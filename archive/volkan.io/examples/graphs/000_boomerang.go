package main

import "fmt"

func distance(a, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]

	return dx*dx + dy*dy
}

func numberOfBoomerangs(points [][]int) int {
	distances := make(map[int]int, 0)
	result := 0

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i != j {
				d :=  distance(points[i], points[j])

				_, ok := distances[d]

				if ok {
					distances[d] += 1
				} else {
					distances[d] = 1
				}
			}
		}

		for _, val := range distances {
			result += val * (val - 1)
		}

		distances = make(map[int]int, 0)
	}

	return result
}

func main() {
	fmt.Println(
		numberOfBoomerangs([][]int{
			[]int{0,0},
			[]int{1,0},
			[]int{-1,0},
			[]int{0,1},
			[]int{0,-1},
		}),
	)
}
