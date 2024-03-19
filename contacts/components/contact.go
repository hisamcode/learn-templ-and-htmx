package components

type Contact struct {
	Name  string
	Email string
}

type Contacts []Contact

func (cs Contacts) Init() *Contacts {
	return &Contacts{
		Contact{Name: "hisam", Email: "hisamcode@gmail.com"},
		Contact{Name: "maulana", Email: "mailanacode@gmail.com"},
	}

}

func (cs *Contacts) New(name string, email string) {
	*cs = append(*cs, Contact{
		Name:  name,
		Email: email,
	})
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
