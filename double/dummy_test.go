package double

import "testing"

type Dummy struct{}

func (ds Dummy) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{}
}

func TestFindReturnsErrorWhenEmptyFirstNameorLastName(t *testing.T) {
	phonebook := &Phonebook{}
	d := &Dummy{}

	_, err := phonebook.Find(d, "", "")

	if err != ErrMissingArgs {
		t.Error("Expected ErrMissingArgs, got", err)
	}
}
