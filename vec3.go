package vec3

import "fmt"

type Vec3 struct {
	X, Y, Z float64
}

func (v *Vec3) ToString() string {
	return fmt.Sprintf("Vec3 struct: %.2f, %.2f, %.2f", v.X, v.Y, v.Z)
}
