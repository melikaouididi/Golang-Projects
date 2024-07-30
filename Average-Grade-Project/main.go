package main

import (
	"fmt"
	"math"
)

/*
  Ask the user for 5 numbers and calculate the average of those 5 numbers
*/
func main() {
	// Declare an array to store 5 integers
	var numbers [5]int

	// Loop to get 5 numbers from the user
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Insert a number:")
		fmt.Scan(&numbers[i]) // Read user input and store it in the array
	}

	// Variable to store the sum of the numbers
	var sum float32 = 0
	
	// Loop to calculate the sum of the numbers
	for i := 0; i < len(numbers); i++ {
		sum = sum + float32(numbers[i]) // Add each number to the sum
	}

	// Calculate the mean (average) of the numbers
	mean := sum / float32(len(numbers))
	fmt.Println("The average is", mean) // Print the average

	// Declare an array to store the numbers as float32
	var floatNumbers [5]float32
	for i := 0; i < len(numbers); i++ {
		floatNumbers[i] = float32(numbers[i]) // Convert each number to float32 and store it in the new array
	}

	// Calculate the standard deviation of the numbers
	sd := standardDeviation(floatNumbers, mean)
	fmt.Println("The Standard Deviation of the above array is:", sd) // Print the standard deviation

	// Loop to print each number with its position
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Number", i+1, "->", numbers[i])
	}
}

/*
  MYSTERY CHALLENGE

  Calculate the standard deviation for this array
*/
func standardDeviation(numbers [5]float32, mean float32) float32 {
	var sd float32
	// Loop to calculate the sum of the squared differences from the mean
	for i := 1; i < len(numbers); i++ {
		sd += float32(math.Pow(float64(numbers[i]-mean), 2)) // Calculate the squared difference and add it to sd
	}

	// Calculate the standard deviation
	sd = float32(math.Sqrt(float64(sd) / float64(len(numbers)))) // Take the square root of the average squared difference
	return sd
}
