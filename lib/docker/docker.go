package docker

import (
	"github.com/3n3a/server-control-api/lib/utils"
	"github.com/go-resty/resty/v2"
)

// Docker Client
//
// * also supports podman (uses libpods docker-compatible api)
// * may need to be run with "sudo"
//
// Sources:
// 
// * Docker Rest Api Reference: https://docs.docker.com/engine/api/latest
// * Podman Rest Api Reference: https://docs.podman.io/en/latest/_static/api.html

type ContainerTool int

const (
	DockerTool ContainerTool = iota
	PodmanTool
)

type Client struct {
	containerTool ContainerTool 
	socketUrl string

	client *resty.Client
}

type CatchAllType []map[string]interface{}

func New(containerTool ContainerTool) Client {
	d := Client{}
	d.containerTool = containerTool
	d.setupSocketUrl()
	d.client = utils.NewUnixHTTPClient(d.socketUrl)
	return d
}

func (d *Client) setupSocketUrl() {
	switch d.containerTool {
	case DockerTool:
		d.socketUrl = "/var/run/docker.sock"
	case PodmanTool:
		d.socketUrl = "/run/podman/podman.sock"
	}
}