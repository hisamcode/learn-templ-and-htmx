package components

import (
	"errors"
	"fmt"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Email string
}

var ID = 0

func NewContact(name string, email string) Contact {
	ID++
	return Contact{
		ID:    ID,
		Name:  name,
		Email: email,
	}

}

type Form struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormContact() *Form {
	return &Form{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Contacts struct {
	Data []Contact
}

func (c *Contacts) Init() {
	c.Data = append(c.Data, NewContact("hisam", "hisam@gmail.com"))
	c.Data = append(c.Data, NewContact("maulana", "maulana@gmail.com"))
	c.Data = append(c.Data, NewContact("pintu", "pintu@gmail.com"))

	for i := 0; i < 20; i++ {
		c.Data = append(c.Data, NewContact(fmt.Sprintf("user %d", i), fmt.Sprintf("email%d@gmail.com", i)))
	}
}

func (c *Contacts) New(name string, email string) {
	c.Data = append(c.Data, NewContact(name, email))
}

func (c *Contacts) DeleteByID(id int) {
	index := c.IndexOfByID(id)
	c.Data = append(c.Data[:index], c.Data[index+1:]...)
}

func (c *Contacts) UpdateByID(id int, contact *Contact) error {
	index := c.IndexOfByID(id)
	if index == -1 {
		return errors.New("index not found")
	}
	c.Data[index].Email = contact.Email
	c.Data[index].Name = contact.Name

	return nil
}

func (c Contacts) IndexOfByID(id int) int {
	for k, v := range c.Data {
		if v.ID == id {
			return k
		}
	}

	return -1
}

func (c *Contacts) FindByID(id int) (*Contact, error) {
	for _, v := range c.Data {
		if id == v.ID {
			return &v, nil
		}
	}

	return nil, errors.New("not found")
}

func (c *Contacts) FindByEmail(email string) (*Contact, error) {
	for _, v := range c.Data {
		if email == v.Email {
			return &v, nil
		}
	}

	return nil, errors.New("not found")
}

func (c Contacts) All(page int) Contacts {
	contacts := Contacts{}
	showOnThePage := 10
	offset := 0
	if page > 1 {
		// bug ketika lebih dari total c.Data
		if (page-1)*showOnThePage > len(c.Data) {
			return contacts
		}
		offset = (page * showOnThePage) - showOnThePage
	}
	// awal nya soalnya 0
	if page <= 1 {
		page = 1
	}

	contacts.Data = append(contacts.Data, c.Data[offset:page*showOnThePage]...)
	return contacts
}

func (c Contacts) Search(search string) Contacts {
	contacts := Contacts{}
	count := 0
out:
	for _, v := range c.Data {
		if strings.Contains(v.Email, search) || strings.Contains(v.Name, search) {
			if count > 5 {
				break out
			}
			contacts.Data = append(contacts.Data, v)
			count++
		}
	}
	return contacts
}

func (c Contacts) Count() int {
	return len(c.Data)

}

func (c Contacts) Bytes() *[]byte {

	b := []byte{}

	for i, v := range c.Data {
		b = append(b, []byte(fmt.Sprintf("id=%d|name=%s|email=%s", v.ID, v.Name, v.Email))...)
		if i < len(c.Data)-1 {
			b = append(b, []byte("\n")...)
		}
	}

	return &b
}
