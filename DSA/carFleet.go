import (
	"slices"
)

type car struct {
	pos  int
	time float64
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}

	cars := make([]car, n)
	for i := 0; i < n; i++ {
		time := float64(target-position[i]) / float64(speed[i])
		cars[i] = car{pos: position[i], time: time}
	}

	slices.SortFunc(cars, func(a, b car) int {
		return b.pos - a.pos
	})

	fleets := 0
	var maxTime float64

	for _, c := range cars {
		if c.time > maxTime {
			fleets++
			maxTime = c.time
		}
	}

	return fleets
}
