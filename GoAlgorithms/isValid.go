package main

func isValid(s string) bool {
	//ss := "()[]{}"40 41 91 93 123 125
	bs := []byte(s)
	if (bs[0] != 40) && (bs[0] != 91) && (bs[0] != 123) {
		return false
	}
	l := len(s)
	if l%2 != 0 {
		return false
	}
	valid := 0
	for i := 0; i <= l/2; i = i + 2 {
		if (bs[i] == bs[i+1]-1) || (bs[i] == bs[i+1]-2) {
			valid++
		}
	}
	if valid == l/2 {
		return true
	}
	return false
}
func main() {
	isValid("{[]}")
}
