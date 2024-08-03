package decrypt

func Nimbus(input string) (output string) {

	for _, c := range(input) {
		asciiCode := int(c) - 5
		output += string(asciiCode)
	}
	return
}