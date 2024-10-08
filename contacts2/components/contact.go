package components

import (
	"fmt"
	"strings"
)

type Pagination struct {
	Page    int
	Limit   int
	MaxPage int
}

func NewPagination() *Pagination {
	return &Pagination{
		Page:    1,
		Limit:   10,
		MaxPage: 0,
	}
}

type Form struct {
	Values map[string]string
	Errors map[string]string
}

func NewForm() *Form {
	return &Form{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var ID int = -1

func NewContact(name, email, phone string) *Contact {
	ID++
	return &Contact{
		ID:    ID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
}

type Contacts struct {
	Data  []Contact
	Count int
}

func (cs *Contacts) Init() {
	for i := 0; i < 20; i++ {
		cs.Add(
			fmt.Sprintf("user %d", i),
			fmt.Sprintf("user%d@gmail.com", i),
			fmt.Sprintf("00000%d", i),
		)
	}
}

func (cs *Contacts) Paging(pagination Pagination) *Contacts {
	contacts := new(Contacts)
	offset := (pagination.Page * pagination.Limit) - pagination.Limit
	// contacts.Data = append(contacts.Data, cs.Data[offset:pagination.Page*pagination.Limit]...)
	for i, v := range cs.Data {
		if i >= offset && i < pagination.Page*pagination.Limit {
			contacts.Data = append(contacts.Data, v)
		}
	}
	return contacts
}

func (cs *Contacts) Add(name, email, phone string) {
	cs.Data = append(cs.Data, *NewContact(name, email, phone))
	cs.Count++
}

func (cs *Contacts) Edit(contact Contact) {
	theContact := cs.FindByID(contact.ID)
	theContact.Name = contact.Name
	theContact.Email = contact.Email
	theContact.Phone = contact.Phone
}

func (cs *Contacts) Delete(id int) {
	i := cs.IndexOf(id)
	cs.Data = append(cs.Data[:i], cs.Data[i+1:]...)
	cs.Count--
}

// if not found return nil
func (cs *Contacts) FindByID(id int) *Contact {
	for i, v := range cs.Data {
		if v.ID == id {
			return &cs.Data[i]
		}
	}
	return nil
}

// if not found return nil
func (cs *Contacts) FindByEmail(email string) *Contact {
	for i, v := range cs.Data {
		if v.Email == email {
			return &cs.Data[i]
		}
	}
	return nil
}

func (cs *Contacts) Search(str string) *Contacts {
	contacts := Contacts{}
	for _, v := range cs.Data {
		if strings.Contains(v.Email, str) || strings.Contains(v.Name, str) || strings.Contains(v.Phone, str) {
			contacts.Data = append(contacts.Data, v)
		}
	}

	return &contacts
}

func (cs *Contacts) IndexOf(id int) int {
	for i, v := range cs.Data {
		if v.ID == id {
			return i
		}
	}
	return -1
}
