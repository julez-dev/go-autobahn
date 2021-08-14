package goautobahn

import "context"

type electricChargingStations struct {
	api *Client
}

func newElectricChargingStations(api *Client) *electricChargingStations {
	return &electricChargingStations{
		api: api,
	}
}

// ElectricChargingStation represents a electric charging station
type ElectricChargingStation struct {
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
}

type GetElectricChargingStationsResp struct {
	ElectricChargingStations []*ElectricChargingStation `json:"electric_charging_station"`
}

type GetDetailsForElectricChargingStationResp struct {
	ElectricChargingStation *ElectricChargingStation `json:"electric_charging_station"`
}

// GetAllForRoad returns all electric charging stations for the provided road
func (e *electricChargingStations) GetAllForRoad(ctx context.Context, roadID string) (*GetElectricChargingStationsResp, error) {
	jsonResp := &GetElectricChargingStationsResp{}

	url := baseURL + roadID + "/services/electric_charging_station"
	err := e.api.doRequest(ctx, "GET", url, jsonResp)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}

// GetDetailsFor returns details for the provided electric charging station
func (e *electricChargingStations) GetDetailsFor(ctx context.Context, identifier string) (*GetDetailsForElectricChargingStationResp, error) {
	jsonResp := &GetDetailsForElectricChargingStationResp{ElectricChargingStation: &ElectricChargingStation{}}

	url := baseURL + "details/electric_charging_station/" + identifier
	err := e.api.doRequest(ctx, "GET", url, jsonResp.ElectricChargingStation)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}
