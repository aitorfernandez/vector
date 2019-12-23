package vec3

import (
	"fmt"
	"log"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v *Vec3) ToString() string {
	return fmt.Sprintf("Vec3 struct: %.2f, %.2f, %.2f", v.X, v.Y, v.Z)
}

func (v *Vec3) Add(a ...interface{}) {
	switch e := a[0].(type) {
	case Vec3:
		v.X += e.X
		v.Y += e.Y
		v.Z += e.Z
	case int, float64:
		for i, arg := range a {
			switch i {
			case 0:
				v.X += getFloat(arg)
			case 1:
				v.Y += getFloat(arg)
			case 2:
				v.Z += getFloat(arg)
			}
		}
	}
}

func (v *Vec3) Sub(a ...interface{}) {
	switch e := a[0].(type) {
	case Vec3:
		v.X -= e.X
		v.Y -= e.Y
		v.Z -= e.Z
	case int, float64:
		for i, arg := range a {
			switch i {
			case 0:
				v.X -= getFloat(arg)
			case 1:
				v.Y -= getFloat(arg)
			case 2:
				v.Z -= getFloat(arg)
			}
		}
	}
}

func (v *Vec3) Mult(n float64) {
	v.X *= n
	v.Y *= n
	v.Z *= n
}

func (v *Vec3) Div(n float64) {
	if n == 0 {
		log.Println("divide by 0")
		return
	}

	v.X /= n
	v.Y /= n
	v.Z /= n
}

func getFloat(i interface{}) float64 {
	switch i.(type) {
	case int:
		return float64(i.(int))
	default:
		return i.(float64)
	}
}
