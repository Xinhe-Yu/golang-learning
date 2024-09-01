# Golang

https://go.dev/doc/#learning

#### Language feature
machine language (CPU instructions represented in binary) / assembly language (CPU instructions with mnemonics) / high-level language (from C, C++ to python)

##### Translation phase

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
