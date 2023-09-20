package datastructures

var QUWCelements = [7]int{0,1,2,3,4,5,6}
var QUWCsz = [7]int{1,1,1,1,1,1,1}

func QUWCRoot(p int) int {
	for p != QUWCelements[p] {
		QUWCelements[p] = QUWCelements[QUWCelements[p]] // path compression
		p = QUWCelements[p]
	}

	return p
}

func QUWCUnion(p, q int) {
	i := QUWCRoot(p)
	j := QUWCRoot(q)
	if i == j { return }
	if QUWCsz[i] < QUWCsz[j] {
		QUWCelements[i] = j
		QUWCsz[j] = QUWCsz[j] + QUWCsz[i]
	} else {
		QUWCelements[j] = i
		QUWCsz[i] = QUWCsz[i] + QUWCsz[j]
	}
}

func QUWCConnected(p, q int) bool {
	return QUWCRoot(p) == QUWCRoot(q)
}
