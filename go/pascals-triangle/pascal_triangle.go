package pascal

const testVersion = 1

func Triangle(n int)(triangle [][]int) {
	for i := 0; i < n; i++ {
		triangle = append(triangle, []int{})
		for c := 0 ; c <= i; c++ {
			el := Factorial(i)/(Factorial(c)*Factorial(i-c))
			triangle[i] = append(triangle[i], el)
		}
	}
	return
}

func Factorial(n int)(result int) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
