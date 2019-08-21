package main

import "fmt"

var data map[string]string

func init() {
	data = make(map[string]string) // Initialise map with make
}

func Add(k,v string) {
	// ToDo: Check if key exists
	fmt.Println("\nAdd")
	if _, ok := data[k]; ok {
		fmt.Println("Key already exists!")
	} else {
		data [k] = v
		fmt.Println("Add successful!")
	}
}

func Update(k,v string) {
	// ToDo: Check if key exists
	fmt.Println("\nUpdate")
	if _, ok := data[k]; ok {
		data [k] = v
		fmt.Println("Key updated!")
	} else {
		fmt.Println("Key doesn't exists!")
	}
}

func Delete(k string) {
	// ToDo: Check if key exists
	fmt.Println("\nDelete")
	if _, ok := data[k]; ok {
		delete(data, k)
		fmt.Println("Delete successful!")
	} else {
		fmt.Println("Key doesn't exists!")
	}
}

func Get(k string) {
	fmt.Println("\nGet")
	if v, ok := data[k]; ok {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	} else {
		fmt.Println("Key doesn't exists!")
	}
}

func GetAll() {
	fmt.Println("\nGetAll")
	for k, v := range data {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}
}

func main() {
	fmt.Println("Assignment 1")
	Add("7", "seven")
	Update("7", "Seven")
	Delete("6")
	Get("6")
	GetAll()
} 