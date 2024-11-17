package main

import (
	"Kubessistant/backend/di"
	"Kubessistant/backend/endpoint"
	"Kubessistant/middleware"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	PROGRAM_NAME = "Kubessistant"
	WIDTH        = 1024
	HEIGHT       = 768
)

func main() {
	dc := di.SetupDeploymentContainer()
	deploymentEp := dc.MustResolve("IDeploymentEndpoint").(endpoint.IDeploymentEndpoint)
	pc := di.SetupPodContainer()
	podEp := pc.MustResolve("IPodEndpoint").(endpoint.IPodEndpoint)

	// Create an instance of the app structure
	app := middleware.NewApp()
	d := middleware.NewDeploymentMiddleware(deploymentEp)
	p := middleware.NewPodMiddleware(podEp)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  PROGRAM_NAME,
		Width:  WIDTH,
		Height: HEIGHT,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
			d,
			p,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
