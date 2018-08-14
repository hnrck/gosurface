package surface

import (
	"fmt"
	"strings"
)

type Polygons []Polygon

func (polygons Polygons) String() string {
	var polygonStrings []string
	for _, polygon := range polygons {
		polygonStrings = append(polygonStrings, polygon.String())
	}
	return strings.Join(polygonStrings, "\n")
}

type Polygon struct {
	a, b, c, d TwoDimPoint
	z          float64
}

func (polygon Polygon) String() string {
	return fmt.Sprintf("%s %s %s %s", polygon.a, polygon.b, polygon.c, polygon.d)
}
