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

### RFC = requests for comments
example protocols: HTML, URI (Uniform resource identifier), HTTP

Golang provides pkg for important RFCs ; functions which encode and decode

`"net/http"` => `http.Get(www.google.com)`

`"net"` for TCP/IP and socket programming : `net.Dial("tcp", "uci.edu:80")`

#### JSON
"`"json"`" (RFC 7159), basic value: bool, num, str, array, "object"

**JSON Marshalling**, generating JSON representation from an object

```go
p1 := Person(...)
barr, err := json.Marshal(p1) // barr = b array
// Marshal() returns JSON representation as []byte
```
JSON Unmarshalling:
```go
var p2 Person
err := json.Unmarshal(barr, &p2)
```

#### File
basic ops: open, read, write, close, seek (move read/write head)

`"io/ioutil"`

`dat, err := ioutil.ReadFile("test.txt")`

dat is []byte

explicit open/close are not needed

large files cause a problem (at least cannot bigger than RAM)

```go
dat = "Hello, world"
err := ioutil.WriteFile("outfile.txt", dat, 0777)
// 0777 : unix-style permission bytes
```

#### OS package
`os.Open()` returns a file descriptor

`os.Close()` closes a file

`os.Read()` reads from a file into a []byte, fills the []byte, control the amount read

`os.Write()` writes the byte you want

##### example for read
```go
f. err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr) // nb => number of bytes
f.Close()
```

##### example for write
```go
f. err := os.Create("outfile.txt")
barr := []byte{1, 2, 3}
nb, err :=f.Write(barr)
nb, err := f.WriteString("Hi")
```

## Functions

```go
func main() {
  // all go program has a main
  // you never call this function
}

func foo(x int, y int) {
  fmt.Print(x * y)
}

func foo1(x int) int { // second int => return value
  return x + 1
}

func foo2(x int) (int, int) { // 2 return values
  return x, x + 1
}


```

**call by value** => passed arguments are copied to parameters, modifying aprameters has no effect outside the function

```go
func foo(y int) {
  y = y + 1
}

func main() {
  x := 2
  foo(x) // 3
  fmt.Print(x) // 2
}
```

**tradeoffs of call by value**

advantage: data encapsulation

disadvantage: copying time, large objects may take a long time to copy

**call by reference** => pass a pointer as an argument, called function has direct access to caller variable in memory.

```go
func foo(y *int) {
  *y = *y + 1
}

func main() {
  x := 2
  foo(&x)
  fmt.Print(x)
}
```

**tradeoffs of call by reference**

advantage: copying time, don't need to copy arguments

disadvantage: data encapsulation, function variables may be changed in called functions

### passing array arguments
array arguments are copied

```go
func foo(x [3]int) int {
  return x[0]
}

func main() {
  a := [3]int{1, 2, 3}
  fmt.Print(foo(a))
}
```
### passing array pointers
```go
func foo(x * [3]int) int {
  (*x)[0] = (*x)[0] + 1

  func main() {
    a := [3]int{1, 2, 3}
    foo(&a)
    fmt.Print(a)
  }
}
```

### pass slices instead
- slices contain a pointer to the array
- passing a slice copies the pointer

```go
func foo(sli[] int) int { // cannot specify the size of slice
  sli[0] = sli[0] + 1
}

func main() {
  a := []int{1, 2, 3} // without number in [] means it is a slice
  foo(a)
  fmt.Print(a)
}
```

### some tips for functions

**understandability**:
- if you are asked to find a feature, you can find it quickly / better : others can find it easily
- if you are asked about where data is used, you know where is used and where is written

**debugging principles**:
- code crashes inside a function
- two options for the cause: function is written incorrectly; data that the function uses is incorrect.

**supporting debugging**:
- functions need to be understandable, determine if actual behavior matches desired behavior
- data needs to be traceable: origin of data, global variables complicate this.

**function Naming**:
- behavior can be undrstood at a glance
- parameter naming counts too

**functional cohesion**
- function should perform only one 'operation'
- an 'operation' depends on the context
- few parameters, because debugging requires tracing function input data

**function complexity**
- function call hierarchy

**control-flow complexity**
- control-flow describes conditional paths
- partitioning conditionals

### first-class values
- Functions are first-class.
- Variables can be declared with a function type
- can be created dynamically

#### variables as functions
- declare a variable as a func
```go
var funcVar func(int) int
func incFn(x int) int {
  return x + 1
}

func main() {
  functVar = incFn // function is on right-hand side, without ()
  fmt.Print(funcVar(1))
}

#####

func applyIt(afunct func (int) int, val int) int {
  return afunct(val)
}

#####

func applyIt(afunct func (int) int, val int) int {
  return afunct(val)
}

func incFn(x int) int {return x + 1}
func decFn(x int) int {return x - 1}

func main() {
  fmt.Println(applyIt(incFn, 2))
  fmt.Println(applyIt(decFn, 2))
}

##### Anonymous functions

func applyIt(afunct func (int) int, val int) int {
  return afunct(val)
}

func main() {
  v := applyIt(func (x int) int {return x + 1}, 2) // it's the increment function
  fmt.Println(v)
}
```

#### functions as return values
```go
func MakeDistOrigin(o_x, o_y float64) func (float64, float64) float64 {
  fn := func(x, y float) float64 {
    return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
  }
  return fn
}

func main() {
  Dist1 := MakeDistOrigin(0, 0)
  Dist2 := MakeDistOrigin(2, 2)
  fmt.Println(Dist1(2, 2))
  fmt.Println(Dist2(2, 2))
}
```
- origin location is passed as an argument
- origin is built into the returned function

**Closure**
- closure = function + its environment

- when functions are passed/returned, their environment comes with them

### variadic and deferred
**variable argument number**
- functions can take a variable number of arguments
- use ellipsis ... to specify
- treated as a slice inside function
```go
func getMax(vals ...int) int { // ... means it takes arguments as much as you like, treat the arguments as a slice
  maxV := -1
  for _, v :range vals {
    if v > maxV {
      maxV = v
    }
    return maxV
  }
}
```

**variadic slice argument**
```go
func main() {
  fmt.Println(getMax(1, 3, 6, 4))
  vslice := []int{1, 3, 6, 4}
  fmt.Println(getMax(vslice...))
}
```
- can pass a slice to a variadic function
- need the ... suffix

**deferred function calls**
- call can be deferred until the surrounding function completes
- typically used for cleanup activities

```go
func main() {
  defer fmt.Println("Bye!") // wont be executed until the main completed

  fmt.Println("Hello!")
}
//  arguments are evaluated immediately but the call is deferred

#####

func main() {
  i := 1
  defer fmt.Println(i+1)
  i++
  fmt.Println("Hello!")
}
```

## Classes (= struct)
- collection of data fields and functions that share a well-defined responsibility
- data can be protected from the programmer
- data can be accessed only using methods

**associating methods with data**
- method has a receiver type that it is associated with
- use dot notation to call the method

```go
type MyInt int
func (mi MyInt) Double () int {
  return int(mi*2) // call by value
}
func main() {
  v := MyInt(3)
  fmt.Println(v.Double())
}

#####

type Point struct {
  x float64
  y float64
}

func (p Point) DistToOrig() {
  t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
  return math.Sqrt(t)
}

func main() {
  p1 := Point(3, 4)
  fmt.Println(p1.DistToOrig())
}
```

### encapsulation
**controlling access**
- Can define public functions to allow access to hidden data

```go
package data
var x int = 1
func PrintX() {fmt.Println(x)}

package main
import "data"
func main() {
  data.PrintX()
}

```
**controlling access to structs**
- hide fields of structs by starting field name with a lower-case letter
- define public methods which access hidden data
```go
package data
type Point struct{
  x float64
  y float64
}
func (p *Point) InitMe(xn, yn float64) {
  p.x = xn
  p.y = yn
} // need InitMe() to assign hidden data fields

#####

func (p *Point) Scale(v float64) {
  p.x = p.x * v
  p.y = p.y * v
}

func (p *Point) PrintMe() {
  fmt.Println(p.x, p.y)
}

package main

func main() {
  var p data.Point
  p.InitMe(3, 4)
  p.Scale(2)
  p.PrintMe()
}
```

### Point receivers
**limitations of methods**
- receiver is passed implicitly as an argument to the method
- method cannot modify the data inside the receiver
```go
// example: OffsetX() should increase x coordinate
func main() {
  p1 := Point(3, 4)
  p1.OffsetX(5)
}
```
- if receiver is large, lots of copying is required
```go
type Image [100][100]int
func main() {
  il := GrabImage()
  il.BlurImage()
}
// 10,000 ints copied to BlurImage()

#####
// using pointer receivers
func (p *Point) OffsetX(v float64) {
  p.x = p.x + v
}
```
- receiver can be a pointer to a type

#### referencing and dereferencing
- no need to dereference
```go
func (p *Point) OffsetX(v int) {
  p.x = p.x + v
}
```
- Point is referenced as p, not *p
- dereferencing is automatic with `.` operator
- no need to reference
```go
func main() {
  p := Point{3, 4}
  p.OffsetX(5)
  fmt.Println(p.x)
}
```

#### using pointer receivers
- good programming pratice: all methods for a type have pointer receivers, or have non-pointer receivers
- mixing pointer/non-pointer reference for a type will get confusing

## Polymorphism
- ability for an object to have different "forms" depending on the context
- different implementations for each class, same signature (name, params, return)
- example: `Area()` function
- - rectangle, area = base * height
- - triangle, area = 0.5 * base * height
- identical at a high level of abstraction and different at a low level of abstraction

- traditional solution : inheritance => golang doesn't have

### interfaces
Go's way to handle polymorphism

- set of method signatures: name, parameters, return values, implementation is NOT defined
- used to express conceptual similarity between types
- example: `Shape2D interface`
- - all 2D shapes must have `Area()` and `Perimeter()`

#### satisfying an interface
- type satisfies an interface if type defines all methods specified in the interface
- - same method signatures
- - additional methods are OK
- similar to inheritance with overriding

**defining an interface type**
```go
type Shape2D interface {
  Area() float64
  Perimeter() float64
}
type Triangle {...}
func (t Triangle) Area() float64 {...}
func (t Triangle) Perimeter() float64 {...}
```
- no need to state it explicitly

#### concerte vs interface types
**concrete types**:
- specify the exact representation of the data and methods
- complete method implementation is included

**interface types**:
- specifies some method signatures
- implementations are abstracted

**interface values**:
- can be treated like other values, assigned to variables, passed, returned
- interface values have two components
- 1. dynamic type: concrete type which it is assigned to
- 2. dynamic value: value of the dynamic type

**defining an interface type**
```go
type Speaker interface { Speak() }
type Dog struct {name string}
func (d Dog) Speak() {
  fmt.Println(d.name)
}
func main() {
  var s1 Speaker
  var d1 Dog{"Brian"}
  s1 = d1 // dog type satisfies speaker type
  s1.Speak()
}
```

**interface with nil dynamic value**
- an interface can have a nil dynamic value
```go
var s1 Speaker
var d1 *Dog
s1 = d1
```
- d1 has no concrete value yet
- s1 has a dynamic type but no dynamic value
- can still call the `Speak()` method of `s1`
- need to check inside the method
```go
func (d *Dog) Speak() {
  if d == nil {
    fmt.Println("<noise>")
  } else {
    fmt.Println(d.name)
  }
var s1 Speaker
var d1 *Dog
s1 = d1
s1.Speak()
}
```
**nil interface value**
- interface with nil dynamic type
- very different from an interface with a nil dynamic value

**nil dynamic value** and **valid dynamic type**:

`var s1 Speaker ; var d1 *Dog ; s1 = d1`

=> can call a method since type is known

**nil dynamic type**

`var s1 Speaker`

=> cannot call a method, runtime error

#### Using interface
**ways to use an interface**
- need a function which takes multiple types of parameter
- Function `foo()` parameter
- - Type X or type Y
- define interface Z
- `foo()` parameter is interface Z
- type X and Y satisfy Z
- interface methods must be those needed by `foo()`

Example: Pool in a yard
- put a pool in my yard
- pool needs to fit in my yard, total area must be limited
- pool needs to be fenced -> total permimeters must be limited
- need to determine if a pool shape satisfies criteria
- `FitInYard()` -> takes a shape as argument and returns `true` if the shape satisfies criteria
- many possible shape types -> Rectangle, triangle, circle, etc
- Valid shape types must have: `Area()` and `Perimeter()`
```go
// interface for shapes
type Shape2D interface {
  Area() float64
  Perimeter() float64
}

type Triangle {...}
func (t Triangle) Area() float64 {...}
func (t Triangle) Perimeter() float64 {...}

type Rectangle {...}
func (t Rectangle) Area() float64 {...}
func (t Rectangle) Perimeter() float64 {...}

func FitInYard(s Shape2D) bool {
  if (s.Area() > 100 && s.Perimeter() > 100 ) {
    return true
  }
  return false
}
```
**empty interface**
- empty interface specifies no methods
- all types satisfy the empty interface
- use it to have a function accept any type as a parameter

```go
func PrintMe(val interface{}) {
  fmt.Println(val)
}
```

#### Type assertions
**concealing type differences**
- interfaces hide the differences between types
- sometimes you need to treat different types in different ways

**exposing type differences**
- example: graphics program
- `DrawShape()` will draw any shape
- - `func DrawShape (s Shape2D) {...}`
- underlying API has different drawing functions for each shape
- - `func DrawRect (r Rectangle) {...}`
- - `func DrawTriangle (t Triangle) {...}`
- concrete type of shape s must be determined

**type assertions for disambiguation**
- type assrtions can be used to determine and extract the underling concrete type
```go
func DrawShape (s Shape2D) bool {
  rect, ok := s.(Rectangle)
  if ok {
    DrawRect(rect)
  }
  tri, ok := s.(Triangle)
  if ok {
    DrawTriangle(tri)
  }
}
```
- type assertion extracts Rectangle from `Shape2D` -> concrete type in parentheses
- if interface contains concrete type => `rect == concrete type, ok == true`
- if interface doesnot contain concrete type => `rect == zero, ok == false`

**type switch**
- switch statement used with a type assertion
```go
func DrawShape(s Shape2D) bool {
  switch := sh := s.(type) {
    case Rectangle:
    DrawRect(sh)
    case Triangle:
    DrawTriangle(sh)
  }
}
```

#### Error handling
**Error interface**
- many go programs return error interface objects to indicate errors
```go
type error interface {
  Error() string
}
```
- correct operation: error == nil
- incorrect operation: Error() prints error message

**handling errors**
```go
f, err := os.Open("/harris/test.txt")
if err != nil {
  fmt.Println(err)
  return
}
```

## Concurrency

### Parallel execution
- two programs execute in parallel if they execute at exactly the same time
- At time t, an instruction is being performed for both P1 and P2
- need replicated hardware
- tasks *may* complete more quickly

**speedup without parallelism
- design faster processors, get speed up without changing software
- design processor with more memory

**Dynamic Power**: `P = alpha * CFV^2`
- `alpha` is percent of time switching
- `C` is capacitance (related to size)
- `F` is the clock frequency
- `V` is voltage swing (from low to high). Important, 0 to 5V uses much power than 0 to 1.3V

**Dennard Scaling**:
- voltage should scale with transistor size
- keep power consumption, and temperature low
- problem: voltage can't go too low, must stay above threshold voltage; noise problems occur
- problem: doesn't consider leakage power

**multi-core systems**
- cannot increase frequency
- can still add processor cores, without increasing frequency: trend is apparent today
- parallel execution is needed to exploit multi-core systems
- code made to execute on multiple cores
- different programs on different cores

### Concurrent execution
- concurrent execution is not necessarily the same as parallel execution
- *Concurrent*: start and end times overlap
- *Parallel*: execute at exactly the same time

**concurrent vs parallel**
- parallel tasks must be executed on different hardware
- concurrent tasks may be executed on the same hardware, only one task actually executed at a time
- mapping from tasks to hardware is not directly controlled by the programmer => at least not in go

**concurrent programming**
- programmer determines which tasks **can** be executed in parallel
- mapping tasks to hardware: operating system, go runtime scheduler

**hiding latency**
- concurrency improves performance, even without parallelism
- tasks must periodically wait for something: waiting for memory
- other concurrent tasks can operate while one task is waiting

**hardware mapping in Go**
- programmer does not determine the hardware mapping
- programmer makes parallelism possible
- hardware mapping depends on many factors: where data; what are the communication costs

### Processes
- an instance of a running program
- thins unique to a process
- 1. memory: virtual address space, code, stack, heap, shared libraries
- 2. registers: program counter, data regs, stack ptr, ...

**operating systems**
- allows many processes to execute concurrently
- processes are switched quickly: 20ms
- usedr has the impression of parallelism
- operating system must give processes fair access to resources

### Scheduling
**scheduling processes**
- operating system schedules processes for execution
- gives the illusion of parallel execution
- os gives fair access to CPU, memory, etc.

**context switch**
- control flow changes from one process to another
- process 'context' must be swapped

### threads vs. processes
- threads share some context
- many threads can exist in one process
- os schedules threads rather than process

### Goroutine (I)
- like a thread in go
- many Goroutines execute within a single os thread
- schedules goroutines inside an os thread
- like a little os inside a single os thread
- **logical processor** is mapped to a thread

### Interleavings
- order of execution within a task is known
- order of execution between concurrent tasks is unknown (non-deterministic)
- many interleavings are possible
- must consider all possibilities

### Race condition
- outcome depends on non-deterministic ordering
- races occur due to **communication**

**communication between tasks**
- threads are largely independent but not completely independent
- example: web server, 1 thread per client
- example: image processing, 1 thread per pixel block

### Goroutine (II)
- one goroutine is created automatically to execute the `main()`
- other goroutines are created using the `go keyword`

**exiting a goroutine**
- a goroutine exists when its code is complete
- when the main goroutine is complete, all other goroutines exit
```go
// early exit
func main() {
  go fmt.Printf("New routine")
  fmt.Printf("Main routine")
}
```
- only "Main routine" is pinted
- main finished before the new goroutine started

**delayed exit**
```go
func main() {
  go fmt.Peintf("New routine")
  time.Sleep(100 * time.Millisecond)
  fmt.Printf("Main routine")
}
```
- add a delay in the main routine to give the new routine a chance to complete
- "New routineMain Routine" is now printed

**timing with goroutines**
- adding a delay to wait for a goroutine is bad!
- time assumption may be wrong
- timing is non-deterministic

#### basic synchronization
- using **global events**, whose execution is viewed by all threads, simultaneously

**synchronization example**
```
// task 1
x = 1
x = x + 1
GLOBAL EVENT

// task 2
if GLOBAL EVENT
  print x
```
- `GLOBAL EVENT` is viewed by all tasks at the same time
- print must occur after update of x
- synchronization is used to restrict bad interleavings

#### wait groups
**sync WaitGroup**:
- sync package contains functions to synchronize between goroutines
- `sync.WaitGroup` forces a goroutine to wait for other goroutines
- example: contains an internal counter
- - increment counter for each goroutine to wait for
- - decrement counter when each goroutine completes
- 1. `wg.Add(1)` => increments the counter
- 2. `wg.Done()` => decrements the counter
- 3. `wg.Wait()` => blocks until `counter == 0`

```go
func foo(wg *sync.WaitGroup) {
  fmt.Printf("New routine")
  wg.Done()
}

func main() {
  var wg sync.WaitGroup
  wg.Add(1)
  go foo(&wg)
  wg.Wait()
  fmt.Printf("Main routine")
}
```

#### Goroutine communication
- goroutines usually work together to perform a bigger task
- often need to send data to collaborate
- example: find the product of 4 integers
- - make 2 goroutines, each multiplies a pair
- - main goroutine multiplies the 2 result
- need to send ints from main routine to the 2 sub-routines
- need to send results from sub-routines back to main routine

**channels**
- transfer data between goroutines
- channels are typed
- use `make()` to crate a channel: `c := make(chan int)`
- send and receive data using the `<-` operator
- send data on a channel: `c <- 3`
- receive data from a channel: `x := <- c`

```go
func prod(v1 int, v2 int, c chan int) {
  c <- v1 * v2
}

func main() {
  c :=make(chan int)
  go prod(1, 2, c)
  go prod(3, 4, c)
  a := <- c
  b := <- c
  fmt.Println(a * b)
}
```
#### blocking on channels
**unbuffered channel**
- unbuffered channels cannot hold data in transit: default is unbuffered
- sending blocks until data is received
- receiving blocks until data is sent

**blocking and synchronization**
- channel communication is synchronous
- blocking is the same as waiting for communication
- receiving and ignoring the result is same as a `Wait()`

#### buffered channel
**channel capacity**
- channels can contain a limited number of objects: default size 0 (unbuffered)
- **capacity** is the number of objects it can hold in transit
- optional argument to make() defines channel capacity: `c := make(chan int, 3)`
- sending only blocks if **buffer is full**
- receiving only blocks if **buffer is empty**

**use of buffering**
- sender and receiver do not need to operate at exactly the same speed
- speed mismatch is acceptable
- average speeds must still match

**iterating through a channel**
- common to iteratively read from a channel
```go
for i := range c {
  fmt.Println(i)
}
```
- continues to read from channel c
- one iteration each time a new value is received
- `i` is assigned to the read value
- iterates when sender calls `close(c)`

**receiving from multiple goroutines**
- multiple channels may be used to receive from multiple sources
- data from both sources may be needed
- read sequentially
```go
a := <- c1
b := <- c2
fmt.Println(a * b)
```

**select statement**
- may have a choice of which data to use: i.e. first-come first-served
- use `select` statement to wait on the first data from a set of channels
```go
select {
  case a = <- c1:
    fmt.Println(a)
  case b = <- c2:
    fmt.Println(b)
}
```

#### select
**select send or receive**
- may select either send or receive operations
```go
select {
  case a = <- inchan:
    fmt.Println("Received a")
  case outchan <- b:
    fmt.Println("sent b")
}
```

**select with an abort channel**
- use select with a **separated abort channel**
- may want to receive data until an **abort signal** is received
```go
for {
  select {
    case a <- c:
      fmt.Println(a)
    case <- abort:
      return
  }
}
```

**default select**
- may want a default operation to avoid blocking
```go
select {
  case a = <- c1:
    fmt.Println(a)
  case b = <- c2:
    fmt.Println(b)
  default:
    fmt.Println("nop")
}
```

#### multual exclusion
**Goroutines sharing variables**
- sharing variables concurrently can case problems
- two goroutines writing to a shared variable can interfere with each other
- (concurrency-safe) function can be invoked cocurrently without interfering with other goroutines

**example:** variable sharing example
```go
var i int = 0
var wg sync.WaitGroup

func inc() {
  i = i + 1
  wg.Done()
}

func main() {
  wg.Add(2)
  go inc()
  go inc()
  wg.Wait()
  fmt.Println(i)
}
// 2 goroutines write to i
// i should equal 2 but not always happen
// explication is following:
```

**granularity of concurrency**:
- concurrency is at the machine code level
- `i = i + 1` might be 3 machine instructions: read, increment, write
- both tasks may read 0 for `i` value

**correct sharing**
- don't let 2 goroutines write to a shared variable at the same time
- need to restrict possible interleavings
- access to shared variables cannot be interleaved

**mutual exclusion (Mutex)**
- code segments in different goroutines which cannot execute concurrently
- writing to shared variables should be mutually exclusive

**`Sync.Mutex`**
- a Mutex ensures mutual exclusion
- uses a `binary semaphore`
- flag up: shared variable is in use; flag down - shared variable is available

**`Sync.Mutex` methods**
- `Lock()` method puts the flag up: shared variable in use
- if lock is already taken by a goroutine, `Lock()` blocks until the flag is put down
- `Unlock()` method puts the flag down: Done using shared variable
- When `Unlock()` is called, a blocked `Lock()` can proceed
```go
// increment operation is now mutually exclusive
var i int = 0
var mut sync.Mutex
func inc() {
  mut.Lock()
  i = i + 1
  mut.Unlock()
}
```
**synchronous initialization**
- must happen once
- must happen before everything else
- how do you perform initialization with multiple goroutines?
- could perfrom initialization before starting the goroutines

**`Sync.Once`**
- has 1 method, `once.Do(f)` => function `f` will only be executed 1 times, even if it is called in multiple goroutines
- all calls to `once.Do()` block until the first returns, ensures that initialization executes first

**Example:**
```go
// make 2 goroutiens, initialization only once
// each goroutine executes dostuff()

var wg sync.WaitGroup

func main() {
  wg.Add(2)
  go dostuff()
  go dostuff()
  wg.Wait()
}

// setup() should execute only once
// "hello" should not print until setup() returns

var  on sync.Once
func setup() {
  fmt.Println("Init")
}

func dostuff() {
  on.Do(setup)
  fmt.Println("hello")
  wg.Done()
}

// init appears only once
// init appears before hello is printed
```

**synchroniazation dependencies**
- synchronization causes the execution of different goroutines to depend on each other

**Deadlock**
- circular dependencies cause all involved goroutines to block
- can be caused by waiting on channels

**Example:**
```go
func dostuff(c1 chan int, c2 chan int) {
  <- c1
  c2 <- 1
  wg.Done()
}
// read from first channel and write on second channel

func main() {
  ch1 := make(chan int)
  ch2 := make(chan int)
  wg.Add(2)
  go dostuff(ch1, ch2)
  go dostuff(ch2, ch1)
  wg.Wait()
}
// dostuff() argument order is swapped
// each goroutine blocked on channel read
```

**Deadlock detection**
- Golang runtime automatically detects when all goroutines are deadlocked `fatal error: all goroutiens are asleep - deadlock!`
- cannot detect when a subset of goroutines are deadlocked

**Dining philosophers problem**
- classic problem involving concurrency and synchronization

*Problem*
- 5 philosophers sitting at a round table
- 1 chopstick is placed between each adjacent pair
- want to eat rice from their plate, but needs 2 chopsticks
- only 1 philosopher can hold a chopstick at a time
- not enough chopsticks for every one to eat at once
- each chopstick is a mutex
- each philosopher is associated with a goroutine and 2 chopsticks

```go
type ChopS struct { sync.Mutex }
type Philo struct {
  leftCS, rightCS *ChopS
}

func (p Philo) eat() {
  for {
    p.leftCS.Lock()
    p.rightCS.Lock()

    fmt.Println("eating")
    p.leftCS.Unlock()
    p.rightCS.Unlock()
  }

func main() {
  CSticks := make([]*ChopS, 5)
  for i := 0; i < 5; i++ {
    CSticks[i] = new(ChopS)
  }
  philos := make([]*Philo, 5)
  for i := 0; i < 5; i++ {
    philos[i] = &Philo{CSticks[i], CSticks[(i+5) % 5]}
  }

  for i := 0; i < 5; i++ {
    go philos[i].eat()
  }
  // all philosophers might lock their left chopsticks concurrently
  // no one can lock their right chopsticks
}
```

**Deadlock solution**
```go
// each philosopher picks up lowest numbered chopstick first
philos[i] = &Philo{CSticks[i], CSticks[(i + 1) % 5]}
//philosopher 4 picks up chopstick 0 before chopstick 4
// philosopher blocks allowing philosopher 3 to eat
// no deadlock, but philosopher 4 may starve
```
