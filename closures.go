package goautobahn

import "context"

type closures struct {
	api *Client
}

func newClosures(api *Client) *closures {
	return &closures{
		api: api,
	}
}

// Closure represents a closure
type Closure struct {
	Extent              string   `json:"extent"`
	Identifier          string   `json:"identifier"`
	RouteRecommendation []string `json:"routeRecommendation"`
	Coordinate          struct {
		Lat  string `json:"lat"`
		Long string `json:"long"`
	} `json:"coordinate"`
	Footer                   []string `json:"footer"`
	Icon                     string   `json:"icon"`
	IsBlocked                bool     `json:"isBlocked"`
	Description              []string `json:"description"`
	Title                    string   `json:"title"`
	Point                    string   `json:"point"`
	DisplayType              string   `json:"display_type"`
	LorryParkingFeatureIcons []string `json:"lorryParkingFeatureIcons"`
	Future                   bool     `json:"future"`
	Subtitle                 string   `json:"subtitle"`
	StartTimestamp           string   `json:"startTimestamp"`
}

type GetClosuresResp struct {
	Closures []*Closure `json:"closure"`
}

type GetDetailsForClosureResp struct {
	Closure *Closure `json:"closure"`
}

// GetAllForRoad returns all closures for the provided road
func (c *closures) GetAllForRoad(ctx context.Context, roadID string) (*GetClosuresResp, error) {
	jsonResp := &GetClosuresResp{}

	url := baseURL + roadID + "/services/closure"
	err := c.api.doRequest(ctx, "GET", url, jsonResp)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}

// GetDetailsFor returns details for the provided closure
func (c *closures) GetDetailsFor(ctx context.Context, identifier string) (*GetDetailsForClosureResp, error) {
	jsonResp := &GetDetailsForClosureResp{Closure: &Closure{}}

	url := baseURL + "details/closure/" + identifier
	err := c.api.doRequest(ctx, "GET", url, jsonResp.Closure)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}
