package docker

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

// Image Operations

const (
	DefaultTag = "latest"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

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

func (d *Client) getErrorResponse(response *resty.Response) error {
	responseString := response.String()

	var errRes ErrorResponse

	err := json.Unmarshal([]byte(responseString), &errRes)
	if err != nil {
		return err
	}
	return errors.New(errRes.Message)
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