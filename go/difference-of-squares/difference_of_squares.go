package diffsquares

const testVersion = 1

// (1 + 2 + ... + 10)² = 55² = 3025.
func SquareOfSums(n int) int {
  out := 0
  for i := 1; i<= n; i++ {
    out += i
  }
  return out * out
}

// 1² + 2² + ... + 10² = 385.
func SumOfSquares(n int) (out int) {
  for i := 1; i<= n; i++ {
    out += i*i
  }
  return
}

func Difference(n int) int {
  return SquareOfSums(n) - SumOfSquares(n)
}
