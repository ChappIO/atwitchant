package twitch

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

type Chat struct {
	api  *Integration
	conn io.ReadWriteCloser
	mux sync.Mutex
}

func (c *Chat) Reconnect() (err error) {
	if c.conn != nil {
		_ = c.conn.Close()
	}
	log.Println("connecting to twitch chat...")

	addr, err := net.ResolveTCPAddr("tcp", "irc.chat.twitch.tv:6697")
	if err != nil {
		return
	}
	conn, err := net.DialTCP(addr.Network(), nil, addr)
	if err != nil {
		return
	}
	c.conn = tls.Client(conn, &tls.Config{
		ServerName: "irc.chat.twitch.tv",
	})
	go func() {
		c, err := io.Copy(os.Stdout, c.conn)
		log.Printf("%d", c)
		if err != nil {
			panic(err)
		}
	}()
	c.Write("PASS oauth:%s", c.api.token)
	c.Write("NICK %s", c.api.user.Login)
	c.Write("USER %s 0 * %s", c.api.user.Login, c.api.user.Login)
	c.Write("JOIN #%s", c.api.user.Login)
	c.Write("CAP REQ :twitch.tv/tags")
	c.Write("CAP REQ :twitch.tv/membership")
	c.Write("CAP REQ :twitch.tv/commands")

	log.Println("connected")
	return
}

func (c *Chat) Write(format string, params ...interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()

	command := fmt.Sprintf(format, params...)
	log.Printf("SEND   : %s", command)
	_, err := c.conn.Write([]byte(command + "\n"))
	if err != nil {
		log.Printf("Failed to send command [%s]: %s", command, err)
	}
}
