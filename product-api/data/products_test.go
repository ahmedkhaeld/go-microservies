package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "shoes",
		Price: 1.00,
		SKU:   "abs-deb-wew",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
