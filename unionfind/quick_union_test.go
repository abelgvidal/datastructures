package unionfind

import (
	"testing"
)

func TestQuickUnion(t *testing.T) {


	QUUnion(3,4)

	a := 3
	b := 4
	if !QUConnected(a, b) {
		t.Errorf("error, %v, %v expected to be connected", a, b)
	}

	a = 2
	b = 4
	if QUConnected(a, b) {
		t.Errorf("error, %v, %v expected to not be connected", a, b)
	}


}


func TestQuickUnionRoot(t *testing.T) {

	QUUnion(1, 4)

	c := 1
	d := 4
	if !(QURoot(c) == d) {
		t.Errorf("error, %v, expected to be root of %v", c, d)
	}
}
