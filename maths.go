package ske

import "math"

type Transform struct {
	P Vec
	R Vec
}

type Vec struct {
	X, Y, Z, W float64
}

func V2(x, y float64) Vec {
	return Vec{x, y, 0, 0}
}

func V3(x, y, z float64) Vec {
	return Vec{x, y, z, 0}
}

func V4(x, y, z, w float64) Vec {
	return Vec{x, y, z, w}
}

func (v Vec) Add(other interface{}) Vec {
	switch o := other.(type) {
	case float64:
		return Vec{v.X + o, v.Y + o, v.Z + o, +v.W + o}
	case Vec:
		return Vec{v.X + o.X, v.Y + o.Y, v.Z + o.Z, v.W + o.W}
	}
	return Vec{}
}

func (v Vec) Sub(other interface{}) Vec {
	switch o := other.(type) {
	case float64:
		return Vec{v.X - o, v.Y - o, v.Z - o, v.W - o}
	case Vec:
		return Vec{v.X - o.X, v.Y - o.Y, v.Z - o.Z, v.W - o.W}
	}
	return Vec{}
}

func (v Vec) Mul(other interface{}) Vec {
	switch o := other.(type) {
	case float64:
		return Vec{v.X * o, v.Y * o, v.Z * o, v.W * o}
	case Vec:
		return Vec{v.X * o.X, v.Y * o.Y, v.Z * o.Z, v.W * o.W}
	}
	return Vec{}
}

func (v Vec) Div(other interface{}) Vec {

	switch o := other.(type) {
	case float64:
		if o != 0 {
			return Vec{v.X / o, v.Y / o, v.Z / o, v.W / o}
		}
		return v
	case int:
		if o != 0 {
			return Vec{v.X / float64(o), v.Y / float64(o), v.Z / float64(o), v.W / float64(o)}
		}
		return v
	case Vec:
		return Vec{v.X / o.X, v.Y / o.Y, v.Z / o.Z, v.W / o.W}
	}
	Assert(false, "cannot div, unknown interface{} type")
	return Vec{}
}

func (v Vec) Mag() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z) + (v.W * v.W))
}

func (v Vec) Normalize() Vec {
	if m := v.Mag(); m != 0 {
		return Vec{v.X / m, v.Y / m, v.Z / m, v.W / m}
	} else {
		return v
	}
}
func Dir(v1, v2 Vec) Vec {
	vec := Vec{v2.X - v1.X, v2.Y - v1.Y, v2.Z - v1.Z, v2.W - v1.W}
	return vec.Normalize()
}

func Lerp(v1, v2 Vec, intensity float64) Vec {
	return (v1.Mul(intensity)).Add(v2.Mul(1 - intensity))
}

func (v Vec) Equals(other Vec) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z && v.W == other.W
}
