package twitch

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type Chat struct {
	api      *Integration
	conn     io.ReadWriteCloser
	mux      sync.Mutex
	handlers map[string][]func(msg ChatMessage)
}

func (c *Chat) recreateSocket() (err error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.conn != nil {
		_ = c.conn.Close()
	}

	log.Println("resolving twitch chat...")
	addr, err := net.ResolveTCPAddr("tcp", "irc.chat.twitch.tv:6697")
	if err != nil {
		return
	}
	log.Println("connecting to twitch chat...")
	conn, err := net.DialTCP(addr.Network(), nil, addr)
	if err != nil {
		return
	}
	log.Println("setting up tls...")
	c.conn = tls.Client(conn, &tls.Config{
		ServerName: "irc.chat.twitch.tv",
	})
	return
}

func (c *Chat) Reconnect() error {
	if err := c.recreateSocket(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(c.conn)
	c.Write("PASS oauth:%s", c.api.Token)
	c.Write("NICK %s", c.api.User.Login)
	time.Sleep(1 * time.Second)
	c.Write("USER %s 0 * %s", c.api.User.Login, c.api.User.Login)
	c.Write("JOIN #%s", c.api.User.Login)
	c.Write("CAP REQ :twitch.tv/tags")
	c.Write("CAP REQ :twitch.tv/membership")
	c.Write("CAP REQ :twitch.tv/commands")

	go func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("RECEIVE: %s", line)
			message := messageFromLine(line)
			switch strings.ToUpper(message.Command) {
			case CommandPing:
				c.Write(strings.Replace(line, "PING", "PONG", 1))
				break
			default:
				c.emitMessage(message)
				break
			}
		}
	}(scanner)

	log.Println("connected")
	return nil
}

func (c *Chat) OnMessage(command string, handler func(msg ChatMessage)) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.handlers == nil {
		c.handlers = map[string][]func(msg ChatMessage){}
	}
	commandKey := strings.ToUpper(command)
	if list, ok := c.handlers[commandKey]; ok {
		c.handlers[commandKey] = append(list, handler)
	} else {
		c.handlers[commandKey] = []func(msg ChatMessage){handler}
	}
}

func (c *Chat) SendMessage(body string) {
	c.Write("PRIVMSG #%s :%s", c.api.User.Login, body)
}

func (c *Chat) Write(format string, params ...interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()

	command := fmt.Sprintf(format, params...)
	log.Printf("SEND   : %s", strings.ReplaceAll(command, c.api.Token, "[REDACTED]"))
	_, err := c.conn.Write([]byte(command + "\n"))
	if err != nil {
		log.Printf("Failed to send command [%s]: %s", command, err)
	}
}

func (c *Chat) emitMessage(message ChatMessage) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.handlers == nil {
		return
	}
	if handlers, ok := c.handlers[strings.ToUpper(message.Command)]; ok {
		for _, handler := range handlers {
			go handler(message)
		}
	}
}
