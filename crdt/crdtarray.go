package crdt

import (
	"fmt"
	"sync"
)

// Element represents an element in the CRDTArray.
type Element struct {
	ID    string
	Value interface{}
}

// CRDTArray is a conflict-free replicated data type for an array.
type CRDTArray struct {
	clientID string
	counter  int
	data     []Element
	lock     sync.RWMutex
}

// NewCRDTArray creates a new CRDTArray with a given client ID.
func NewCRDTArray(clientID string) *CRDTArray {
	return &CRDTArray{
		clientID: clientID,
		data:     make([]Element, 0),
	}
}

// Add appends a new element to the CRDTArray.
func (a *CRDTArray) Add(value interface{}) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.counter++
	element := Element{
		ID:    fmt.Sprintf("%s-%d", a.clientID, a.counter),
		Value: value,
	}
	a.data = append(a.data, element)
}

// Merge merges another CRDTArray into this one, maintaining order and uniqueness.
func (a *CRDTArray) Merge(other *CRDTArray) {
	a.lock.Lock()
	defer a.lock.Unlock()

	for _, element := range other.data {
		if !a.contains(element.ID) {
			a.data = append(a.data, element)
		}
	}
}

// contains checks if an element ID is already in the array.
func (a *CRDTArray) contains(id string) bool {
	for _, element := range a.data {
		if element.ID == id {
			return true
		}
	}
	return false
}
