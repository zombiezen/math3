// Package quat32 operates on float32 quaternions.
package quat32

import (
	"bitbucket.org/zombiezen/math3/vec32"
	"math"
)

// A Quaternion holds a four-dimensional quaternion.  The values are stored in {X, Y, Z, W} order.
type Quaternion [4]float32

// Length returns the norm of q.
func (q Quaternion) Length() float32 {
	return float32(math.Sqrt(float64(q.LengthSqr())))
}

// LengthSqr returns the norm squared of q.  This is cheaper to compute than Length.
func (q Quaternion) LengthSqr() float32 {
	return q[0]*q[0] + q[1]*q[1] + q[2]*q[2] + q[3]*q[3]
}

// Negate returns a new quaternion that applies an opposite rotation.
func (q Quaternion) Negate() Quaternion {
	return Quaternion{q[0], q[1], q[2], -q[3]}
}

// Conjugate returns q's conjugate.
func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{-q[0], -q[1], -q[2], q[3]}
}

// Transform computes the rotation q for point v.
// Mathematically, this is the product of q, v, and q's conjugate.
func (q Quaternion) Transform(v vec32.Vector) vec32.Vector {
	qq := Mul(Mul(q, Quaternion{v[0], v[1], v[2]}), q.Conjugate())
	return vec32.Vector{qq[0], qq[1], qq[2]}
}

// Mul calculates the product of q1 and q2.  The result corresponds to the rotation q2 followed by the rotation q1.
func Mul(q1, q2 Quaternion) Quaternion {
	v1, v2 := vec32.Vector(q1).Vec3(), vec32.Vector(q2).Vec3()
	v3 := vec32.Add(vec32.Add(v2.Scale(q1[3]), v1.Scale(q2[3])), vec32.Cross(v1, v2))
	return Quaternion{v3[0], v3[1], v3[2], q1[3]*q2[3] - vec32.Dot(v1, v2)}
}

// AxisAngle builds a new quaternion from a vector (which will be normalized)
// and an angle in radians.
func AxisAngle(axis vec32.Vector, angle float32) Quaternion {
	axis = axis.Normalize()
	sin, cos := float32(math.Sin(float64(angle/2))), float32(math.Cos(float64(angle/2)))
	return Quaternion{
		-axis[0] * sin,
		-axis[1] * sin,
		-axis[2] * sin,
		cos,
	}
}
