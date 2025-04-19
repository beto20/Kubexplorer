package middleware

import (
	"context"
	"fmt"
)

type App struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	//StartBackgroundProcesses()
	a.ctx = ctx
}

//func StartBackgroundProcesses() {
//	metrics := di.SetupMetricsContainer()
//	metricsBackground := metrics.MustResolve("IMetricBackground").(background.MetricBackground)
//	go func() {
//		metricsBackground.PodMetricsBackgroundProcess()
//		//background.MetricBackground{}.PodMetricsBackgroundProcess()
//	}()
//}

func (a *App) Shutdown(ctx context.Context) {
	fmt.Println("🛑 App shutting down")
	if a.cancelFunc != nil {
		a.cancelFunc()
	}
}

//// TODO: Se agrego para Wails events
//// EmitMessages emits messages to the frontend every second
//func (a *App) EmitMessages() {
//	for i := 0; i < 10; i++ {
//		message := fmt.Sprintf("Message %d", i)
//		runtime.EventsEmit(a.ctx, "new-message", message)
//		time.Sleep(1 * time.Second)
//	}
//}
