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

	pr := di.SetupParameterContainer()
	parameterEp := pr.MustResolve("IParameterEndpoint").(endpoint.IParameterEndpoint)

	e := di.SetupEnvironmentContainer()
	environmentEp := e.MustResolve("IEnvironmentEndpoint").(endpoint.IEnvironmentEndpoint)

	s := di.SetupServiceContainer()
	serviceEp := s.MustResolve("IServiceEndpoint").(endpoint.IServiceEndpoint)

	n := di.SetupNodeContainer()
	nodeEp := n.MustResolve("INodeEndpoint").(endpoint.INodeEndpoint)

	// Create an instance of the app structure
	app := middleware.NewApp()
	deploymentMiddleware := middleware.NewDeploymentMiddleware(deploymentEp)
	podMiddleware := middleware.NewPodMiddleware(podEp)
	parameterMiddleware := middleware.NewParameterMiddleware(parameterEp)
	environmentMiddleware := middleware.NewEnvironmentMiddleware(environmentEp)
	serviceMiddleware := middleware.NewServiceMiddleware(serviceEp)
	nodeMiddleware := middleware.NewNodeMiddleware(nodeEp)

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
			deploymentMiddleware,
			podMiddleware,
			parameterMiddleware,
			environmentMiddleware,
			serviceMiddleware,
			nodeMiddleware,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
