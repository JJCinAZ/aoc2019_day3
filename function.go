// function.go
package function

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type intersection struct {
	wirebits    int
	wirelengths [2]int
}

var (
	pointMap map[int64]intersection
)

func Part1(w http.ResponseWriter, r *http.Request) {
	wires := make([]string, 0, 2)
	scanner := bufio.NewScanner(r.Body)
	for scanner.Scan() {
		wires = append(wires, scanner.Text())
	}
	if len(wires) != 2 {
		http.Error(w, "Missing input, expect two wire paths", http.StatusNotAcceptable)
		return
	}
	initMap()
	buildMap(parseLine(wires[0]), 1)
	buildMap(parseLine(wires[1]), 2)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("Closest is %d", getClosetIntersection())))
}

func Part2(w http.ResponseWriter, r *http.Request) {
	wires := make([]string, 0, 2)
	scanner := bufio.NewScanner(r.Body)
	for scanner.Scan() {
		wires = append(wires, scanner.Text())
	}
	if len(wires) != 2 {
		http.Error(w, "Missing input, expect two wire paths", http.StatusNotAcceptable)
		return
	}
	initMap()
	buildMap(parseLine(wires[0]), 1)
	buildMap(parseLine(wires[1]), 2)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("Shortest is %d", getShortestIntersection())))
}

func getClosetIntersection() int {
	shortest := int32(0)
	for p, intSect := range pointMap {
		if intSect.wirebits == 3 && p != 0 {
			var x, y int32
			x = int32(p >> 32)
			y = int32(p & 0xFFFFFFFF)
			d := abs32(x) + abs32(y)
			if shortest == 0 {
				shortest = d
			} else if d < shortest {
				shortest = d
			}
		}
	}
	return int(shortest)
}

func getShortestIntersection() int {
	shortest := 0
	for p, intSect := range pointMap {
		if intSect.wirebits == 3 && p != 0 {
			d := intSect.wirelengths[0] + intSect.wirelengths[1]
			if shortest == 0 {
				shortest = d
			} else if d < shortest {
				shortest = d
			}
		}
	}
	return shortest
}

func abs32(i int32) int32 {
	if i < 0 {
		return -i
	}
	return i
}
func parseLine(input string) []string {
	a := strings.Split(input, ",")
	return a
}

func initMap() {
	pointMap = make(map[int64]intersection)
}

func buildMap(line []string, wirebit int) error {
	var (
		curX, curY int32
		dist       int
	)
	for _, m := range line {
		count, err := strconv.Atoi(m[1:])
		if err != nil {
			return err
		}
		switch m[0] {
		case 'R':
			for x := 0; x < count; x++ {
				markPoint(curX, curY, wirebit, dist)
				curX++
				dist++
			}
		case 'L':
			for x := 0; x < count; x++ {
				markPoint(curX, curY, wirebit, dist)
				curX--
				dist++
			}
		case 'U':
			for y := 0; y < count; y++ {
				markPoint(curX, curY, wirebit, dist)
				curY++
				dist++
			}
		case 'D':
			for y := 0; y < count; y++ {
				markPoint(curX, curY, wirebit, dist)
				curY--
				dist++
			}
		}
	}
	return nil
}

func markPoint(x, y int32, wirebit int, dist int) {
	lenIdx := 0
	if wirebit == 2 {
		lenIdx = 1
	}
	u := int64(x)<<32 | (int64(y) & 0xFFFFFFFF)
	if intSect, found := pointMap[u]; found {
		intSect.wirebits |= wirebit
		if intSect.wirelengths[lenIdx] == 0 {
			intSect.wirelengths[lenIdx] = dist
		}
		pointMap[u] = intSect
	} else {
		intSect = intersection{wirebits: wirebit}
		intSect.wirelengths[lenIdx] = dist
		pointMap[u] = intSect
	}
}
