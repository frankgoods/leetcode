package problems

func addBinary(a string, b string) string {
	var res []byte
	var other string
	if len(a) > len(b) {
		res = []byte(a)
		other = b
	} else {
		res = []byte(b)
		other = a
	}
	l := len(res)

	carry := 0
	i := l - 1
	j := len(other) - 1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		v1 := int(res[i] - '0')
		v2 := int(other[j] - '0')
		sum := v1 + v2 + carry
		carry = sum / 2
		res[i] = byte('0' + sum%2)
	}

	for ; carry == 1 && i >= 0; i-- {
		v := int(res[i] - '0')
		sum := v + carry
		carry = sum / 2
		res[i] = byte('0' + sum%2)
	}

	if carry == 1 {
		return "1" + string(res)
	} else {
		return string(res)
	}
}
