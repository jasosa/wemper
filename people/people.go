package people

// Person represents a person in the most simple way
type Person struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

//People represents a collection of sinlge persons
var People []Person
