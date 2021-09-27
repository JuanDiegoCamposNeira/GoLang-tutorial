package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//-------------------------------------------------------------------------------------------
//									Go functions
//-------------------------------------------------------------------------------------------
//
//	Declare a function :
//
//			func <Function_Name> ( [ <parameter_name> <parameter_type>, ... ] ) [ <return_type>, ... ] {
//				Code goes here
//			}
//
//			Return value(s) : A go function could return multiple values
//
//
//			IMPORTANT NOTE:  If the name of the function begins with a capital letter can be
//							 called by a function not in the same package.
//
//------------------------------  End of Functions ------------------------------------------

//-------------------------------------------------------------------------------------------
//									Go variables
//-------------------------------------------------------------------------------------------
//
//	- Declare a variable :      			var <variable_name> <variable_type>
// 	- Assign a value to a variable : 		<vairable_name> = <value>
//
//	- Shortcut to declare a variable and assign a value to it:
//
//					<variable_name> := <variable_value>
//
//------------------------------  End of Variables ------------------------------------------

//-------------------------------------------------------------------------------------------
//									Go data structures
//-------------------------------------------------------------------------------------------
//--- Slice (vector) :
// 						- Declare a slice :
//							var <slice_name> []<slice_type>
//
//						- Initialize the slice :
//							<slice_name> = make(<slice_type>, <slice_length>[, <slice_capacity>])
//
//										-> slice_length : length of the slice
// 										-> slice_capacity : Max capacity of the slice
//
//						- Shortcut to declare and initialize the slice
//							a)  <slice_name> := make(<slice_type>, <slice_length>[, <slice_capacity>])
//							b) 	<slice_name> := []<slite_type>{ [<slice_0>, <slice_1> ... ] }
//
//
//--- Map :
//			- A Go map type will look like = map[KeyType]ValueType
//
//			- Declare a map :
//				var <map_name> map[<key_type>]<value_type>
//								   ^^^^^^^^^^ - NOT OPTIONAL
//
//			- Initialize a map :
//				<map_name> = make(map[<key_type>]<value_type>)
// 									  ^^^^^^^^^^ - NOT OPTIONAL
//
//			- Shortcut to declare an initialize the map
// 				a)  <map_name> := make(map[<key_type>]<value_type>)
// 									  ^^^^^^^^^^ - NOT OPTIONAL
//
//				b)  <map_name> := map[<key_type>]<value_type>{
//									<key_1>: <value_1>,
//									...
//									<key_n>: <value_n>,
//								  }
//
//			- Working with maps :
//				-> Access an element :
//					a) <variable_name>  := <map_name>[<key>] -> returns the value of the given key
//					b) <var_name>, <var_name> := <map_name>[<key>] -> returns the value of the key
//																	  and a boolean containign true if
//																	  the key exists, otherwise false
//																	  if the key doesn't exists
//				-> Delete an element :
//					delete(<map_name>, <key_to_be_deleted>)
//
//---------------------------------  End of Structures --------------------------------------

//------------------------------------
//			greetings.go -> 1.0
//------------------------------------
// Hello returns a greeting for the named person.
// func Hello(name string) string {
// 	// Return a greeting that embeds the name in a message
// 	message := fmt.Sprintf("Hi, %v. \n Welcome :)", name)

// 	// fmt.Sprintf( "<string>", <string> ) -> This function will format the string of the first argument
// 	// 										  and will replace the values in it with the rest of the arguments

// 	message += fmt.Sprintf("\n this is added text ... %v, %v", name, name)
// 	return message
// }

func Test(name string) {
	message := fmt.Sprintf("Hello %v, %v", name, name)
	fmt.Println(message)
}

//-------------------------------------------
//	greetings.go -> 2.0 ----> Error handling
//-------------------------------------------
// Hello returns a greeting for the named person.
// func Hello(name string) (string, error) {

// 	// Check if no name was given ...
// 	if name == "" {
// 		return "", errors.New("empty name given ... :(")
// 	}

// 	// Return a greeting that embeds the name in a message
// 	message := fmt.Sprintf("Hi, %v. \nWelcome :)", name)

// 	// fmt.Sprintf( "<string>", <string> ) -> This function will format the string of the first argument
// 	// 										  and will replace the values in it with the rest of the arguments

// 	message += fmt.Sprintf("\nThis is added text ... %v, %v", name, name)
// 	return message, nil
// }

//-----------------------------------------------------------------------------------------
//					greetings.go -> 3.0 ----> Return random greeting
//-----------------------------------------------------------------------------------------
//----------------------------------------------------------------
//	Init fucntion : This function will be execute
//	   				once the package is called.
//
//	- This effectively allows us to set up database connections,
//	  or register with various service registries, or perform
//	  any number of other tasks that you typically only want to
//	  do once.
//---------------------------------------------------------------
// init sets initial values for variables used in the function. (In this case)
func init() {
	rand.Seed(time.Now().UnixNano())
}

//------------------  End of Init function  ---------------------

// // Function to send a random greeting message
// func Hello(name string) (string, error) {

// 	// Check if no name was given ...
// 	if name == "" {
// 		return "", errors.New("empty name given ... :(")
// 	}

// 	// Return a greeting that embeds the name in a message
// 	// message := fmt.Sprintf("Hi, %v. \nWelcome :)", name)
// 	message := fmt.Sprintf(randomFormat(), name)

// 	// fmt.Sprintf( "<string>", <string> ) -> This function will format the string of the first argument
// 	// 										  and will replace the values in it with the rest of the arguments

// 	message += fmt.Sprintf("\nThis is added text ... %v, %v", name, name)
// 	return message, nil
// }

// // randomFormat returns one of a set of greeting messages. The returned
// // message is selected at random.
// func randomFormat() string {
// 	// A slice of message formats.
// 	formats := []string{
// 		"Hi, %v. Welcome!",
// 		"Great to see you, %v!",
// 		"Hail, %v! Well met!",
// 	}

// 	// Return a randomly selected message format by specifying
// 	// a random index for the slice of formats.
// 	return formats[rand.Intn(len(formats))]
// 	// Intn returns, as an int, a non-negative
// 	// pseudo-random number in [0,n) from the default Source.
// 	// It panics if n <= 0
// }

//-----------------------------------------------------------------------------------------
//			greetings.go -> 4.0 ----> Return greetings for multiple people
//-----------------------------------------------------------------------------------------
// Function to send a random greeting message
func Hello(name string) (string, error) {

	// Check if no name was given ...
	if name == "" {
		return "", errors.New("empty name given ... :(")
	}

	// Return a greeting that embeds the name in a message
	// message := fmt.Sprintf("Hi, %v. \nWelcome :)", name)
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat()) // This is intentional to make a test fail (TestHelloName)

	// fmt.Sprintf( "<string>", <string> ) -> This function will format the string of the first argument
	// 										  and will replace the values in it with the rest of the arguments

	// message += fmt.Sprintf("\nThis is added text ... %v, %v", name, name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
	// Intn returns, as an int, a non-negative
	// pseudo-random number in [0,n) from the default Source.
	// It panics if n <= 0
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}
