package encrypt

func Nimbus(input string) (output string) {

	for _, c := range(input) {
		asciiCode := int(c) + 3
		output += string(asciiCode)
	}
	return
}