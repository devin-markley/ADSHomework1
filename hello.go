package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
)

type Contact struct {
	firstName string
	lastName  string
	street    string
	city      string
	state     string
	zip       string
	phone     string
	email     string
}

func main() {

	f, err := os.Open("us-contacts.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//Create a contact list from the csv data
	contactList := createContactList(data)

	//Sort it by last name
	sort.Slice(contactList, func(i, j int) bool {
		return contactList[i].lastName < contactList[j].lastName
	})

	//Print out every 50th contact
	for i := 0; i < len(contactList); i += 50 {
		fmt.Printf("%+v\n", contactList[i])
	}

}

/*
The contacts are in the form:
First name, Last name, Street, City, State, Zip, Phone, Email.
*/
func createContactList(data [][]string) []Contact {
	var contactList []Contact
	for _, line := range data {
		var contact Contact
		for j, field := range line {
			switch j {
			case 0:
				contact.firstName = field
			case 1:
				contact.lastName = field
			case 2:
				contact.street = field
			case 3:
				contact.city = field
			case 4:
				contact.state = field
			case 5:
				contact.zip = field
			case 6:
				contact.phone = field
			case 7:
				contact.email = field
			}
		}

		contactList = append(contactList, contact)

	}

	return contactList
}
