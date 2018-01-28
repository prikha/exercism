package triangle

import "math"

const testVersion = 3

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

// Pick values for the following identifiers used by the test program.
const NaT Kind = Kind("NaT") // not a triangle
const Equ Kind = Kind("Equ") // equilateral
const Iso Kind = Kind("Iso") // isosceles
const Sca Kind = Kind("Sca") // scalene


func KindFromSides(a, b, c float64) Kind {
  if isNat(a, b, c) {
    return NaT
  }
  if isEqu(a, b, c) {
    return Equ
  }
  if isIso(a, b, c) {
    return Iso
  }
  return Sca
}

func isNat(a, b, c float64) bool {
  sides := []float64{a, b, c}
  anyNaN := Any(sides, nanCheck)
  anyInf := Any(sides, infCheck)
  anyNegative := Any(sides, negativeCheck)
  wrongSides := natSides(a, b, c)
  return anyNaN || anyInf || anyNegative || wrongSides
}

func isEqu(a, b, c float64) bool {
  return a == b && b == c
}

func isIso(a, b, c float64) bool {
  return (a == b || a == c || b == c)
}

func natSides(a, b, c float64) bool {
  return (a + b < c || c + a < b || b + c < a)
}

type floatChecker func(float64) bool

func Any(in []float64, checker floatChecker) (out bool) {
  out = false
  for _, el := range in {
    out = checker(el)
    if out {
      break
    }
  }
  return
}

func nanCheck(in float64) bool {
  return math.IsNaN(in)
}

func infCheck(in float64) bool {
  return math.IsInf(in, 1) || math.IsInf(in, -1)
}

func negativeCheck(in float64) bool {
  return (in <= 0)
}
