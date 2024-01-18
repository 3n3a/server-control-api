package docker

import (
	"errors"
	"fmt"
	"strings"
)

// Image Operations

const (
	DefaultTag = "latest"
)

func (d *Client) PullImage(imageName string) (error) {
	imageUrl, tag := d.getImageAndTag(imageName)

	dockerUrl := fmt.Sprintf("http://localhost%s", "/images/create")
	res, err := d.client.R().
		SetQueryParam("fromImage", imageUrl).
		SetQueryParam("tag", tag).
		Post(dockerUrl)

	// Specific Error Responses
	if res.StatusCode() == 404 || res.StatusCode() == 500 {
		return d.getErrorResponse(res)
	}

	// Default Error Response
	if err != nil {
		return errors.New("image could not be pulled")
	}

	return nil
}



func (d *Client) getImageAndTag(imageName string) (string, string) {
	splitResult := strings.Split(imageName, ":")

	imageUrl := splitResult[0]
	tag := DefaultTag

	// set tag only if it has a content
	if len(splitResult) > 1 && splitResult[1] != "" {
		tag = splitResult[1]
	}
	
	return imageUrl, tag
}