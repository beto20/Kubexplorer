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
	pc := di.SetupWorkloadContainer()
	podEp := pc.MustResolve("IWorkloadEndpoint").(endpoint.IWorkloadEndpoint)

	pr := di.SetupParameterContainer()
	parameterEp := pr.MustResolve("IParameterEndpoint").(endpoint.IParameterEndpoint)

	e := di.SetupEnvironmentContainer()
	environmentEp := e.MustResolve("IEnvironmentEndpoint").(endpoint.IEnvironmentEndpoint)

	s := di.SetupNetworkContainer()
	serviceEp := s.MustResolve("INetworkEndpoint").(endpoint.INetworkEndpoint)

	n := di.SetupGeneralContainer()
	nodeEp := n.MustResolve("IGeneralEndpoint").(endpoint.IGeneralEndpoint)

	stg := di.SetupStorageContainer()
	storageEp := stg.MustResolve("IStorageEndpoint").(endpoint.IStorageEndpoint)

	mtc := di.SetupMetricsContainer2()
	metricEp := mtc.MustResolve("IMetricEndpoint").(endpoint.IMetricEndpoint)

	// Create an instance of the app structure
	app := middleware.NewApp()
	workloadMiddleware := middleware.NewWorkloadMiddleware(podEp)
	parameterMiddleware := middleware.NewParameterMiddleware(parameterEp)
	environmentMiddleware := middleware.NewEnvironmentMiddleware(environmentEp)
	networkMiddleware := middleware.NewNetworkMiddleware(serviceEp)
	generalMiddleware := middleware.NewGeneralMiddleware(nodeEp)
	storageMiddleware := middleware.NewStorageMiddleware(storageEp)
	metricMiddleware := middleware.NewMetricMiddleware(metricEp)

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
		OnShutdown:       app.Shutdown,
		Bind: []interface{}{
			app,
			workloadMiddleware,
			parameterMiddleware,
			environmentMiddleware,
			networkMiddleware,
			generalMiddleware,
			storageMiddleware,
			metricMiddleware,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
