package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sides int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const (
	SidesTriangle sides = 3
	SidesSquare   sides = 4
	SidesCircle   sides = 0
)

// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	if sidesNum == SidesTriangle {
		result := (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
		return result
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	} else if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	} else {
		return 0
	}
}
