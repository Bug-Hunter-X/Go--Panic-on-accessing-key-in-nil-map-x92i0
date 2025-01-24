# Go: Panic on Accessing Key in Nil Map

This repository demonstrates a common error in Go: panicking when accessing a key in a `nil` map.  Go's maps don't have a default behavior like returning `null` or zero-value; instead, accessing a key in a `nil` map causes a runtime panic.

## The Bug

The provided `bug.go` file contains a simple function that attempts to access a key in an uninitialized map.  This will cause the program to crash.

## The Solution

The `bugSolution.go` file provides a corrected version which uses a check to avoid accessing a `nil` map and instead handles the case gracefully.

## How to Reproduce

1. Clone this repository.
2. Run `go run bug.go` to see the panic.
3. Run `go run bugSolution.go` to see the corrected behavior.