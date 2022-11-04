package composition

import "testing"

func TestManagerBirthday(t *testing.T) {
	m := Manager{
		Employee: Employee{
			Name: "Tom",
			Age:  23,
		},
		Department: "Engineering",
	}

	m.Birthday()

	if m.Age != 24 {
		t.Errorf("Expected age to be 24, got %d", m.Age)
	}

	if m.Employee.Age != 24 {
		t.Errorf("Expected age to be 24, got %d", m.Employee.Age)
	}
}
