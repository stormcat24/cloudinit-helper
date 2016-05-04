package meta

import (
	"time"
	"net/http"
)

type Config struct {
	UrlBase string
	Timeout time.Duration
	UseMock bool
}

func NewClient(conf *Config) Client {

	if conf.UseMock {
		return ClientMock{}
	} else {

		httpCli := &http.Client{
			Timeout: conf.Timeout,
		}

		return ClientImpl{
			urlBase: conf.UrlBase,
			httpCli: httpCli,
		}
	}
}

