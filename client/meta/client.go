package meta

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

type Client interface {
	GetAvailabilityZone() (*AvailabilityZone, error)
	GetInstanceID() (string, error)
}

type ClientImpl struct {
	httpCli *http.Client
	urlBase string
}

func (c ClientImpl) GetAvailabilityZone() (*AvailabilityZone, error) {

	apiUrl := fmt.Sprintf("%s%s", c.urlBase, "/latest/meta-data/placement/availability-zone")
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("StatusCode: %v, Message: %v", resp.StatusCode, string(data))
	}

	return &AvailabilityZone{
		Name: string(data),
	}, nil
}

func (c ClientImpl) GetInstanceID() (string, error) {
	apiUrl := fmt.Sprintf("%s%s", c.urlBase, "/latest/meta-data/instance-id")
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpCli.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("StatusCode: %v, Message: %v", resp.StatusCode, string(data))
	}

	return string(data), nil
}