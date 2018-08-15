package surface

import (
	"fmt"
	"math"
)

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

func (point ThreeDimPoint) isoprojection(width, height uint, xyrange, angle float64) TwoDimPoint {
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4
	return TwoDimPoint{
		float64(width)/2 + (point.x-point.y)*math.Cos(angle)*xyscale,
		float64(height)/2 + (point.x+point.y)*math.Sin(angle)*xyscale - point.z*zscale,
	}
}

func (point ThreeDimPoint) String() string {
	return fmt.Sprintf("%g,%g,%g", point.x, point.y, point.z)
}
