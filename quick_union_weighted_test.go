package datastructures

import (
	"testing"
)

func TestQuickUnionWeighted(t *testing.T) {

	QUWUnion(3,4)

	QUWUnion(1,2)
	QUWUnion(6,2)
	QUWUnion(3,2)

	a := 3
	b := 2
	if !QUWConnected(a, b) {
		t.Errorf("error, %v, %v expected to be connected", a, b)
	}

	a = 6
	b = 2
	if !QUWConnected(a, b) {
		t.Errorf("error, %v, %v expected to be connected", a, b)
	}

	a = 0
	b = 4
	if QUWConnected(a, b) {
		t.Errorf("error, %v, %v expected to not be connected", a, b)
	}


}


func TestQuickUnionWeightedRoot(t *testing.T) {

	QUWUnion(1, 4)

	c := 4
	d := 1
	if !(QUWRoot(c) == d) {
		t.Errorf("error, %v, expected to be root of %v and its %v", c, d, QUWRoot(c))
	}
}
