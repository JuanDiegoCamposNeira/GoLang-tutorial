// * Package : A way to group functions and a package it's made up by the files
// 			   in THE SAME DIRECTORY
//
// Declare a main package
package main

// Import packages ...
//
// Standard library :
//
// 	- fmt (package) : Contains functions to : - Format text
//										      - Print in console
import (
	"fmt"

	"rsc.io/quote"
)

// Call code from an external package :
//
// - You can find packages or modules from other people at : https://pkg.go.dev/
//
// For this example we'll be using rsc.io/quote package
// import required package

// Main function :  This function will execute by default when
//					the MAIN PACKAGE is runned.
func main() {
	// Print in console "Hello world"
	fmt.Println("Hello, World!")

	// Print in console the function called of rsc.io/quote
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
}

// Execute Go file :
//
//	- Open terminal
// 	- $ go run .         ----> This will execute the Main package

// Go commands ...
//
// - Open terminal
// - $ go help

// Aditional go commands (Not the only ones):
//
// Usage:
//
//         go <command> [arguments]
//
// The commands are:
//
//         bug         start a bug report
//         build       compile packages and dependencies
//         clean       remove object files and cached files
//         doc         show documentation for package or symbol
//         env         print Go environment information
//         fix         update packages to use new APIs
//         fmt         gofmt (reformat) package sources
//         generate    generate Go files by processing source
//         get         add dependencies to current module and install them
//         install     compile and install packages and dependencies
//         list        list packages or modules
//         mod         module maintenance
//         run         compile and run Go program
//         test        test packages
//         tool        run specified go tool
//         version     print Go version
//         vet         report likely mistakes in packages
