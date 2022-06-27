package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
	"golang.org/x/oauth2/clientcredentials"
)

const TokenURL string = "https://auth.lumeris.io/as/token"
const ApiURL string = "https://api.pingone.com/v1"
const EnvironmentPath string = "environments"

type Client struct {
	TokenURL    string
	ApiURL      string
	Context     context.Context
	AccessToken string
}

func NewClient(ctx context.Context, clientID, clientSecret, tokenURL, apiURL *string) (*Client, error) {
	client := Client{
		TokenURL: TokenURL,
		ApiURL:   ApiURL,
		Context:  ctx,
	}

	if tokenURL != nil {
		client.TokenURL = *tokenURL
	}
	if apiURL != nil {
		client.ApiURL = *apiURL
	}

	conf := clientcredentials.Config{
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		TokenURL:     client.TokenURL,
	}

	token, err := conf.Token(ctx)
	if err != nil {
		return nil, err
	}

	client.AccessToken = token.AccessToken

	return &client, nil
}

func (client *Client) doRequest(req *http.Request) ([]byte, error) {
	token := client.AccessToken
	httpClient := &http.Client{Timeout: 10 * time.Second}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func GetAll[V models.AllTheThings](c *Client, params map[string]string, segments ...string) (*models.GenericResponse[V], error) {
	segments = append([]string{c.ApiURL}, segments...)
	uri := strings.Join(segments, "/")

	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	genericResponse := models.GenericResponse[V]{}
	err = json.Unmarshal(body, &genericResponse)
	if err != nil {
		return nil, err
	}

	return &genericResponse, nil
}

func GetAllFromEnvironment[V models.AllTheThings](c *Client, environmentId string, params map[string]string, segments ...string) (*models.GenericResponse[V], error) {
	segments = append([]string{"environments", environmentId}, segments...)
	return GetAll[V](c, params, segments...)
}

func Get[T models.AllTheThings](c *Client, segments ...string) (*T, error) {
	segments = append([]string{c.ApiURL}, segments...)
	url := strings.Join(segments, "/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response T
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func GetFromEnvironment[V models.AllTheThings](c *Client, environmentId string, segments ...string) (*V, error) {
	segments = append([]string{"environments", environmentId}, segments...)
	return Get[V](c, segments...)
}

// GetStringFromEnvironment return non-json content as a text string.
func GetStringFromEnvironment(c *Client, environmentId string, acceptHeader string, segments ...string) (*string, error) {
	segments = append([]string{c.ApiURL, "environments", environmentId}, segments...)
	url := strings.Join(segments, "/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", acceptHeader)

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var bodyString = string(body)

	return &bodyString, nil
}

func CreateForEnvironment[V models.AllTheThings](c *Client, environmentId string, obj V, segments ...string) (*V, error) {
	segments = append([]string{c.ApiURL, "environments", environmentId}, segments...)
	uri := strings.Join(segments, "/")

	encoded, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, strings.NewReader(string(encoded)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response V
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PutForEnvironment[V models.AllTheThings](c *Client, environmentId string, obj V, segments ...string) (*V, error) {
	segments = append([]string{c.ApiURL, "environments", environmentId}, segments...)
	uri := strings.Join(segments, "/")

	encoded, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", uri, strings.NewReader(string(encoded)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response V
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func DeleteForEnvironment(c *Client, environmentId string, segments ...string) error {
	segments = append([]string{c.ApiURL, "environments", environmentId}, segments...)
	uri := strings.Join(segments, "/")

	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
