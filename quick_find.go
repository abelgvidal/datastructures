package datastructures

var elements = [7]int{0,1,2,3,4,5,6}

func QFConnected(p, q int) bool {
	return elements[p] == elements[q]
}

func QFUnion(p, q int) {
	positionP := elements[p]
	positionQ := elements[q]
	for x := 0; x < len(elements); x++ {
		if elements[x] == positionP {
			elements[x] = positionQ
		}
	}
}

func QFElements() [7]int {
	return elements
}
