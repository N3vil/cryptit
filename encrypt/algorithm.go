// This package provides simple text encryption 
package encrypt

// The Nimbus function encrypts a string by adding 4 to its ascii value
func Nimbus(input string) (output string) {

	for _, c := range(input) {
		asciiCode := int(c) + 5
		output += string(asciiCode)
	}
	return
}