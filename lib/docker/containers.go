package docker

import (
	"errors"
	"fmt"
)

const (
	DefaultContainerKillSignal = "SIGTERM"
	DefaultContainerKillTimeout = "10"
)

// Container Operations
func (d *Client) RestartContainer(containerName string) (error) {
	err := d.restartContainerByIdOrName(containerName)

	// Default Error Response
	if err != nil {
		return errors.New("container could not be restarted")
	}

	return nil
}

func (d *Client) restartContainerByIdOrName(containerIdOrName string) (error) {
	dockerUrl := fmt.Sprintf("http://localhost/containers/%s/restart", containerIdOrName)
	res, err := d.client.R().
		SetQueryParam("signal", DefaultContainerKillSignal).
		SetQueryParam("t", DefaultContainerKillTimeout).
		Post(dockerUrl)

	// Specific Error Responses
	if res.StatusCode() == 404 || res.StatusCode() == 500 {
		return d.getErrorResponse(res)
	}

	return err
}