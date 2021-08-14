package goautobahn

import "context"

type roads struct {
	api *Client
}

func newRoads(api *Client) *roads {
	return &roads{
		api: api,
	}
}

type GetAllRoadsResp struct {
	Roads []string `json:"roads"`
}

func (r *roads) GetAll(ctx context.Context) (*GetAllRoadsResp, error) {
	jsonResp := &GetAllRoadsResp{}

	err := r.api.doRequest(ctx, "GET", baseURL, jsonResp)

	if err != nil {
		return nil, err
	}

	return jsonResp, err
}
