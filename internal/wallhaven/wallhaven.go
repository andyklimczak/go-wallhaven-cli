package wallhaven

import (
	"encoding/json"
	"fmt"
	"github.com/andyklimczak/go-wallhaven/internal/logger"
	"net/http"
)

const WallhavenApiHost = "https://wallhaven.cc"

type Client struct {
	httpClient HttpClient
	host       string
	apikey     string
	username   string
	log        *logger.GoHavenLogger
}

func NewClient(username string, apikey string, log *logger.GoHavenLogger) *Client {
	client := DefaultHttpClient()
	return &Client{
		httpClient: client,
		host:       WallhavenApiHost,
		apikey:     apikey,
		username:   username,
		log:        log,
	}
}

func (c *Client) CollectionsForApikey() (CollectionData, error) {
	endpoint := fmt.Sprintf("api/v1/collections/%s?apikey=%s", c.username, c.apikey)
	c.log.Debug("Getting collection at url: %s", endpoint)
	res, err := c.do(http.MethodGet, endpoint)
	var collectionData CollectionData
	if err != nil {
		return collectionData, fmt.Errorf("Unable to get collections at %s: %w", endpoint, err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&collectionData)
	if err != nil {
		return collectionData, fmt.Errorf("Unable to decode body: %w")
	}
	return collectionData, nil
}

func (c *Client) ListingsForCollection(collection *Collection) (SearchData, error) {
	endpoint := fmt.Sprintf("api/v1/collections/%s/%d?apikey=%s", c.username, collection.Id, c.apikey)
	c.log.Debug("Getting listings for collection at url: %s", endpoint)
	res, err := c.do(http.MethodGet, endpoint)
	var searchData SearchData
	if err != nil {
		return searchData, fmt.Errorf("Unable to get collections at %s: %w", endpoint, err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&searchData)
	if err != nil {
		return searchData, fmt.Errorf("Unable to decode body: %w")
	}
	return searchData, nil
}

func (c *Client) do(method, endpoint string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.httpClient.Do(req)
}
