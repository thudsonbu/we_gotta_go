package types

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Methods that modify the value of a struct should take a pointer as a
// parameter.
func (p *Person) Birthday() {
	p.Age++
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// Methods that take a pointer can determing if a value is nil by checking
// if the pointer is nil.
func (p *Person) IsNil() bool {
	return p == nil
}

// Go supports subtyping but not inhertiance. Thought HighScore is a subtype of
// Score it cannot be used or compared to a Score without an explicit
// conversion.
type Score struct {
  score int
  name string
}

type HighScore Score

// HighScore will also not include any methods defined on Score.
func (s *Score) increment() int {
  s.score += 1

	return s.score
}

func Main() {
	tom := Person{
		Name: "Tom",
		Age:  23,
	}

	tom.Birthday()

	fmt.Println(tom.Age) // 24
}
