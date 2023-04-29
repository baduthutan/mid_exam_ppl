package main

import "fmt"

// The ellipsis before the string type in this example indicates that the variable parameter 'dishes' of type string is a part of the 'eat' function. This means that any number of string parameters can be passed to the 'eat' function.
func eat(dishes ...string) {
    fmt.Println("I am eating:")
	// A message indicating that we are eating is printed inside the 'eat' function before using a range loop to iterate through and print each dish in the 'dishes' slice.
    for _, dish := range dishes {
        fmt.Println("- " + dish)
    }
}

// We call the 'eat' function four times in the 'main' function, passing in various quantities of dishes each time. The first call is passed with three dishes, the second with two dishes, the third with one dish, and the fourth with none.
func main() {
    eat("pizza", "sushi", "tacos") // passing in three dishes
    eat("ramen", "curry") // passing in two dishes
    eat("salad") // passing in a dish
    eat() // passing in zero dishes
}
