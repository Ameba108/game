package life_test

import (
	"testing"

	"github.com/Ameba108/game/pkg/life"
)

func TestNewWorld(t *testing.T) {
	height := 10
	width := 4
	world, err := life.NewWorld(height, width)
	if err != nil {
		t.Errorf("somrthing went wrong with func NewWorld")
	}
	if world.Height != height {
		t.Errorf("expected height: %d, actual height: %d", height, world.Height)
	}
	if world.Width != width {
		t.Errorf("expected width: %d, actual width: %d", width, world.Width)
	}
	if len(world.Cells) != height {
		t.Errorf("expected height: %d, actual number of rows: %d", height, len(world.Cells))
	}
	for i, row := range world.Cells {
		if len(row) != width {
			t.Errorf("expected width: %d, actual row's %d len: %d", width, i, world.Width)
		}
	}
}
