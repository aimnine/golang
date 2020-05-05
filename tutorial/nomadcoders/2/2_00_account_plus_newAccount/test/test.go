package test

type test struct {
	a string
	b int
}

// Test function
func Test(a string) *test {
	t := test{a:a, b:0}
	return &t
}