package goautobahn

import (
	"context"
)

type warnings struct {
	api *Client
}

func newWarnings(api *Client) *warnings {
	return &warnings{
		api: api,
	}
}

// Warning represents a warning
type Warning struct {
	Extent              string   `json:"extent"`
	Identifier          string   `json:"identifier"`
	RouteRecommendation []string `json:"routeRecommendation"`
	Coordinate          struct {
		Lat  string `json:"lat"`
		Long string `json:"long"`
	} `json:"coordinate"`
	Footer                   []string   `json:"footer"`
	Icon                     string     `json:"icon"`
	IsBlocked                bool       `json:"isBlocked"`
	Description              []string   `json:"description"`
	Title                    string     `json:"title"`
	Point                    string     `json:"point"`
	DisplayType              string     `json:"display_type"`
	LorryParkingFeatureIcons []string   `json:"lorryParkingFeatureIcons"`
	Future                   bool       `json:"future"`
	Subtitle                 string     `json:"subtitle"`
	StartTimestamp           CustomTime `json:"startTimestamp"`
}

type GetWarningsResp struct {
	Warnings []*Warning `json:"warning"`
}

type GetDetailsForwarningResp struct {
	Warning *Warning `json:"warning"`
}

// GetAllForRoad returns all warnings for the provided road
func (w *warnings) GetAllForRoad(ctx context.Context, roadID string) (*GetWarningsResp, error) {
	jsonResp := &GetWarningsResp{}

	url := baseURL + roadID + "/services/warning"
	err := w.api.doRequest(ctx, "GET", url, jsonResp)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}

// GetDetailsFor returns details for the provided warning
func (w *warnings) GetDetailsFor(ctx context.Context, identifier string) (*GetDetailsForwarningResp, error) {
	jsonResp := &GetDetailsForwarningResp{Warning: &Warning{}}

	url := baseURL + "details/warning/" + identifier
	err := w.api.doRequest(ctx, "GET", url, jsonResp.Warning)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}
