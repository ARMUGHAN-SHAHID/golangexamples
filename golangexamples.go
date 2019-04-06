package golangexamples

import (
	"fmt"

	"github.com/ehteshamz/greetings"
)

func ConcatSlice(sliceToConcat []byte) string {
	var concatString string
	for index, element := range sliceToConcat {
		concatString += string(element)
		if index < len(sliceToConcat)-1 {
			concatString += string('-')
		}

	}
	return concatString
}
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for index, element := range sliceToEncrypt {
		sliceToEncrypt[index] = element + byte(ceaserCount)
	}
}
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}

func Test() {
	temp := []byte{'a', 'r', 'm'}
	fmt.Println(ConcatSlice(temp))
	Encrypt(temp, 1)
	fmt.Println(string(temp))
	fmt.Println(EZGreetings("ARMUGHAN"))

}
