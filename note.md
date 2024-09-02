# Golang

https://go.dev/doc/#learning

## Basic

### Language feature
machine language (CPU instructions represented in binary) / assembly language (CPU instructions with mnemonics) / high-level language (from C, C++ to python)

#### Translation phase

compiled vs interpreted

Compilation: once before running the code

Interpretation: translate instructions while code is executed. Garbage collection

- Garbage collection in Go: automatic memory management.

#### OOP
organize code through encapsulation.

term "**structs**" instead of the term "class"

no inheritance, no constructors, no generic

#### Concurency

1. performance limits of machines : Moore's Law: Number of transistors doubles every 18 months. But there's a limit

2. Parallelism : multiple tasks may be performed at the same time on different cores. Difficulties : when start/stop ? do tasks conflict in memory ? What if one task needs data from another task ?

Concurrent programming: management of multiple tsks at the same time.

Key requirement for large systems

Go : Goroutines, represent concurrent tasks, Channels used to communicate between tasks, Select enables task synchronization

#### Workspaces

hierarchy of directories : 3 subdirectories : src, pkg, bin (contains executables)

common organization is good for sharing

workspace's path is set in $GOPATH system variable

#### Package

first line of file names the packages

`package xinhepkg`

import :

```Go
import (
  xinhepkg
)
```

must be one package called `main`

Building the main packagegenerated an executable program

main package needs a `main()` function

`main()` is where code execution starts

#### import
`import` keyword make Go searches directories specified by GOROOT and GOPATH env variables.

#### Go Tool

`go build` compiles the program. arguments can be a list of packages or a list of .go files; create an executable for the main package, same name as the first .go file

`go fmt` formats source code files

`go doc` prints documentation for a package

`go run` compile .go files and runs the executable

`go test` runs test using files ending in "_test.go"


### variables

#### Naming
case sensitive

don't use keywords : `if`, `case`, `package`...

all var must have declarations: `var x int`, keyword, name, type

can declare many on the same line: `var x, y int`

#### Types

- integer
- floating point, fractional (decimal) values, floating point arithmetic (may use different hardware)
- strings, byte sequences

**Type declarations**: defining an alias, may improve clarity

`type Clesius float64`
`type IDnum int`

uninitialized variables have a zero value:

`var x int // x = 0`

`var x string // x = ""`

short variable declarations: with `:=` operator, to declare and assign in the same time, but can only do this inside a function

`x := 100 // `

#### Pointers (`ptr`)
`&` operator returns the address of a variable/function

`*` returns data at an address

```go
var x int = 1
var y int
var ip *int // ip is pointer to int

ip = &x // ip now points to x
y = *ip // y is now 1
```

`new` : alternate way to create a variable

`new()` function creates a variable and returns a pointer to the variable

#### Scope

the places in code where a variable can be accessed

blocks: sequence of declarations and statement within matching brackets, `{}`, including function definitions

hierarchy of **implicit blocks**

- universe block - all Go source

- package block - all source in a package

- file block - all source in a file

- `if` `for` `switch` - all code inside the statem,ent

- clause in `switch` or `select` - individual clauses each get a block

##### lexical scoping

"defined inside" is transitive

#### Deallocating memory && Garbage collection

heap: assign outside of functions, doesn't go away if you don't explicitly deallocate

stack: deallocate when the function is completed

manual deallocation, error-prone, but fast

```go
func foo() *int {
  x := 1
  return &x
}

func main() {
  var y *int
  y = foo()
  fmt.Printf("%d", *y)
}
```

Java Virtual Machine & Python Interpreter do the garbage collection, easy fo the programmer, but slow

Go is a compiled language which enables garbage collection, implementation is fast

Go's Compiler determines stack vs heap, garbage collction in the background

#### Comment
Single-line comments: `// this is a comment`

Block comments:
```go
/* comment 1
   comment 2*/
```

#### Printing
import from the fmt package. String in double quote.

`fmt.Printf()` prints a string

`fmt.Printf("Hi %s", x)` - `%s` sonvert character for a string

#### Integers
Generic int declaration : `var x int`
different lengths and sings : `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`...

**Binary operators:**

Arthmetic: `+ - * / % << >>`

comparison: `== != > < >= <=`

boolean: `&& ||`

#### Type conversion
most binary operations need operands of the same type, including assignments

```go
var x int32 = 1
var y int16 = 2
x = y // will fail because the type is different
```

convert type with `T()` operation:
```go
x = int32(y)
```

#### Floating point
`float32` - ~6 digits of precision

`float64` ~15 digits

expressed using decimals or sceintific notation
```go
var x float64 = 123.45
var y float64 = 1.2345e2
var z complex128 = complex(2,3)
```

#### Strings

ASCII (8-bit) & Unicode

Default Go : UTF-8

Code points, unicode characters, Rune, a code point in Go

strings: sequence of arbitrary bytes, Read-only, Often meant to be printed

literal: notated by double quotes, each byte is a rune (UTF-8 code point)

##### String packages
`Unicode packages`: runes are divided into many different categories, provides a set of function to test categories of runes
```go
IsDigit(r rune)
IsSpace(r rune)
IsLetter(r rune)
IsLower(r rune)
IsPunct(r rune)

ToUpper(r rune)
ToLower(r rune)
```

`strings package`
```go
Compare(a, b) // returns an integer comparing two strings, 0 if a ==b, -1 if a < b

Contains(s, substr)

HasPrefix(s, prefex)

Index(s, substr)

// String are immutable, but modified strings are returned

Replace(s, old, new, n)

ToLower(s)

ToUpper(s)

TrimSpace(s) // strip
```

`Strconv Pakcage`
conversions to and from string representations of basic data types
```go
Atoi(s) // converts string to int

FormatFloat(f, fmt, prec, bitSize) // convert floating point number to a string

ParseFloat(s, bitSize) // converta string to float
```

#### Constant
expression whose value is know at compile time

Type is inferred from righthand side (boolean, string, number), no need to declare the data type

```go
const x = 1.3
const (
  y = 4
  z = "Hi"
)

```

##### iota
starts at 1 and increment

generate a set of related but distinct constants, pratical when each constant need a unique integer

often represents a property which has several distinct possible values

Constants must be different but actual value is not important

like an enumerated type

```go
type Grades int
const (
  A Grades = iota
  B
  C
  D
  F
)
```
### Control flow
control structures

#### If statement
```go
if <condition> {
  <consequent>
}

if x > 5 {
  fmt.Printf("Youhou")
}
```

#### For loops
iterates while a condition is true; may have an initialization and update operation

```go
for <init>; <cond>;
<update> {
  <stmts>
}

for i := 0;  i < 10; i ++ {
  fmt.Printf("hi")
}

i = 0
for i < 10 {
  fmt.Printf("hi")
  i++
}

```

#### Switch/case
a multi-way if statement

```go
switch x {
  case x > 1:
  fmt.Printf("case1")
  case x < -1:
  fmt.Printf("case2")
  default:
  fmt.Printf("no case")
}
```
##### Tagless switch

#### break and continue
`break` exits the containing loop

```go
i = 0
for i < 10 {
  if i == 5 {break}
  fmt.Printf("hi")
  i++
}
```

`continue` skips the rest of the current iteration

#### scan
- scan reads user input
- takes a pointer as an argument
- typed data is written to pointer
- returns number of scanned items, and error

```go
var appleNum int

fmt.Printf("number of apples ?")
num, err := fmt.Scan(&appleNum)
fmt.Printf(appleNum)
```

### Arrays

elments accessed using subscript notation `[]`

elements initialized to zero value

```go
var x [5]int // array of 5 integers

x[0] = 2
fmt.Pinrtf(x[1]) // 0
```

#### Aray Literal
length of literal must be length of array

`var x [5]int = [5] {1, 2, 3, 4, 5}`

`x := [...]int{1, 2, 3, 4} // ... automatically give the length of array`

Use a for loop with the range keyword
```go
x := [...]int {1, 2, 3}
for i, v range x {
  fmt.Printf("ind %d, val %d", i, v)
}
```
i: index; v: value

### Slices

- a "window" on an underlying array

- variable size, up to the whole array

- `pointers` indicates the start of the slice

- `length` is the number of elements in the slice

- `capacity` is maximum number of elements => from start of slice to end of array

```go
arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
s1 := arr[1:3]
s2 := arr[2:5]

```

#### length and capacity

`len()` returns the length

`cap()` returns the capacity

```go
a1 := [3]string("a", "b", "c")
sli1 := a1[0:1]
fmt.Printf(len(sli1), cap(sli1)) // 1, 3

fmt.Printf(s1[1]) // c
fmt.Printf(s2[0]) // c
```

#### slice literals

- can be used to initialize a slice

- creates the underlying array and creates a slice to reference it

- slice points to the start of the array, length is capacity

```go
cli := []int{1, 2, 3} // dont put anything in bracket so it will be seen as a slice, and create the array
```

#### `make`

`make()` create a slice (and array)

2 argument version: specify type and length/capacity

init to zero, length = capacity: `sli = make([]int, 10)`

3 argument version: specify length and capacity separately

`sli = make([]int, 10, 15)`

#### `append`

`append()` can
- increase the size of a slice
- add elements to the end of a slice
- inserts into underlying array
- increase size of array if necessary

```go
sli = make([]int, 0, 3) // length of sli is 0
sli = append(sli, 100)
```

### Hash & Maps
- contains key/value pairs

- hash function is used to compute the slot for a key

advantage:

faster lookup than lists, constant-time vs linear-time

Arbitrary keys: not ints, like slices or arrays

disavantage:

may have collisions when 2 keys hash to same slot

**Maps**: Golang's implementation of a hash table

use `make()` to create a map

```go
var idMap map[string]int // [keyType]valueType
idMap = make(map[string]int)
```
may define a map literal
```go
idMap := map[string] int {
  "joe": 123
}
```

#### Accessing maps
referencing a value with `[key]`

returns zero if key is not present
`fmt.Println(idMap["joe"])`

add a key/value pair
`idMap["jane"] = 456`

deleting a key/value pair
`delete(idMap, "joe")`

#### Map functions

two value assignment tests for existence of the key
`id, p := idMap["joe"]`
`id` is value, `p` is prensence of key => a boolean

`len()` returns number of values

#### Iterating through a map
```go
for key, val := range idMap {
  fmt.Println(key, val)
}
```

### Struct (class)
- aggregate data type

- groups together other objects of arbitrary type

```go
type struct Person {
  name string
  addr string
  phone string
}
var p1 Person
```
- each property is a field
- p1 contains value of all fields

use dot notaion
```go
p1.name = "joe"
x = p1.addr

```

init

`p1 := new(Person)` to initializes fileds to zero

`p1 := Person(name: "joe", addr: "a st.", phone: "123")` to initialize using a struct literal
