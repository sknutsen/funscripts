package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "name",
		Price: 1.00,
		SKU:   "abd-asd-sfd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
