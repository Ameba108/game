package life

import (
	"errors"
	"math/rand"
	"time"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("Height and width must be positive numbers")
	}

	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) next(x, y int) bool {
	n := w.neighbors(x, y)
	alive := w.Cells[y][x]
	if n < 4 && n > 1 && alive {
		return true
	} else if n == 3 && !alive {
		return true
	} else {
		return false
	}
}

func (w *World) neighbors(x, y int) int {
	number := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			neighborX := (x + i + w.Width) % w.Width
			neighborY := (y + j + w.Height) % w.Height
			if w.Cells[neighborY][neighborX] {
				number++
			}
		}
	}
	return number
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.next(j, i)
		}
	}
}

func (w *World) RandInit(percentage int) {
	numAlive := percentage * w.Height * w.Width / 100
	w.fillAlive(numAlive)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
