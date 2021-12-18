package isdayoff

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

// Client for requests to isdayoff.ru
type Client struct {
	httpClient *http.Client
}

// New initiates client with default http client
func New() *Client {
	return NewWithClient(http.DefaultClient)
}

// NewWithClient initiates client with custom http client
func NewWithClient(client *http.Client) *Client {
	return &Client{client}
}

func LibVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Can't read build info")
		return "0.0.0"
	}
	for _, dep := range bi.Deps {
		if dep.Path == PkgRepoUrl {
			return dep.Version
		}
	}
	fmt.Println("Can't get pkg version")
	return "0.0.0"
}

// IsLeap checks if year is leap
func (c *Client) IsLeap(year int) (bool, error) {
	url := fmt.Sprintf("https://isdayoff.ru/api/isleap?year=%d", year)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, fmt.Errorf("http.NewRequest failed: %v", err)
	}
	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s (%s)", PkgLibName, LibVersion(), PkgRepoUrl))
	res, err := c.httpClient.Do(req)
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

var boolToStr = map[bool]string{
	false: "0",
	true:  "1",
}

// Params contains various filters for request
type Params struct {
	Year        int
	Month       *time.Month
	Day         *int
	CountryCode *CountryCode
	Pre         *bool
	Covid       *bool
	TZ          *string
}

type ParamsRange struct {
	Params
	AfterYear   *int
	AfterMonth  *time.Month
	AfterDay    *int
	BeforeYear  *int
	BeforeMonth *time.Month
	BeforeDay   *int
}

// GetBy Get data by particular params
func (c *Client) GetBy(params Params) ([]DayType, error) {
	url := fmt.Sprintf("https://isdayoff.ru/api/getdata?year=%d", params.Year)
	// ugly. change it later
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
		url += fmt.Sprintf("&pre=%s", boolToStr[*params.Pre])
	}
	if params.Covid != nil {
		url += fmt.Sprintf("&covid=%s", boolToStr[*params.Covid])
	}
	if params.TZ != nil {
		url += fmt.Sprintf("&tz=%s", *params.TZ)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest failed: %v", err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s (%s)", PkgLibName, LibVersion(), PkgRepoUrl))

	res, err := c.httpClient.Do(req)
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

// GetByRange Get data by particular params with date range provided
func (c *Client) GetByRange(params ParamsRange) ([]DayType, error) {
	if params.AfterYear == nil ||
		params.AfterMonth == nil ||
		params.AfterDay == nil ||
		params.BeforeYear == nil ||
		params.BeforeMonth == nil ||
		params.BeforeDay == nil {
		return nil, fmt.Errorf("need to define all of ParamsRange{} values")
	}

	url := fmt.Sprintf("https://isdayoff.ru/api/getdata?date1=%d", *params.AfterYear)
	// ugly. change it later :)
	if *params.AfterMonth < 10 {
		url += fmt.Sprintf("0%d", *params.AfterMonth)
	} else {
		url += fmt.Sprintf("%d", *params.AfterMonth)
	}
	if *params.BeforeDay < 10 {
		url += fmt.Sprintf("0%d", *params.AfterDay)
	} else {
		url += fmt.Sprintf("%d", *params.AfterDay)
	}

	url += fmt.Sprintf("&date2=%d", *params.BeforeYear)
	if *params.BeforeMonth < 10 {
		url += fmt.Sprintf("0%d", *params.BeforeMonth)
	} else {
		url += fmt.Sprintf("%d", *params.BeforeMonth)
	}
	if *params.BeforeDay < 10 {
		url += fmt.Sprintf("0%d", *params.BeforeDay)
	} else {
		url += fmt.Sprintf("%d", *params.BeforeDay)
	}

	if params.CountryCode != nil {
		url += fmt.Sprintf("&cc=%v", *params.CountryCode)
	}
	if params.Pre != nil {
		url += fmt.Sprintf("&pre=%s", boolToStr[*params.Pre])
	}
	if params.Covid != nil {
		url += fmt.Sprintf("&covid=%s", boolToStr[*params.Covid])
	}
	if params.TZ != nil {
		url += fmt.Sprintf("&tz=%s", *params.TZ)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest failed: %v", err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s (%s)", PkgLibName, LibVersion(), PkgRepoUrl))

	res, err := c.httpClient.Do(req)
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
	var result []DayType

	days := strings.Split(string(body), "")
	for _, day := range days {
		result = append(result, DayType(day))
	}

	return result, nil
}

// Today get data for today by particular params
func (c *Client) Today(params Params) (*DayType, error) {
	return c.aliasRequest("today", params)
}

// Tomorrow get data for tomorrow by particular params
func (c *Client) Tomorrow(params Params) (*DayType, error) {
	return c.aliasRequest("tomorrow", params)
}

func (c *Client) aliasRequest(alias string, params Params) (*DayType, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://isdayoff.ru/%s", alias), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest failed: %v", err)
	}

	q := req.URL.Query()
	if params.CountryCode != nil {
		q.Add("cc", string(*params.CountryCode))
	}
	if params.Pre != nil {
		q.Add("pre ", boolToStr[*params.Pre])
	}
	if params.Covid != nil {
		q.Add("covid", boolToStr[*params.Covid])
	}
	if params.TZ != nil {
		q.Add("tz", string(*params.TZ))
	}

	req.URL.RawQuery = q.Encode()

	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s (%s)", PkgLibName, LibVersion(), PkgRepoUrl))

	res, err := c.httpClient.Do(req)
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

	result := DayType(body)

	return &result, nil
}
