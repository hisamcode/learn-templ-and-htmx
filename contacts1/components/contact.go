package components

import (
	"errors"
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

type FormContact struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormContact() *FormContact {
	return &FormContact{
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

	// for i := 0; i < 100; i++ {
	// 	c.Data = append(c.Data, NewContact(fmt.Sprintf("user %d", i), fmt.Sprintf("email%d@gmail.com", i)))
	// }
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
