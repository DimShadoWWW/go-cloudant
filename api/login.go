package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	request "github.com/parnurzeal/gorequest"
)

// Client ...
type Client struct {
	URL      string
	Username string
	Password string
}

// APIKey ...
type APIKey struct {
	Ok       bool   `json:"-"`
	Key      string `json:"key"`
	Password string `json:"password"`
}

func (c *Client) conn() *request.SuperAgent {
	return request.New().Timeout(30*time.Second).SetBasicAuth(c.Username, c.Password)
}

// NewClient login to Cloudant API
func NewClient(username, password string) (*Client, error) {
	u, err := url.Parse(fmt.Sprintf("https://%s.cloudant.com/", username))
	if err != nil {
		return nil, err
	}
	return &Client{
		URL:      u.String(),
		Username: username,
		Password: password,
	}, nil
}

// CreateAPIKey create a new key/password pair
func (c *Client) CreateAPIKey(api *APIKey) error {
	path := "_api/v2/api_keys"

	resp, body, errs := c.conn().Post(c.URL + path).End()
	if errs != nil {
		for i, e := range errs {
			fmt.Printf("CreateAPIKey error %v: %#v\n", i, e.Error())
		}
		return errs[0]
	}
	if resp.StatusCode >= 400 {
		return errors.New("Error in setting index: " + strconv.Itoa(resp.StatusCode))
	}

	err := json.Unmarshal([]byte(body), &api)
	if err != nil {
		return err
	}
	fmt.Printf("apiKEY: %#v\n", api)
	return nil
}
