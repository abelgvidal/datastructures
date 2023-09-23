package unionfind

var QUWelements = [7]int{0,1,2,3,4,5,6}
var QUWsz = [7]int{1,1,1,1,1,1,1}

func QUWRoot(p int) int {
	for p != QUWelements[p] {
		p = QUWelements[p]
	}
	return p
}

func QUWUnion(p, q int) {
	i := QUWRoot(p)
	j := QUWRoot(q)
	if i == j { return }
	if QUWsz[i] < QUWsz[j] {
		QUWelements[i] = j
		QUWsz[j] = QUWsz[j] + QUWsz[i]
	} else {
		QUWelements[j] = i
		QUWsz[i] = QUWsz[i] + QUWsz[j]
	}
}

func QUWConnected(p, q int) bool {
	return QUWRoot(p) == QUWRoot(q)
}
