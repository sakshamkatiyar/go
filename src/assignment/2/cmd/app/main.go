package main

import (
	"fmt"
	"assignment/2/pkg/domain"
	"assignment/2/pkg/mapstore"
)

type CustomerController struct {
	// Explicit dependency and declarative programming that hides dependent logic
	store domain.CustomerStore // It can be any Store including MapStore
}

func (cc CustomerController ) Add (c domain.Customer) {
	err:= cc.store.Create(c)
	if err!=nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

// Delete deletes a customer from the store.
func (cc CustomerController) Delete(c domain.Customer) {
	err := cc.store.Delete(c.ID)
	if err != nil {
		fmt.Println("Could not delete a customer:", err)
		return
	}
	fmt.Println("New Customer has been deleted")
}

func main() {
	controller := CustomerController{
		store : mapstore.NewMapStore(), // Inject the dependency
		// store := mongodb.NewMongoStore()
	}

	customer := domain.Customer {
		ID : "cust101",
		Name: "JP Morgan",
		Email: "abc@xyz.com",
	}


	err := controller.store.Create(customer)

	if err!=nil{
		fmt.Println("Cannot create customer:", err)
	}

	err = controller.store.Update(customer.ID, customer)

	if err!=nil{
		fmt.Println("Cannot update customer:", err)
	}

	// err = controller.store.Delete(customer.ID)

	// if err!=nil{
	// 	fmt.Println("Cannot delete customer:", err)
	// }
	
	_, err = controller.store.GetById(customer.ID)

	if err!=nil{
		fmt.Println("Cannot print customer:", err)
	}

	val, err := controller.store.GetAll()

	if err!=nil{
		fmt.Println("Cannot print database:", err)
	} else {
		fmt.Println(val)
	}

	// controller.Add(customer) // Create new Customer
}