package unionfind

var QUWCelements = [20]int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19}
var QUWCsz = [20]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}

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
