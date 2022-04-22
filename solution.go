package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SideNum int

const (
	SidesTriangle SideNum = 3
	SidesSquare   SideNum = 4
	SidesCircle   SideNum = 0
)

func CalcSquare(sideLen float64, sidesNum SideNum) float64 {
	if sideLen < 0 {
		return 0
	}
	switch sidesNum {
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	case SidesTriangle:
		return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0
	}
}
