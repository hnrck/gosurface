package surface

import "fmt"

type TwoDimPoint struct {
	x, y float64
}

func TwoDimPointOrigin() TwoDimPoint {
	return TwoDimPoint{0.0, 0.0}
}

func (point TwoDimPoint) String() string {
	return fmt.Sprintf("%g,%g", point.x, point.y)
}

type ThreeDimPoint struct {
	TwoDimPoint
	z float64
}

func ThreeDimPointOrigin() ThreeDimPoint {
	return ThreeDimPoint{TwoDimPointOrigin(), 0.0}
}

func (point ThreeDimPoint) String() string {
	return fmt.Sprintf("%g,%g,%g", point.x, point.y, point.z)
}
