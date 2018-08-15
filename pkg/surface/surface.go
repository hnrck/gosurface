package surface

import (
	"fmt"
	"io"
	"math"
)

type Surface struct {
	Width, Height uint
	Cells         uint
	XYRange       float64
	Angle         float64
	polygons      Polygons
}

func (surface *Surface) GeneratePolygons() {
	for i := 0; i < int(surface.Cells); i++ {
		for j := 0; j < int(surface.Cells); j++ {
			polygon := surface.cell(TwoDimPoint{float64(i), float64(j)})
			surface.appendPolygon(polygon)
		}
	}
}

func (surface Surface) cell(point TwoDimPoint) Polygon {
	threeDimPoints := make([]ThreeDimPoint, 4)
	twoDimPoints := make([]TwoDimPoint, 4)
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			threeDimPoints[2*i+j] = surface.computeCurve(TwoDimPoint{
				point.x + float64(i),
				point.y + float64(j),
			})
		}
	}
	for i, p := range threeDimPoints {
		twoDimPoints[i] = p.isoprojection(surface.Width, surface.Height, surface.XYRange, surface.Angle)
	}
	return Polygon{
		a: twoDimPoints[2],
		b: twoDimPoints[0],
		c: twoDimPoints[1],
		d: twoDimPoints[3],
	}
}

func (surface Surface) computeCurve(point TwoDimPoint) ThreeDimPoint {
	tmpPoint := TwoDimPoint{
		surface.XYRange * (float64(point.x)/float64(surface.Cells) - 0.5),
		surface.XYRange * (float64(point.y)/float64(surface.Cells) - 0.5),
	}
	r := math.Hypot(tmpPoint.x, tmpPoint.y)
	return ThreeDimPoint{tmpPoint, math.Sin(r) / r}
}

func (surface *Surface) appendPolygon(polygon Polygon) {
	surface.polygons.appendPolygon(polygon)
}

func (surface Surface) Writer(w io.Writer) {
	fmt.Fprintf(w, "%s", surface)
}

func (surface Surface) String() string {
	return fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>\n%s</svg>", surface.Width, surface.Height, surface.polygons)
}
