package components

import "slices"

type Contact struct {
	ID    int
	Name  string
	Email string
}

type Contacts []Contact

func (cs Contacts) Init() *Contacts {
	return &Contacts{
		Contact{ID: 0, Name: "hisam", Email: "hisamcode@gmail.com"},
		Contact{ID: 1, Name: "maulana", Email: "mailanacode@gmail.com"},
	}

}

func (cs *Contacts) New(name string, email string) {
	*cs = append(*cs, Contact{
		Name:  name,
		Email: email,
	})

	id := slices.IndexFunc(*cs, func(c Contact) bool {
		return c.Email == email
	})

	(*cs)[id].ID = id
}

// IndexOfByEmail get index by email return -1 if not found
func (cs Contacts) IndexOfByEmail(email string) int {
	for k, v := range cs {
		if v.Email == email {
			return k
		}
	}
	return -1
}

func (cs *Contacts) HasEmail(email string) bool {
	for _, v := range *cs {
		if v.Email == email {
			return true
		}
	}
	return false
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}
