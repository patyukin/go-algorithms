package main

import (
	"fmt"
	"math"
)

type polar struct {
	radius float64
	angle  float64
}

type cartesian struct {
	x, y float64
}

func main() {
	questions := make(chan polar)
	answers := createSolver(questions)
	defer close(questions)
	interact(questions, answers)
	defer close(answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		p := <-questions
		phi := p.angle * math.Pi / 180.0
		x := p.radius * math.Cos(phi)
		y := p.radius * math.Sin(phi)
		answers <- cartesian{x, y}
	}()
	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	var radius, angle float64
	fmt.Print("radius: ")
	fmt.Scan(&radius)

	fmt.Print("angle: ")
	fmt.Scan(&angle)

	polar := polar{radius: radius, angle: angle}
	questions <- polar

	cartesianCoord := <-answers
	fmt.Printf("Cartesian coordinates (x, y): (%f, %f)\n", cartesianCoord.x, cartesianCoord.y)
}
