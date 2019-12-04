package function

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func formatIntersections() string {
	a := make([]string, 0)
	for p, intSect := range pointMap {
		if intSect.wirebits == 3 && p != 0 {
			var x, y int32
			x = int32(p >> 32)
			y = int32(p & 0xFFFFFFFF)
			a = append(a, fmt.Sprintf("%d,%d", x, y))
		}
	}
	sort.Strings(a)
	return strings.Join(a, ";")
}

func Test_buildMap1(t *testing.T) {
	type args struct {
		wire1, wire2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"R8,U5,L5,D3", "U7,R6,D4,L4"}, "3,3;6,5"},
		{"test2", args{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, "146,46;155,11;155,4;158,-12"},
		{"test3", args{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, "107,47;107,51;107,71;124,11;157,18"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMap()
			buildMap(parseLine(tt.args.wire1), 1)
			buildMap(parseLine(tt.args.wire2), 2)
			if got := formatIntersections(); got != tt.want {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildMap2(t *testing.T) {
	type args struct {
		wire1, wire2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 30},
		{"test2", args{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 610},
		{"test3", args{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 410},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMap()
			buildMap(parseLine(tt.args.wire1), 1)
			buildMap(parseLine(tt.args.wire2), 2)
			if got := getShortestIntersection(); got != tt.want {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
