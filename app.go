package kick

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrAppAlreadyStarted = errors.New("kick: app already started")
	ErrAppAlreadyStopped = errors.New("kick: app already stopped")
)

type App struct {
	started    bool
	stopped    bool
	startFuncs []Func
	stopFuncs  []Func

	mu sync.Mutex
}

type Func func(ctx context.Context) error

func NewApp() *App {
	return &App{}
}

// OnStart registers a function to be called when the app is started.
func (a *App) OnStart(f ...Func) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.started {
		err := fmt.Errorf("kick: failed to add hook: %w", ErrAppAlreadyStarted)
		panic(err)
	}
	a.startFuncs = append(a.startFuncs, f...)
}

// OnStop registers a function to be called when the app is stopped.
func (a *App) OnStop(f ...Func) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.started {
		err := fmt.Errorf("kick: failed to add hook: %w", ErrAppAlreadyStarted)
		panic(err)
	}
	a.stopFuncs = append(a.stopFuncs, f...)
}

// Start the app.
// It will call all the registered OnStart hooks.
func (a *App) Start(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.started {
		return ErrAppAlreadyStarted
	}
	a.started = true
	for _, f := range a.startFuncs {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Stop the app.
// It will call all the registered OnStop hooks.
func (a *App) Stop(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.stopped {
		return ErrAppAlreadyStopped
	}
	a.stopped = true
	for _, f := range a.stopFuncs {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}
