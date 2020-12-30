package isdayoff

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client клиент для запросов в isdayoff.ru
type Client struct {
	*http.Client
}

// New ...
func New(client *http.Client) *Client {
	return &Client{client}
}

// IsLeap Проверка года на високосность
func (c *Client) IsLeap(year int) (bool, error) {
	url := fmt.Sprintf("https://isdayoff.ru/api/isleap?year=%d", year)
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return false, fmt.Errorf("http.NewRequest failed: %v", err)
	}
	res, err := c.Do(req)
	if err != nil {
		return false, fmt.Errorf("client.Do(req) failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, fmt.Errorf("ioutil.ReadAll failed: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf(string(body))
	}

	return YearType(string(body)) == YearTypeLeap, nil
}
