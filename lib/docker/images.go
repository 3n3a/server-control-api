package docker

// Image Operations

func (d *Client) PullImage(imageName string) (CatchAllType, error) {
	response, err := d.client.R().
		SetQueryParam("fromImage", imageName).
		SetResult(&CatchAllType{}).
		Post("/images/create")
	return response.Result().(CatchAllType), err
}