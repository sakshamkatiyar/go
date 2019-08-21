package mapstore

import (
	"fmt"
	"errors"
	"assignment/2/pkg/domain"
)

type MapStore struct {
	store map[string]domain.Customer // An in-memory store with a map
}

// Factory method gives a new instance of MapStore
// This is for caller packages, not for mapstore
func NewMapStore() *MapStore {
	return &MapStore { store: make(map[string]domain.Customer)}
}

// Implement interface methods of domain.CustomerStore

func (ms *MapStore) Create(c domain.Customer) error {
	if c.ID == "" {
		return errors.New("ID not found for the customer")
	}
	ms.store[c.ID] = c
	fmt.Println("User created successfully, User ID:", c.ID)
	return nil
}

func (ms *MapStore) Delete(ID string) error {
	_, ok := ms.store[ID]
	if ok {
		delete(ms.store, ID)
		fmt.Println("User deleted successfully")
		return nil
	}
	return errors.New("Cannot delete: the customer does not exist")
}

func (ms *MapStore) Update(ID string, c domain.Customer) error {
	_, ok := ms.store[ID]
	if ok {
		ms.store[c.ID] = c
		fmt.Println("User updated successfully, User Name:", c.Name)
		return nil
	}
	return errors.New("Cannot update: the customer does not exist")
}

func (ms *MapStore) GetById(ID string) (domain.Customer, error) {
	val, ok := ms.store[ID]
	if ok {
		fmt.Println("User details:", val.Name, val.Email)
		// return domain.Customer{}, nil
		return ms.store[ID], nil
	}
	return domain.Customer{}, errors.New("The customer does not exist") /// returning nil customer
}

func (ms *MapStore) GetAll() ([]domain.Customer, error) {
	ok := len(ms.store)
	if ok!=0 {
		// dc := make([]domain.Customer, ok) // extra slice in front
		var dc [] domain.Customer
		for c := range ms.store {
			dc = append(dc, ms.store[c])
			// fmt.Println(ms.store[c])
		}	
		return dc, nil
	}
	return nil, errors.New("Database empty")
}