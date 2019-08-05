package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

const result = "Polar radius=%.02f θ=%.02f° → Cartesian x=%.02f y=%.02f\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " +
	"or %s to quit."

type polar struct {
	radius float64
	theta float64
}

type cartesian struct {
	x float64
	y float64
}

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)

	go func() {
		for {
			polarCoord := <-questions
			theta := polarCoord.theta * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(theta)
			y := polarCoord.radius * math.Sin(theta)

			answers <- cartesian{x, y}
		}
	}()

	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, theta float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &theta); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, theta}
		coord := <-answers
		fmt.Printf(result, radius, theta, coord.x, coord.y)
	}
	fmt.Println()

}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}
