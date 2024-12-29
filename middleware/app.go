package middleware

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// TODO: Se agrego para Wails events
// EmitMessages emits messages to the frontend every second
func (a *App) EmitMessages() {
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Message %d", i)
		runtime.EventsEmit(a.ctx, "new-message", message)
		time.Sleep(1 * time.Second)
	}
}
