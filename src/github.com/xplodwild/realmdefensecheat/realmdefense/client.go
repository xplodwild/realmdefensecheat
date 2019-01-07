package realmdefense

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	EndpointBabeltimeUS = "http://td-users2.babeltimeus.com"
	EndpointBabeltimeCA = "http://td-users2.babeltimeca.com"
)

type Client struct {
	endpoint  string
	userAgent string
	cli       *http.Client
}

func NewClient(endpoint string, userAgent string) *Client {
	return &Client{
		endpoint:  endpoint,
		userAgent: userAgent,
		cli:       &http.Client{},
	}
}

func (c *Client) POST(api string, body []byte, gz bool, acceptGz bool) ([]byte, error) {
	// Create a new request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.endpoint, api), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Common headers
	req.Header.Set("X-Unity-Version", "2017.4.8f1")
	req.Header.Set("User-Agent", c.endpoint)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Gzip support
	if acceptGz {
		req.Header.Set("Accept-Gzip", "True")
	}

	if gz {
		fmt.Printf("Body length: %d\n", len(body))
		req.Header.Set("Gzip", "True")

		var buf bytes.Buffer
		zw, _ := gzip.NewWriterLevel(&buf, gzip.BestCompression)
		zw.Write(body)
		zw.Close()

		body = buf.Bytes()
	}

	// Compute hash
	req.Header.Set("Hash", ComputeHash(body))

	// Send the request
	res, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}

	// Read the response
	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	/*fmt.Printf("Raw result: %s\n", resBody)
	fmt.Printf("Raw headers: %+v\n", res.Header)
	fmt.Printf("Status: %d\n", res.StatusCode)*/

	if res.StatusCode != 200 {
		fmt.Printf("Got %d HTTP status code!\n", res.StatusCode)
	}

	// Handle Gzip responses
	if res.Header.Get("Gzip") == "True" {
		zr, err := gzip.NewReader(bytes.NewReader(resBody))
		if err != nil {
			return nil, err
		}
		resBodyGunzip, err := ioutil.ReadAll(zr)
		if err != nil {
			return nil, err
		}

		resBody = resBodyGunzip
		zr.Close()
	}

	return resBody, err
}
