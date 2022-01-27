package main

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func funcs() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)

	val := 10
	double(val)
	doublePtr(&val)
}
