package controller

import (
	"fmt"
	// "mux"
	"net/http"
	"assignment/3/pkg/domain"
	"assignment/3/pkg/mapstore"
)

type CustomerController struct {
	// Explicit dependency and declarative programming that hides dependent logic
	store domain.CustomerStore // It can be any Store including MapStore
}

var controller = CustomerController {
	store : mapstore.NewMapStore(), // Inject the dependency
	// store := mongodb.NewMongoStore()
}


func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	html :=
		`<doctype html>
        <html>
	<head>
		<title>Assignment 3</title>
	</head>
	<body>
		<b>Hello Gopher!</b>
        <p>
            <a href="/create">Create</a> |  <a href="/update">Update</a> |  <a href="/delete">Delete</a> |  
            <a href="/getbyid">Get By ID</a> |  <a href="/getall">Get All</a>
        </p>
	</body>
</html>`
	fmt.Fprintf(w, html)
}

func Add (w http.ResponseWriter, r *http.Request) {

	customer := domain.Customer {
		ID : "cust101",
		Name: "JP Morgan",
		Email: "abc@xyz.com",
	}

	err := controller.store.Create(customer)

	if err!=nil {
		fmt.Println("Cannot create customer:", err)
		fmt.Fprintf(w, "Cannot create customer!")
		return
	}

	str := "New customer successfully created!\n\nCustommer ID: " + customer.ID + "\nCustomer Name: " + customer.Name + "\nCustomer Email: " + customer.Email
	fmt.Fprintf(w, str)
}

func Update (w http.ResponseWriter, r *http.Request) {

	customer := domain.Customer {
		ID : "cust101",
		Name: "HG Wells",
		Email: "read@book.com",
	}

	err := controller.store.Update(customer.ID, customer)

	if err!=nil {
		fmt.Println("Cannot update customer:", err)
		fmt.Fprintf(w, "Cannot update customer!")
		return
	}

	str := "Customer successfully updated!\n\nCustommer ID: " + customer.ID + "\nCustomer Name: " + customer.Name + "\nCustomer Email: " + customer.Email
	fmt.Fprintf(w, str)
}

// Delete deletes a customer from the store.
func Delete (w http.ResponseWriter, r *http.Request) {

	customer := domain.Customer {
		ID : "cust101",
	}

	err := controller.store.Delete(customer.ID)

	if err!=nil {
		fmt.Println("Cannot delete customer:", err)
		fmt.Fprintf(w, "Cannot delete customer!")
		return
	}

	str := "Customer successfully deleted!\n\nCustomer ID: " + customer.ID
	fmt.Fprintf(w, str)
}

func GetById (w http.ResponseWriter, r *http.Request) {

	customer := domain.Customer {
		ID : "cust101",
	}

	val, err := controller.store.GetById(customer.ID)

	if err!=nil {
		fmt.Println("Cannot print customer:", err)
		fmt.Fprintf(w, "Cannot print customer!")
		return
	}

	str := "Customer successfully printed!\n\nCustommer ID: " + val.ID + "\nCustomer Name: " + val.Name + "\nCustomer Email: " + val.Email
	fmt.Fprintf(w, str)
}

func GetAll (w http.ResponseWriter, r *http.Request) {

	val, err := controller.store.GetAll()

	if err!=nil {
		fmt.Println("Cannot print database:", err)
		fmt.Fprintf(w, "Cannot print database!")
	} else {
		str := "All Customer successfully printed!\n"
		for _,v := range val {
			str += "\nCustomer ID: " + v.ID + "\nCustomer Name: " + v.Name + "\nCustomer Email: " + v.Email
		}
		// fmt.Println(val)
		fmt.Fprintf(w, str)
	}
}
