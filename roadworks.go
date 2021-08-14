package goautobahn

import (
	"context"
)

type roadworks struct {
	api *Client
}

func newRoadworks(api *Client) *roadworks {
	return &roadworks{
		api: api,
	}
}

type Roadwork struct {
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
	Impact                   struct {
		Lower   string   `json:"lower"`
		Upper   string   `json:"upper"`
		Symbols []string `json:"symbols"`
	} `json:"impact,omitempty"`
}

type GetRoadworksResp struct {
	Roadworks []*Roadwork `json:"roadworks"`
}

type GetDetailsForRoadworkResp struct {
	Roadwork *Roadwork
}

func (r *roadworks) GetAllForRoad(ctx context.Context, roadID string) (*GetRoadworksResp, error) {
	jsonResp := &GetRoadworksResp{}

	url := baseURL + roadID + "/services/roadworks"
	err := r.api.doRequest(ctx, "GET", url, jsonResp)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}

func (r *roadworks) GetDetailsFor(ctx context.Context, identifier string) (*GetDetailsForRoadworkResp, error) {
	jsonResp := &GetDetailsForRoadworkResp{Roadwork: &Roadwork{}}

	url := baseURL + "details/roadworks/" + identifier
	err := r.api.doRequest(ctx, "GET", url, jsonResp.Roadwork)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}
