package main

import (
	"fmt" // Importing the fmt package for formatted I/O operations
)

func main() {
	/* 

	Rochelle owns a Beauty company and needs a program to manage massage appointments. 
	Each appointment has a client name, appointment time, and duration in minutes. 
	The program should be able to add new appointments, cancel appointments, and list 
	all appointments.

	Requirements:

	Create 3 slices to represent the appointment schedule, where each slice contains 
	client name (string), time (string), and duration (int) respectively.
	Write code to add a new appointment.
	Write code to cancel an appointment by client name.
	Write code to list all appointments.

	*/
	names := []string{}       // Slice to store client names
	times := []string{}       // Slice to store appointment times
	durations := []int{}      // Slice to store appointment durations

	choice := 0               // Variable to store user menu choice

	for true {                // Infinite loop to repeatedly show menu and process user input
		fmt.Println("1 - Add New Appointment")
		fmt.Println("2 - Cancel Appointment")
		fmt.Println("3 - List All Appointments")
		fmt.Println("4 - Exit")
		fmt.Println()
		fmt.Println("Please enter an option")
		fmt.Scan(&choice)     // Read user's menu choice

		if choice == 1 {      // If user chooses to add a new appointment
			var name, time string
			var duration int
			
			// Get appointment details from user
			fmt.Println("Enter the name of the client: ")
			fmt.Scan(&name)
			fmt.Println("Enter the time of the appointment: ")
			fmt.Scan(&time)
			fmt.Println("Enter the duration of the appointment: ")
			fmt.Scan(&duration)
			
			// Add details to the slices
			names = append(names, name)
			times = append(times, time)
			durations = append(durations, duration)

		} else if choice == 2 {  // If user chooses to cancel an appointment
			var name string
			fmt.Print("Enter the name of the client to cancel their appointment: ")
			fmt.Scan(&name)

			index := -1          // Initialize index to an invalid value
			for i, n := range names {  // Find the index of the appointment to cancel
				if n == name {
					index = i
					break
				}
			}

			if index != -1 {     // If appointment is found, remove it from the slices
				names = append(names[:index], names[index+1:]...)
				times = append(times[:index], times[index+1:]...)
				durations = append(durations[:index], durations[index+1:]...)
				fmt.Println("Appointment canceled successfully!")
			} else {             // If appointment is not found, notify the user
				fmt.Println("Appointment not found!")
			}

		} else if choice == 3 {  // If user chooses to list all appointments
			for i := range names {  // Iterate over the slices and print the details
				fmt.Printf("Client Name: %s, Appointment Time: %s, Appointment Duration %d\n", names[i], times[i], durations[i])
			}

		} else if choice == 4 {  // If user chooses to exit
			fmt.Println("Exiting...")
			return              // Exit the program

		} else {                 // If user enters an invalid choice
			fmt.Println("Invalid choice. Please enter a number between 1 and 4")
		}
	}
}
