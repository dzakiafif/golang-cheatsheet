// variables 
var msg string
msg = "Hello"

// shortcut 
msg := "Hello"

// constants
const Phi = 1.618

// string 
str := "Hello"

// string multiline
str := `Multiline
string`

// number
num := 3          // int
num := 3.         // float64
num := 3 + 4i     // complex128
num := byte('a')  // byte (alias for uint8)

var u uint = 7        // uint (unsigned)
var p float32 = 22.7  // 32-bit float

// arrays
// var numbers [5]int
numbers := [...]int{0, 0, 0, 0, 0}

// slices
slice := []int{2, 3, 4}
slice := []byte("Hello")