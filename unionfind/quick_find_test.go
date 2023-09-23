package unionfind

import (
	"testing"
)

func TestQuickFind(t *testing.T) {

	QFUnion(3,4)

	if !QFConnected(3,4) {
		t.Errorf("error, 3, 4 expected to be connected")
	}

	if QFConnected(3,5) {
		t.Errorf("error, 3, 5 expected to not be connected")
	}
}
