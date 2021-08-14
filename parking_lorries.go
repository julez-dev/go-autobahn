package goautobahn

import (
	"context"
	"log"
)

type parkingLorries struct {
	api *Client
}

func newParkingLorries(api *Client) *parkingLorries {
	return &parkingLorries{
		api: api,
	}
}

// ParkingLorry represents a parking lorry
type ParkingLorry struct {
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
	LorryParkingFeatureIcons []struct {
		Icon        string `json:"icon"`
		Description string `json:"description"`
		Style       string `json:"style"`
	} `json:"lorryParkingFeatureIcons"`
	Future   bool   `json:"future"`
	Subtitle string `json:"subtitle"`
}

type GetParkingLorriesResp struct {
	ParkingLorries []*ParkingLorry `json:"parking_lorry"`
}

type GetDetailsForParkingLorryResp struct {
	ParkingLorry *ParkingLorry `json:"parking_lorry"`
}

// GetAllForRoad returns all parking lorries for the provided road
func (pl *parkingLorries) GetAllForRoad(ctx context.Context, roadID string) (*GetParkingLorriesResp, error) {
	jsonResp := &GetParkingLorriesResp{}

	url := baseURL + roadID + "/services/parking_lorry"
	log.Println(url)
	err := pl.api.doRequest(ctx, "GET", url, jsonResp)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}

// GetDetailsFor returns details for the provided parking lorry
func (pl *parkingLorries) GetDetailsFor(ctx context.Context, identifier string) (*GetDetailsForParkingLorryResp, error) {
	jsonResp := &GetDetailsForParkingLorryResp{ParkingLorry: &ParkingLorry{}}

	url := baseURL + "details/parking_lorry/" + identifier
	err := pl.api.doRequest(ctx, "GET", url, jsonResp.ParkingLorry)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}
