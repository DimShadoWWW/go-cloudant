package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Security security object
type Security struct {
	Cloudant map[string]SecurityRoles `json:"cloudant"`
	ID       string                   `json:"_id"`
}

// SecurityRoles security roles of the key
type SecurityRoles []string

// Security retrieve the security object
func (c *Client) Security(name string, api *Security) error {
	path := fmt.Sprintf("_api/v2/db/%s/_security", name)

	resp, body, errs := c.conn().Get(c.URL + path).End()
	if errs != nil {
		for i, e := range errs {
			fmt.Printf("error %v: %#v\n", i, e.Error())
		}
		return errs[0]
	}
	if resp.StatusCode >= 400 {
		return errors.New(resp.Status)
	}

	err := json.Unmarshal([]byte(body), &api)
	if err != nil {
		return err
	}
	fmt.Printf("security object: %#v\n", api)
	return nil
}

// PutSecurity apply the security object
func (c *Client) PutSecurity(name string, obj Security) error {
	path := fmt.Sprintf("_api/v2/db/%s/_security", name)

	resp, _, errs := c.conn().Put(c.URL + path).Send(obj).End()
	if errs != nil {
		return errs[0]
	}
	if resp.StatusCode >= 400 {
		return errors.New("Error in setting index: " + strconv.Itoa(resp.StatusCode))
	}
	return nil
}
