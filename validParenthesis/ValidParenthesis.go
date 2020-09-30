package main

func ValidParenthesis(parens string) bool {
	count := 0
	for _, c := range parens {
		if string(c) == "(" {
			count++
		}
		if string(c) == ")" {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

func main() {
	ValidParenthesis("()()()")
}
