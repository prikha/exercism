package series

const testVersion = 2

func All(n int, s string)(out []string){
	for i := 0; (i + n) <= len(s); i++ {
		out = append(out, s[i:(i+n)])
	}
	return
}

func UnsafeFirst(n int, s string)(out string){
	out = s[0:n]
	return
}

func First(n int, s string)(first string, ok bool){
	if n > len(s) {
		ok = false
		return
	} else {
		ok = true
	}
	first = UnsafeFirst(n, s)
	return
}
