package src

import (
	"net/http"
	"os"
)

type ConfigSingleton interface {
	GetToken() string
	GetCache() *Cache
	GetHttpClient() *http.Client
}

type configSingleton struct {
	cache      *Cache
	token      string
	httpClient *http.Client
}

func (c *configSingleton) GetHttpClient() *http.Client {
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}
	return c.httpClient
}

func (c *configSingleton) GetToken() string {
	if c.token == "" {
		c.token = os.Getenv("API_KEY")
		if c.token == "" {
			c.token = "YourApiKeyToken"
		}
	}
	return c.token
}

func (c *configSingleton) GetCache() *Cache {
	if c.cache == nil {
		c.cache = &Cache{}
		c.cache.cache = make(map[int]*BlockTotalData)
	}
	return c.cache
}

var instance *configSingleton

func GetConfig() ConfigSingleton {
	if instance == nil {
		instance = new(configSingleton)
	}
	return instance
}
