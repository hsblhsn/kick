package kick_test

import (
	"context"
	"github.com/hsblhsn/kick"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	app := kick.NewApp()

	err := app.Start(context.Background())
	require.NoError(t, err)

	err = app.Start(context.Background())
	require.Error(t, err)
	require.ErrorIs(t, err, kick.ErrAppAlreadyStarted)

	err = app.Stop(context.Background())
	require.NoError(t, err)

	err = app.Stop(context.Background())
	require.Error(t, err)
	require.ErrorIs(t, err, kick.ErrAppAlreadyStopped)
}

func TestApp_ConcurrentStart(t *testing.T) {
	app := kick.NewApp()

	counter := 0
	app.OnStart(func(ctx context.Context) error {
		counter++
		return nil
	})

	app.OnStop(func(ctx context.Context) error {
		counter++
		return nil
	})

	for i := 0; i < 1000; i++ {
		go func() {
			_ = app.Start(context.Background())
		}()
	}
	time.Sleep(time.Second)
	require.Equal(t, 1, counter)

	for i := 0; i < 1000; i++ {
		go func() {
			_ = app.Stop(context.Background())
		}()
	}
	time.Sleep(time.Second)
	require.Equal(t, 2, counter)
}
