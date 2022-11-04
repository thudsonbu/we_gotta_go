package composition

type Employee struct {
	Name string
	Age  int
}

func (e *Employee) Birthday() {
	e.Age++
}

// Fields on the Employee type will be _promoted_ to the Manager type, including
// any methods defined on the Employee type.
type Manager struct {
	Employee
	Department string
}
