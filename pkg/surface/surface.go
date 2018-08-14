package surface

import "fmt"

type Surface struct {
}

func (s Surface) String() string {
	return fmt.Sprint("Surface")
}
