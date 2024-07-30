package main

import "fmt" // Importing the fmt package for formatted I/O operations

func main() {
	var title string          // Variable to store the book title
	var author string         // Variable to store the book author
	var publicationYear int   // Variable to store the book publication year

	var exit bool = false     // Variable to control the exit condition of the loop

	for exit == false {       // Loop runs until exit is set to true
		fmt.Println("Please enter the book title:")
		fmt.Scan(&title)      // Read the book title from the user

		if title == "exit" || title == "Exit" { // Check if the user wants to exit
			exit = true       // Set exit to true to break the loop
			continue          // Skip the rest of the loop and start the next iteration
		}
	
		fmt.Println("Please enter the book author:")
		fmt.Scan(&author)     // Read the book author from the user
	
		fmt.Println("Please enter the book publication year:")
		fmt.Scan(&publicationYear)  // Read the book publication year from the user
	
		// Check if the publication year is not one of the specified years
		if publicationYear != 1635 && publicationYear != 1724 && publicationYear != 1820 && publicationYear != 1920 && publicationYear != 2020 {
			exit = true       // Set exit to true to break the loop
			continue          // Skip the rest of the loop and start the next iteration
		}
			
		// Determine the century of the book based on the publication year
		switch publicationYear {
		case 1635:
			fmt.Println("your book belongs to the 17th century")
		case 1724:
			fmt.Println("your book belongs to the 18th century")
		case 1820:
			fmt.Println("your book belongs to the 19th century")
		case 1920:
			fmt.Println("your book belongs to the 20th century")
		case 2020:
			fmt.Println("your book belongs to the 21st century")
		}
	}
}
