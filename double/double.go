package double

import "errors"

var ErrMissingArgs = errors.New("FirstName and LastName are mandatory arguments")
var ErrNoPersonFound = errors.New("no person found")

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

type Phonebook struct {
	People []*Person
}

type Queryer interface {
	Search(people []*Person, firstName string, lastName string) *Person
}

func (p *Phonebook) Find(query Queryer, firstName, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArgs
	}

	person := query.Search(p.People, firstName, lastName)

	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}
