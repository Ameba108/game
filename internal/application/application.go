package application

import (
	"context"
	"fmt"
	"time"

	"github.com/Ameba108/game/pkg/life"
)

type Config struct {
	Width  int
	Height int
}

type Application struct {
	Cfg Config
}

func New(config Config) *Application {
	return &Application{
		Cfg: config,
	}
}

func (a *Application) Run(ctx context.Context) error {
	currentWorld := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	nextWorld := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	currentWorld.RandInit(30)
	for {
		fmt.Println(currentWorld)
		life.NextState(currentWorld, nextWorld)
		currentWorld = nextWorld
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(100 * time.Millisecond)
			break
		}
		fmt.Print("\033[H\033[2J")
	}
}
