// demonstrate the use of a buffered channel with make.
package main

import "fmt"

func main() {

	message := make(chan string, 4)  //when i put less than of execution than it gives deadlock situation 

	message <- "Rajan"
	message <- "Pintu"
	message <- "Mohan"
	message <- "Raghu"

	fmt.Println(<-message, "is goog Guy")
	fmt.Println(<-message, "Looks like osm")
	fmt.Println(<-message, "It's nyc.")
	fmt.Println(<-message, "It's nyc also")
	
}
