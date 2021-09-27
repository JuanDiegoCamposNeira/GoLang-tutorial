package main

// NOTE : Because the module isn't published yet we need to tell Go to find
//		  the file in our local file system, to do so, run the next command :
//					$ go mod edit -replace=<module_name>=<path_to_local_module>
//
//					I.e:
//						- Module import
//							import "example.com/greetings"
//
//						- Run the command:
//							$ go mod edit -replace=example.com/greetings=../greetings
//
//						- After running the command, go will
// 						  interpretate "example.com/greetings" as "../greetings"
//						  This would be reflected in go.mod as "replace example.com/greetings => ../greetings"

// Import packages
import (
	"fmt"
	"log"

	"example.com/greetings"
)

//-------------------------------------------
//			hello.go -> 1.0
//-------------------------------------------
// func main() {
// 	// Use Hello function from greetings module
// 	message := greetings.Hello("Juan Diego")
// 	fmt.Println(message)

// 	// Use the Test function from greetings module
// 	greetings.Test("J")
// }

//-------------------------------------------
//	   hello.go -> 2.0 ----> Error handling
//-------------------------------------------
// func main() {

// 	// Set properties of the predefined Logger, including
// 	// the log entry prefix and a flag to disable printing
// 	// the time, source file, and line number.
// 	log.SetPrefix("greetings: ")
// 	log.SetFlags(0)

// 	//-------------------------------------------------------
// 	//					Hello function
// 	//-------------------------------------------------------
// 	//----------------- Valid message request ---------------
// 	// Request a greeting message WITH A VALID NAME
// 	message, err := greetings.Hello("Juan Diego")
// 	if err != nil {
// 		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
// 		log.Fatal(err)
// 	}
// 	// If no error exist, print the message returned
// 	fmt.Println(message)

// ----------------- Invalid message request ---------------
// // Request a greeting message WITH AN INVALID NAME.
// message1, err1 := greetings.Hello("")
// if err1 != nil {
// 	// Fatal is equivalent to Print() followed by a call to os.Exit(1).
// 	log.Fatal(err1)
// }
// // If no error exist, print the message returned
// fmt.Println(message1)

// 	//-------------------------------------------------------
// 	//					Test function
// 	//-------------------------------------------------------
// 	// Use the Test function from greetings module
// 	greetings.Test("J")
// }

//---------------------------------------------------
//	   hello.go -> 3.0 ----> Return multiple values
//---------------------------------------------------
func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//-------------------------------------------------------
	//					Hello function
	//-------------------------------------------------------
	//----------------- Valid message request ---------------
	// Request a greeting message WITH A VALID NAME
	message, err := greetings.Hello("Juan Diego")
	if err != nil {
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(err)
	}
	// If no error exist, print the message returned
	fmt.Println(message)

	//----------------- Invalid message request ---------------
	// // Request a greeting message WITH AN INVALID NAME.
	// message1, err1 := greetings.Hello("")
	// if err1 != nil {
	// 	// Fatal is equivalent to Print() followed by a call to os.Exit(1).
	// 	log.Fatal(err1)
	// }
	// // If no error exist, print the message returned
	// fmt.Println(message1)

	//-------------------------------------------------------
	//					Hellos function
	//-------------------------------------------------------
	// Declare and initialize slice to be passed as a parameter to the function
	names := []string{"Juan", "Diego", "Campos"}
	// Call Hellos function in greetings module
	hellos, err2 := greetings.Hellos(names)
	// If error exists, print error and exit program
	if err2 != nil {
		log.Fatal(err2)
	}
	// Otherwise, print the returned map
	fmt.Println(hellos)

	//-------------------------------------------------------
	//					Test function
	//-------------------------------------------------------
	// Use the Test function from greetings module
	greetings.Test("J")
}
