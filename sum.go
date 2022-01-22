package sum

func Sum(xs ...int) int {
	s := 0
	for _, v := range xs {
		s += v
	}
	return s
}
