package isdayoff

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, fmt.Errorf("ioutil.ReadAll failed: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf(string(body))
	}

	return YearType(string(body)) == YearTypeLeap, nil
}

var boolToInt = map[bool]int{
	false: 0,
	true:  1,
}

// Params параметры запроса
type Params struct {
	Year        int
	Month       *time.Month
	Day         *int
	CountryCode *CountryCode
	Pre         *bool
	Covid       *bool
}

// GetBy Получение данных за определенное время (год, месяц, день)
// https://isdayoff.ru/api/getdata?year=YYYY&month=MM&day=DD[&cc=xx&pre=[0|1]&covid=[0|1]]
func (c *Client) GetBy(params Params) ([]DayType, error) {
	url := fmt.Sprintf("https://isdayoff.ru/api/getdata?year=%d", params.Year)
	if params.Month != nil {
		if *params.Month < 10 {
			url += fmt.Sprintf("&month=0%d", *params.Month)
		} else {
			url += fmt.Sprintf("&month=%d", *params.Month)
		}
	}
	if params.Day != nil {
		if *params.Day < 10 {
			url += fmt.Sprintf("&day=0%d", *params.Day)
		} else {
			url += fmt.Sprintf("&day=%d", *params.Day)
		}
	}
	if params.CountryCode != nil {
		url += fmt.Sprintf("&cc=%v", *params.CountryCode)
	}
	if params.Pre != nil {
		url += fmt.Sprintf("&pre=%d", boolToInt[*params.Pre])
	}
	if params.Covid != nil {
		url += fmt.Sprintf("&covid=%d", boolToInt[*params.Covid])
	}
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest failed: %v", err)
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do(req) failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll failed: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := []DayType{}

	days := strings.Split(string(body), "")
	for _, day := range days {
		result = append(result, DayType(day))
	}

	return result, nil
}
