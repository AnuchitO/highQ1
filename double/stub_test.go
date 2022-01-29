package double

import "testing"

type Stub struct {
	phone string
}

func (s *Stub) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{
		FirstName: "",
		LastName:  "",
		Phone:     s.phone,
	}
}

func TestFindReturnsStubPerson(t *testing.T) {
	phone := "+66 33 222 333"
	phonebook := &Phonebook{}

	s := &Stub{phone}

	p, _ := phonebook.Find(s, "swerwer", "werwerwer")

	if p != phone {
		t.Error("Expected", phone, "got", p)
	}

}
