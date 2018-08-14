package surface

import "fmt"

type Surface struct {
	polygons Polygons
}

func (surface Surface) String() string {
	return fmt.Sprintf("%s", surface.polygons)
}
