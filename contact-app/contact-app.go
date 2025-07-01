// mini-projects/contact-app/contact-app.go
package contact_manager_app // Declares this file as part of the 'contact_manager_app' package

import (
	"bufio"   // For reading user input
	"fmt"     // For printing output to the terminal
	"strconv" // For converting strings to numeric types
	"strings" // For string manipulation (e.g., TrimSpace)
)

// Contact represents a contact with a Name and Phone Number.
// The capitalized first letter (Contact) makes this struct exported,
// meaning it can be accessed from other packages if needed.
type Contact struct {
	Nama    string // Name field, capitalized for exportability
	Telepon string // Phone number field, also exported
}

// contacts is a slice (dynamic array) of Contact structs.
// The lowercase first letter (contacts) makes this variable unexported,
// meaning it can only be accessed within the 'contact_manager_app' package.
var contacts []Contact

// AddContact adds a new contact to the list.
// This function is exported (capitalized AddContact) so it can be called from other packages.
func AddContact(nama, telepon string) {
	newContact := Contact{
		Nama:    nama,
		Telepon: telepon,
	}
	contacts = append(contacts, newContact) // Appends the new contact to the 'contacts' slice
	fmt.Printf("Contact of '%s' has been added.\n", nama)
}

// ListContacts displays all contacts currently in the list.
// This function is also exported (capitalized ListContacts).
func ListContacts() {
	if len(contacts) == 0 { // Checks if the contact list is empty
		fmt.Println("Empty contact list.")
		return // Exits the function if there are no contacts
	}

	fmt.Println("\nContact List:")
	fmt.Println("-------------")
	for i, c := range contacts { // Iterates through each contact in the slice
		fmt.Printf("%d. Name: %s, Phone Number: %s\n", i+1, c.Nama, c.Telepon) // Displays contact information
	}
	fmt.Println("-------------")
}

// DeleteContact removes a contact based on their name.
// This function is exported (capitalized DeleteContact).
func DeleteContact(nama string) bool {
	found := false // Flag to indicate if the contact was found and deleted
	// Creates a new slice to store contacts that are NOT being deleted
	var updatedContacts []Contact
	for _, c := range contacts { // Iterates through each contact in the current list
		if c.Nama == nama { // If the contact's name matches the one to be deleted
			found = true // Mark that the contact was found
			continue     // Skip this contact (don't add it to updatedContacts, effectively deleting it)
		}
		updatedContacts = append(updatedContacts, c) // Add contact to the new slice if it's not being deleted
	}

	if found { // If the contact was successfully found and removed
		contacts = updatedContacts // Update the main 'contacts' slice with the new (filtered) list
		fmt.Printf("Contact of '%s' has been deleted.\n", nama)
	} else { // If the contact was not found
		fmt.Printf("Contact of '%s' is not available.\n", nama)
	}
	return found // Returns true if deleted, false otherwise
}

// RunContactManagerCLI is the main function that runs the Contact Manager CLI application.
// This function is exported (capitalized 'R') so it can be called from the 'main' package.
// The 'reader' parameter is needed to read user input,
// allowing integration with the input system in main.go.
func RunContactManagerCLI(reader *bufio.Reader) {
	fmt.Println("\n== Contact Manager CLI ==")

	for { // Main loop for the contact manager, runs until the user chooses to exit
		// Displaying menu options
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View All Contacts")
		fmt.Println("3. Delete Contact")
		fmt.Println("4. Back to Main Menu")

		fmt.Print("Choose option (1-4): ")
		input, _ := reader.ReadString('\n') // Read user input
		input = strings.TrimSpace(input)    // Remove spaces or newlines
		choice, err := strconv.Atoi(input)  // Convert string to int
		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue // Continue to next iteration (display menu again)
		}

		switch choice { // Process user's choice
		case 1:
			fmt.Print("Enter Name: ")
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			fmt.Print("Enter Phone Number: ")
			telepon, _ := reader.ReadString('\n')
			telepon = strings.TrimSpace(telepon)
			AddContact(nama, telepon) // Call the exported function from this package
		case 2:
			ListContacts() // Call the exported function from this package
		case 3:
			fmt.Print("Enter Name of contact to delete: ")
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			DeleteContact(nama) // Call the exported function from this package
		case 4:
			fmt.Println("Returning to Main Menu.")
			return // Exit this function, returning to the main menu loop in main.go
		default:
			fmt.Println("Invalid option. Please choose between 1-4.")
		}
	}
}
