package unionfind

import (
	"testing"
)

func TestQuickUnionWeightedCompression(t *testing.T) {

	QUWCUnion(3,4)

	QUWCUnion(1,2)
	QUWCUnion(6,2)
	QUWCUnion(3,2)

	a := 3
	b := 2
	if !QUWCConnected(a, b) {
		t.Errorf("error, %v, %v expected to be connected", a, b)
	}

	a = 6
	b = 2
	if !QUWCConnected(a, b) {
		t.Errorf("error, %v, %v expected to be connected", a, b)
	}

	a = 0
	b = 4
	if QUWCConnected(a, b) {
		t.Errorf("error, %v, %v expected to not be connected", a, b)
	}


}


func TestQuickUnionWeightedCompresionRoot(t *testing.T) {

	QUWCUnion(1, 4)

	c := 4
	d := 1
	if !(QUWCRoot(c) == d) {
		t.Errorf("error, %v, expected to be root of %v and its %v", c, d, QUWCRoot(c))
	}
}
