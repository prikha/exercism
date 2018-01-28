package secret

const testVersion = 1

var words = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(in uint) []string {
	stack := ToBinaryStack(in)
	reverse := false
	var outStack []string
	for idx, flag := range stack {
		switch {
		case idx > 4:
			continue
		case idx == 4:
			reverse = true
		default:
			// idx < 4
			if flag == 1 {
				outStack = append(outStack, words[idx])
			}
		}
	}
	outStack = CondReverse(outStack, reverse)
	return outStack
}

func Reverse(in []string) []string {
	var out = make([]string, len(in))
	for idx, el := range in {
		out[len(in)-idx-1] = el
	}
	return out
}

func CondReverse(arr []string, reverse bool) []string {
	if reverse {
		return Reverse(arr)
	} else {
		return arr
	}
}

func ToBinaryStack(in uint) (stack []int) {
	n := int(in)
	for n > 0 {
		stack = append(stack, n%2)
		n = n / 2
	}
	return
}
