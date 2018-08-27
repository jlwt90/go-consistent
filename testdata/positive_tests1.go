package tests1

// In this test suite, (1) option is always preferred.

// T is an example type.
type T struct {
	integer int
}

func zeroValuePointerAllocation() {
	_ = new(T)
	_ = new(T)
	//= zeroValPtrAlloc: use new instead of address-of-literal
	_ = &T{}
}

func emptySlice() {
	_ = make([]int, 0)
	_ = make([]float64, 0)
	//= emptySlice: use make instead of literal
	_ = []string{}
}

func nilSlice() {
	var s1 []int
	var s2 []*T
	//= nilSliceDecl: use var instead of literal
	s3 := []string(nil)
}

func emptyMap() {
	_ = make(map[T]T)
	_ = make(map[*T]*T)
	//= emptyMap: use make instead of literal
	_ = map[int]int{}
}
