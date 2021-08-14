package goautobahn

import "context"

type webcams struct {
	api *Client
}

func newWebcams(api *Client) *webcams {
	return &webcams{
		api: api,
	}
}

type Webcam struct {
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
	Operator                 string   `json:"operator"`
	Point                    string   `json:"point"`
	DisplayType              string   `json:"display_type"`
	LorryParkingFeatureIcons []string `json:"lorryParkingFeatureIcons"`
	Future                   bool     `json:"future"`
	Imageurl                 string   `json:"imageurl"`
	Subtitle                 string   `json:"subtitle"`
	Linkurl                  string   `json:"linkurl"`
}

type GetWebcamsResp struct {
	Webcams []*Webcam `json:"webcam"`
}

type GetDetailsForWebcamResp struct {
	Webcam *Webcam `json:"webcam"`
}

func (w *webcams) GetAllForRoad(ctx context.Context, roadID string) (*GetWebcamsResp, error) {
	jsonResp := &GetWebcamsResp{}

	url := baseURL + roadID + "/services/webcam"
	err := w.api.doRequest(ctx, "GET", url, jsonResp)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}

func (w *webcams) GetDetailsFor(ctx context.Context, identifier string) (*GetDetailsForWebcamResp, error) {
	jsonResp := &GetDetailsForWebcamResp{Webcam: &Webcam{}}

	url := baseURL + "details/webcam/" + identifier
	err := w.api.doRequest(ctx, "GET", url, jsonResp.Webcam)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}
