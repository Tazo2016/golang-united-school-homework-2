package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesNumType int

const (
	SidesTriangle sidesNumType = 3
	SidesSquare   sidesNumType = 4
	SidesCircle   sidesNumType = 0
)

func CalcSquare(sideLen float64, sidesNum sidesNumType) float64 {
	if sideLen > 0 {
		switch sidesNum {
		case SidesTriangle:
			return (math.Sqrt(3.0) / 4) * math.Pow(sideLen, 2.0)
		case SidesSquare:
			return math.Pow(sideLen, 2.0)
		case SidesCircle:
			return math.Pi * math.Pow(sideLen, 2.0)
		default:
			return 0
		}
	} else {
		return 0
	}
}
