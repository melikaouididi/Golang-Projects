// Package main is the entry point for the Go program.
package main

// Importing the fmt package for formatted I/O operations.
import "fmt"

// The main function serves as the entry point of the program.
func main() {
    // Declaring and initializing variables.
    var (
        temps  []float32 = []float32{22.5, 25.3, 19.8, 21.0, 23.4} // Slice of temperature readings.
        total  float32                                       // Variable to store the sum of temperatures.
        average float32                                      // Variable to store the average temperature.
        length float32 = float32(len(temps))                 // Variable to store the length of the temps slice.
    )

    // Loop through each temperature in the temps slice.
    for _, temp := range temps {
        total += temp // Accumulate the total sum of temperatures.
    }
    
    // Calculate the average temperature.
    average = total / length

    // Print the average temperature to the console.
    fmt.Println(average, "is your average")
}
