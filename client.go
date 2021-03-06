package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Client struct {
	u string
	c *http.Client
}

func New(u string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	c := &http.Client{
		Timeout: time.Second * 10,
		Jar:     jar,
	}
	return &Client{
		u: u,
		c: c,
	}, nil
}

func (c *Client) get(p string, v interface{}) error {
	resp, err := c.c.Get(c.u + p)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Failed http request. Status:%d, Error:%s", resp.StatusCode, string(body))
	}
	if v == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(v)
}
func (c *Client) postWithResponse(p string, v interface{}, ret interface{}) error {
	buf := new(bytes.Buffer)
	if v != nil {
		if err := json.NewEncoder(buf).Encode(v); err != nil {
			return err
		}
	}
	req, err := http.NewRequest("POST", c.u+p, buf)
	if err != nil {
		return err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Failed http request. Status:%d, Error:%s", resp.StatusCode, string(body))
	}
	if ret == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(ret)
}

func (c *Client) post(p string, v interface{}) error {
	return c.postWithResponse(p, v, nil)
}
func (c *Client) put(p string, v interface{}) error {
	buf := new(bytes.Buffer)
	if v != nil {
		if err := json.NewEncoder(buf).Encode(v); err != nil {
			return err
		}
	}
	req, err := http.NewRequest("PUT", c.u+p, buf)
	if err != nil {
		return err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Failed http request. Status:%d, Error:%s", resp.StatusCode, string(body))
	}
	return nil
}

func (c *Client) delete(p string) error {
	buf := new(bytes.Buffer)
	req, err := http.NewRequest("DELETE", c.u+p, buf)
	if err != nil {
		return err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if err != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			return fmt.Errorf("Failed http request. Status:%d, Error:%s", resp.StatusCode, string(body))
		}
	}
	return nil
}
