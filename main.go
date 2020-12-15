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
So UTF8 was build into Go from the get-go, using the Go data types string and 'rune'.

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
	str4 := "    The is a Test     "

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

	// Can we copy a string to a slice of runes?
	rune1 := []rune(str2 + str1)
	fmt.Println("yes we can:  str2+str1 to rune1", rune1)

	// Can we copy a string to a slice of bytes?
	byte1 := []byte(str2 + str1)
	fmt.Println("yes we can:  str2+str1 to byte1", byte1)

	// Can we copy a slice of runes to a slice of bytes?
	// Yes but, we can't directly copy a variable of type []rune) to a variable of type []byte
	// We need to run it through a string data type first, like so:
	byte2 := []byte(string(rune1))
	fmt.Println("yes we can: rune1 to byte2", byte2)

	// Can we copy a slice of bytes to a slice of runes?
	// Yes but, we can't directly copy a variable of type []byte) to a variable of type []rune
	// We need to run it through a string data type first, like so:
	rune2 := []rune(string(byte1))
	fmt.Println("yes we can: byte1 to rune2", rune2)

	b := byte(65)    //uint8
	r := rune(20320) //uint32
	fmt.Println(string(b))
	fmt.Println(string(r))

	// Let me explain: UTF-8 is the native/default character encoding scheme in GO.
	// UTF-8 character is encoded using a uint32 integer, containing 4 (8bit) bytes.
	// The low order six bits per byte represent the actual characters being encoded.
	// The high order two bits represent continuation info, for one or more 8 bit values.
	// Some UTF8 character only need 6 bits, other are much larger, they vary.

	// In the old days using ASCII, all chars were 8 bits, and are compatible with UTF8, because
	// UTF8 supports variable bit lengths using those two hi-order bits per byte. For this reason
	// You can convert a string to a []rune, and maintain all UTF8 data. However, when you convert
	// a string to a []byte you will lose data, or at least the meaning of the data. Likewise,
	// this is why you can't convert a []rune to []byte, or a []byte to a []rune, without going
	// through a string first.

	os.Exit(1)
	// Take a look at strings, notice that some char sets require multiple bytes per char.
	// evalString(str1)
	// evalString(str2)
	// evalString(str3)
	//
	// String tranformations, Reverse, toUpper, toLower, and trim.
	evalString(str4)

	// // See why processing UTF8 data by the byte is wrong.
	// // Watch what happens to the char data when transform by bytes.
	//    revByte() will attemp to reverse the string passed to it.
	// fmt.Println("************ B Y T E S***************")
	// evalBytehex(str1)
	// evalBytehex(revByte(str1))
	// evalBytehex(str2)
	// evalBytehex(revByte(str2))
	// evalBytehex(str3)
	// evalBytehex(revByte(str3))

	// // See why processing UTF data by rune works perfect.
	// fmt.Println("*********** R U N E S ****************")
	// evalRunehex(str1)
	// evalRunehex(revRune(str1))
	// evalRunehex(str2)
	// evalRunehex(revRune(str2))
	// evalRunehex(str3)
	// evalRunehex(revRune(str3))

	// Think about all the small function you might use everyday to process data, and
	// then consider what a disaster it would be without runes.

	// Lower case
	fmt.Println(toLower(str1))

	// Right trim
	evalString(rTrim(str4))
	fmt.Println("[" + str4 + "]")

	// Left trim
	evalString(rTrim(str4))
	fmt.Println("[" + str4 + "]")

	// Left & Right trim
	str5 := lTrim(rTrim(str4))
	fmt.Println("[" + str5 + "]")

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

	fmt.Println("Character count: ", counter)
	fmt.Println("     Byte count: ", len(sentence))
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

//************** evalBytehex **************"
// evalBytehex is like evalByte, except it show the hex values.
// ALSO: We will not show all index positions, rather focus on the order of bytes.
func evalBytehex(str string) {
	fmt.Println(str)
	sentence := []byte(str)
	counter := 0
	for _, letter := range sentence {
		counter++
		fmt.Printf("%x ", letter)
	}
	fmt.Printf(" > Chars: %v Bytes: %v\n", counter, len(sentence))

}

// ************** evalRunehex **************
// evalRunehex is like evalRune, except it show the hex values.
// ALSO: We will not show all index positions, rather focus on the order of runes.
func evalRunehex(str string) {

	fmt.Println(str)
	sentence := []rune(str)
	counter := 0
	for _, letter := range sentence {
		counter++
		fmt.Printf("%x ", letter)
	}

	fmt.Printf(" > Chars: %v Bytes: %v\n", counter, len(string(sentence)))

}

//************** R E V S T R 1 **************"
// revByte Is the wrong way to reverse a string.
// It will work correctly within the ASCII character range.
// IT WILL FAIL using UTF8 characters that require more than 1 byte per characters.
func revByte(str string) string {
	theString := []byte(str)
	for i, j := 0, len(theString); i < j; i++ {
		j--
		theString[i], theString[j] = theString[j], theString[i]
	}
	return string(theString)
}

//************** R E V S T R 2 **************")
// revRune is the correct way to reverse a string, using a slice of runes.
func revRune(str string) string {
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
	theString := str
	for i, j := 0, len(theString); i < j; i++ {
		j--
		theString[i], theString[j] = theString[j], theString[i]
	}
	return theString
}
*/

func toUpper(str string) string {
	theString := []rune(str)
	for i, j := 0, len(theString); i < j; i++ {
		j--
		if theString[i] < 123 && theString[i] > 96 {
			theString[i] = theString[i] - 32
		}
		if theString[j] < 123 && theString[j] > 96 {
			theString[j] = theString[j] - 32
		}
	}
	return string(theString)
}
func toLower(str string) string {
	theString := []rune(str)
	for i, j := 0, len(theString); i < j; i++ {
		j--
		if theString[i] < 91 && theString[i] > 64 {
			theString[i] = theString[i] + 32
		}
		if theString[j] < 91 && theString[j] > 64 {
			theString[j] = theString[j] + 32
		}
	}
	return string(theString)
}

func lTrim(str string) string {
	theStr := []rune(str)
	x := 0
	for i, j := 0, len(str); i < j; i++ {
		j--
		if theStr[i] != 32 && x == 0 {
			x = i
			break // break out of the loop, we got her did.
		}
	}
	return string(theStr[x:])
}

func rTrim(str string) string {
	theStr := []rune(str)
	x := 0
	for i, j := 0, len(str); i < j; i++ {
		j--
		if theStr[j] != 32 && x == 0 {
			x = j + 1
			break // break out of the loop, we got her did.
		}
	}
	return string(theStr[:x])
}
