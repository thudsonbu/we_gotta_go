package nullobject

import "fmt"

// AbstractObject is the interface for all objects
type AbstractObject interface {
	IsNil() bool
	GetName() string
}

// RealObject is a struct representing a real object
type RealObject struct {
	name string
}

func (realObject *RealObject) IsNil() bool {
	return false
}

func (realObject *RealObject) GetName() string {
	return realObject.name
}

// NullObject is a struct representing a null object
type NullObject struct{}

func (nullObject *NullObject) IsNil() bool {
	return true
}

func (nullObject *NullObject) GetName() string {
	return "Not Available"
}

// ObjectFactory is a struct for creating real and null objects
type ObjectFactory struct {
	objects []string
}

func (objectFactory *ObjectFactory) GetObject(index int) AbstractObject {
	if index < len(objectFactory.objects) {
		return &RealObject{objectFactory.objects[index]}
	}
	return &NullObject{}
}

func main() {
	objectFactory := &ObjectFactory{
		objects: []string{"Object 1", "Object 2"},
	}

	for i := 0; i < 5; i++ {
		object := objectFactory.GetObject(i)
		fmt.Printf("Name: %s\n", object.GetName())
	}
}
