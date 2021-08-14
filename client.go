package goautobahn

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://verkehr.autobahn.de/o/autobahn/"

// Client holds all endpoints
type Client struct {
	http *http.Client

	Roadworks                *roadworks                // Endpoint for roadworks
	Roads                    *roads                    // Endpoint for roads (autobahnen)
	Webcams                  *webcams                  // Endpoint for webcams
	ParkingLorries           *parkingLorries           // Endpoint for parking lorries
	Warnings                 *warnings                 // Endpoint for warning
	Closures                 *closures                 // Endpoint for colusures
	ElectricChargingStations *electricChargingStations // Endpoint for electric charging stations
}

// doRequest makes a request and decodes the json response into v
// will return error if the returned status code is not 200
func (c *Client) doRequest(ctx context.Context, method string, url string, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)

	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status (%s) not ok or empty body", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return nil
	}

	return nil
}

type OptionFunc func(*Client)

// New creates a new client with the provided options applied
func New(options ...OptionFunc) *Client {
	client := &Client{}

	for _, f := range options {
		f(client)
	}

	if client.http == nil {
		client.http = &http.Client{}
	}

	client.Roadworks = newRoadworks(client)
	client.Roads = newRoads(client)
	client.Webcams = newWebcams(client)
	client.ParkingLorries = newParkingLorries(client)
	client.Warnings = newWarnings(client)
	client.Closures = newClosures(client)
	client.ElectricChargingStations = newElectricChargingStations(client)

	return client
}

// WithHTTPClient sets a custom http.Client
// If not set a new client will be created
func WithHTTPClient(httpClient *http.Client) OptionFunc {
	return func(c *Client) {
		c.http = httpClient
	}
}
