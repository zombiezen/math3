package quat32

import (
	"bitbucket.org/zombiezen/math3/vec32"
	"math"
	"testing"
)

const checkTol = 0.01

func TestAxisAngle(t *testing.T) {
	tests := []struct {
		Axis  vec32.Vector
		Angle float32
		Q     Quaternion
	}{
		{vec32.Vector{0, 0, 1}, math.Pi / 2, Quaternion{0, 0, -float32(1 / math.Sqrt(2)), float32(1 / math.Sqrt(2))}},
	}
	for _, test := range tests {
		q := AxisAngle(test.Axis, test.Angle)
		if !checkQuaternion(q, test.Q, checkTol) {
			t.Errorf("AxisAngle(%v, %v) = %v; want %v", test.Axis, test.Angle, q, test.Q)
		}
	}
}

func TestTransform(t *testing.T) {
	tests := []struct {
		Q   Quaternion
		V   vec32.Vector
		Out vec32.Vector
	}{
		{AxisAngle(vec32.Vector{0, 0, 1}, -math.Pi/2), vec32.Vector{2, 0, 0}, vec32.Vector{0, 2, 0}},
	}
	for _, test := range tests {
		out := test.Q.Transform(test.V)
		if !checkVector(out, test.Out, checkTol) {
			t.Errorf("%v.Transform(%v) = %v; want %v", test.Q, test.V, out, test.Out)
		}
	}
}

// checkVector returns whether v1 ~ v2, given a tolerance.
func checkVector(v1, v2 vec32.Vector, tol float32) bool {
	for i := range v1 {
		if v2[i] > v1[i]+tol || v2[i] < v1[i]-tol {
			return false
		}
	}
	return true
}

// checkQuaternion returns whether q1 ~ q2, given a tolerance.
func checkQuaternion(q1, q2 Quaternion, tol float32) bool {
	for i := range q1 {
		if q2[i] > q1[i]+tol || q2[i] < q1[i]-tol {
			return false
		}
	}
	return true
}
