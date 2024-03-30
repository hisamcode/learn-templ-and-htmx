package components

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

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

type Pagination struct {
	Page  int
	Limit int
	Total int
}

func NewPagination(page, limit int, total *int) *Pagination {
	return &Pagination{
		Page:  page,
		Limit: limit,
		Total: *total,
	}

}

func (p *Pagination) Offset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return (p.Page * p.Limit) - p.Limit
}

var ID int = 0

type Contact struct {
	ID                 int
	Name, Email, Phone string
}

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
	Total int
}

func NewContacts() *Contacts {
	contacts := new(Contacts)

	var count int = 0
	for i := 0; i < 20; i++ {
		contacts.Data = append(contacts.Data, *NewContact(
			fmt.Sprintf("user %d", i),
			fmt.Sprintf("user-%d@email.com", i),
			fmt.Sprintf("0080802340%d", i),
		))
		count++
	}

	contacts.Total = count
	return contacts
}

// CRUD

// if found return copyan nya
// if not return nil
func (cs Contacts) FindByID(id int) *Contact {
	for _, v := range cs.Data {
		if v.ID == id {
			return &v
		}
	}

	return nil

}

func (cs Contacts) Search(str string) (*Contacts, error) {
	contacts := Contacts{}
	count := 0

out:
	for _, v := range cs.Data {
		if count >= 5 {
			break out
		}
		if strings.Contains(v.Name, str) || strings.Contains(v.Email, str) {
			count++
			contacts.Data = append(contacts.Data, v)
		}
	}

	if count < 1 {
		return nil, errors.New("not found")
	}

	return &contacts, nil
}

// if page > maxPage return nil
// otherwise return *Contacts
func (cs Contacts) Paging(pagination *Pagination) *Contacts {
	contacts := new(Contacts)

	if pagination.Page == 0 {
		pagination.Page = 1
	}

	maxPage := math.Ceil(float64(pagination.Total) / float64(pagination.Limit))

	if pagination.Page > int(maxPage) {
		return nil
	}

	// linear, kalo data banyak bakal lama beut
out:
	for i, v := range cs.Data {
		// page 2, offset = 10
		if i >= pagination.Offset() && i < pagination.Page*pagination.Limit {
			contacts.Data = append(contacts.Data, v)
		}

		if i == (pagination.Page*pagination.Limit)-1 {
			break out
		}
	}

	return contacts
}

func (cs *Contacts) Create(contact *Contact) error {
	cs.Data = append(cs.Data, *NewContact(
		contact.Name, contact.Email, contact.Phone,
	))
	cs.Total++
	return nil
}

func (cs *Contacts) Edit(contact *Contact) error {
	for i, v := range cs.Data {
		if v.ID == contact.ID {
			cs.Data[i].Name = contact.Name
			cs.Data[i].Email = contact.Email
			cs.Data[i].Phone = contact.Phone
		}
	}
	return nil
}

func (cs *Contacts) Delete(id int) error {
	for i, v := range cs.Data {
		if v.ID == id {
			cs.Data = append(cs.Data[:i], cs.Data[i+1:]...)
		}
	}
	return nil
}
