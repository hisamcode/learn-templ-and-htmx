package main

import (
	"fmt"
	"slices"

	"github.com/hisamcode/try-htmx/contacts/components"
)

func main() {

	contacts := components.Contacts{}
	contacts = *contacts.Init()
	contacts.New("pintu", "pintu@gmail.com")
	contacts.New("meja", "meja@gmail.com")
	contacts.New("monitor", "monitor@gmail.com")
	contacts.New("keyboard", "keyboard@gmail.com")
	fmt.Println(contacts)

	id := contacts.IndexOfByEmail("meja@gmail.com")
	// contacts = append(contacts[:id], contacts[id+1:]...)
	contacts = slices.Delete(contacts, id, id+1)

	id = contacts.IndexOfByEmail("pintu@gmail.com")
	contacts = slices.Delete(contacts, id, id+1)
	// contacts = append(contacts[:id], contacts[id+1:]...)

	// fmt.Println(contacts)

	id = contacts.IndexOfByEmail("mailanacode@gmail.com")
	contacts = slices.Delete(contacts, id, id+1)

	// fmt.Println(id, id+1)
	// contacts = append(contacts[:id], contacts[id+1:]...)
	fmt.Println(contacts)

}
