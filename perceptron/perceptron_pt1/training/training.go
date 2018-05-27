package training

import (
	"math/rand"

	"github.com/fogleman/gg"
)

type Point struct {
	x     int64
	y     int64
	label int
}

type Training struct {
	points []Point
	dd     *gg.Context
}

func New(size int) Training {
	var newTraining Training
	newTraining.dd = gg.NewContext(500, 500)
	newTraining.dd.DrawLine(0, 0, 500, 500)
	newTraining.dd.SetRGB255(255, 255, 255)
	newTraining.dd.Fill()

	newTraining.points = make([]Point, size)

	for i := 0; i < size; i++ {
		newTraining.points[i].x = int64(rand.Intn(500))
		newTraining.points[i].y = int64(rand.Intn(500))
		if newTraining.points[i].x > newTraining.points[i].y {
			newTraining.points[i].label = 1
		} else {
			newTraining.points[i].label = -1
		}
	}
	return newTraining
}

func (currentTraining *Training) Show() {

	for i := range currentTraining.points {
		x := float64(currentTraining.points[i].x)
		y := float64(currentTraining.points[i].y)
		label := currentTraining.points[i].label

		currentTraining.dd.DrawEllipse(x, y, 8.0, 8.0)
		if label == 1 {
			currentTraining.dd.SetRGB255(255, 255, 255)
		} else {
			currentTraining.dd.SetRGB255(0, 0, 0)
		}
		currentTraining.dd.Fill()
	}
	currentTraining.dd.SavePNG("out.png")
}
