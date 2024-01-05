package main

import (
	"time"

	server "github.com/3n3a/server-control-api/lib/server"
)

var version string
var appConfig = server.AppConfig{
	CACHE_INCLUDE_RAW: "", // regex for caching only certain routes
	CACHE_LENGTH:      30 * time.Minute,
	APP_PORT:          3000,
	VERSION:           version,
}

func main() {
	appConfig.Setup()
}
