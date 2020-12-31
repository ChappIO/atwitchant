package twitch

import (
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type cacheItem struct {
	Data      []byte
	ExpiresOn time.Time
}

type cache struct {
	Data map[string]cacheItem
	Lock sync.Mutex
}

func (c *cache) Get(token string, url string) ([]byte, error) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	if elem, ok := c.Data[url]; ok {
		if elem.ExpiresOn.After(time.Now()) {
			return elem.Data, nil
		}
	}

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Client-Id", clientId)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	c.Data[url] = cacheItem{
		Data:      data,
		ExpiresOn: time.Now().Add(15 * time.Second),
	}
	return data, nil
}
