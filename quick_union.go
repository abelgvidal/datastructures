package datastructures

import (
	"fmt"
)

var QUelements = [7]int{0,1,2,3,4,5,6}

func QURoot(p int) int {
	for p != QUelements[p] {
		p = QUelements[p]
	}
	return p
}

func QUUnion(p, q int) {
	i := QURoot(p)
	j := QURoot(q)
	QUelements[i] = j
}

func QUConnected(p, q int) bool {
	return QURoot(p) == QURoot(q)
}

func QUStatus() {
	fmt.Println(QUelements)
}
