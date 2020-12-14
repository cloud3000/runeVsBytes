package main

/*
A few examples; how to process text or character data in Golang.
In Go we have strings, bytes, and runes to process text or character data.
Each data type has it's place,
	these examples show which data type is best for addressing a single character.
	Also, how Go solved the problem that has existed for decades with character data.

In the beginning each hardware vendor had their own method to address character data.
Therefore they could not communicate with oneanother.
Then IBM came up with EBCDIC which stands for the Extended Binary Coded Decimal Interchange.
Non-IBM vendors came up with ASCII which stands for American Standard Code for Information Interchange.
With standard functions developed to convert between EBCDIC & ASCII everyone could communicate.
That worked from the 1960's through the 1995, because EBCDIC & ASCII only addressed the english alphabet.

Microsoft invented it's own methods to address international character sets,
 similar to IBM this was also a proprietary solution.

However, smarter minds decided on UTF8 (unicode).
But unlike ASCII, many UTF8 characters required more than one byte per character.
Programming languages didn't support multi-byte characters, and had to be patched for UTF8.

Fortunitely one of the co-developers of Go, before developing GO, was the creator of UTF8.
So UTF8 was build into Go from the get-go, using the Go data type 'rune'.

Lets test this. ;-)

*/
import (
	"fmt"
	"os"
)

func main() {
	str1 := "Hello gophers"
	str2 := "你好地鼠"
	str3 := "Nǐ hǎo dì shǔ"

	// You can pass up to three string on the command line.
	if len(os.Args) > 1 {
		str1 = os.Args[1]
	}
	if len(os.Args) > 2 {
		str2 = os.Args[2]
	}
	if len(os.Args) > 3 {
		str3 = os.Args[3]
	}

	fmt.Printf("\n\n\rCharacter String to process: [%s] [%s] [%s]\n\r\n", str1, str2, str3)

	evalRunehex(revstr2(str1))
	evalRunehex(str1)
	evalRunehex(revstr2(str2))
	evalRunehex(str2)
	evalRunehex(revstr2(str3))
	evalRunehex(str3)

}

// evalString will iterate the length of a string.
// It will also show:
// 	 each character and it's index position,
// 	 the total count of iterations,
//   the complete string value.
func evalString(str string) {
	fmt.Println("************** evalString **************")
	sentence := str
	counter := 0

	for index, letter := range sentence {
		counter++
		fmt.Printf("Index: %v letter: %c\n", index, letter)
	}

	fmt.Printf("      Counter value is: %v\n", counter)
	fmt.Println("Length of sentence is: ", len(sentence))
	fmt.Println(sentence)

}

// evalByte is like evalString, except with a slice of bytes instead of a string.
// NOTE: The difference in the output.
func evalByte(str string) {
	fmt.Println("************** evalByte **************")
	sentence := []byte(str)
	counter := 0

	for index, letter := range sentence {
		counter++
		fmt.Printf("Index: %v letter: %c\n", index, letter)
	}

	fmt.Printf("      Counter value is: %v\n", counter)
	fmt.Println("Length of sentence is: ", len(sentence))
	fmt.Println(sentence)

}

// evalRune is like evalByte, except it uses runes instead of bytes.
// NOTE: The difference in the output.
func evalRune(str string) {
	fmt.Println("************** evalRune **************")
	sentence := []rune(str)
	counter := 0

	for index, letter := range sentence {
		counter++
		fmt.Printf("Index: %v letter: %c\n", index, letter)
	}

	fmt.Printf("Counter value: %v\n", counter)
	fmt.Printf("%x \n", sentence)
	fmt.Println("Length of sentence is: ", len(sentence))

}

// evalBytehex is like evalByte, except it show the hex values.
// ALSO: We will not show all index positions, rather focus on the order of bytes.
func evalBytehex(str string) {
	fmt.Println("************** evalBytehex **************")
	fmt.Println(str)
	sentence := []byte(str)
	counter := 0
	for _, letter := range sentence {
		counter++
		fmt.Printf("%x ", letter)
	}
	fmt.Println(" ")

	fmt.Printf("        Counter value: %v\n", counter)
	fmt.Printf("Length of sentence is: %v\n", len(sentence))
}

// evalRunehex is like evalRune, except it show the hex values.
// ALSO: We will not show all index positions, rather focus on the order of runes.
func evalRunehex(str string) {
	fmt.Println("************** evalRunehex **************")
	fmt.Println(str)
	sentence := []rune(str)
	counter := 0
	for _, letter := range sentence {
		counter++
		fmt.Printf("%x ", letter)
	}
	fmt.Println(" ")

	fmt.Printf("        Counter value: %v\n", counter)
	fmt.Printf("Length of sentence is: %v\n", len(sentence))
}

// revstr1 Is the wrong way to reverse a string.
// It will work correctly within the ASCII character range.
// IT WILL FAIL using UTF8 characters that require more than 1 byte per characters.
func revstr1(str string) string {
	fmt.Println("************** R E V S T R 1 **************")
	fmt.Println(str)
	theString := []byte(str)
	for i, j := 0, len(theString); i < j; i++ {
		j--
		theString[i], theString[j] = theString[j], theString[i]
	}
	return string(theString)
}

// revstr2 is the correct way to reverse a string, using a slice of runes.
func revstr2(str string) string {
	fmt.Println("************** R E V S T R 2 **************")
	fmt.Println(str)
	theString := []rune(str)
	for i, j := 0, len(theString); i < j; i++ {
		j--
		theString[i], theString[j] = theString[j], theString[i]
	}
	return string(theString)
}

// Will revstr3 be another way to reverse a string, using a simple string?
// Really, WHY USE A SLICE OF RUNES? Lets see..... uncomment and try it!
/*
func revstr3(str string) string {
	fmt.Println("************** R E V S T R 2 **************")
	fmt.Println(str)
	theString := str
	for i, j := 0, len(theString); i < j; i++ {
		j--
		theString[i], theString[j] = theString[j], theString[i]
	}
	return theString
}
*/
