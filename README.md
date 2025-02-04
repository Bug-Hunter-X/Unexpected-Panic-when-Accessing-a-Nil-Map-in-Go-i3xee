# Unexpected Panic when Accessing a Nil Map in Go

This repository demonstrates a common error in Go: accessing a map without checking if it's nil. This can lead to a runtime panic, which is often difficult to debug.

## The Problem

Go's maps are initialized to `nil` if not explicitly initialized.  Attempting to access a key in a `nil` map will cause a runtime panic.  The code in `bug.go` demonstrates this problem.

## The Solution

The best way to avoid this panic is to check if the map is `nil` before accessing its elements.  The `bugSolution.go` file shows a corrected version of the code that gracefully handles the case where the map is `nil`.

## How to Run

1. Clone this repository.
2. Navigate to the directory.
3. Run `go run bug.go` to see the panic.
4. Run `go run bugSolution.go` to see the corrected version.