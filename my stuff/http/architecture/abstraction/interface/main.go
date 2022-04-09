package main

import "fmt"

// contact model
type contact struct {
	firstName string
	lastName  string
	phone     string
}

// mock databases
type mongo map[int]contact
type postgres map[int]contact

// set and get mongo database methods, not accessed directly
func (m mongo) set(n int, c contact) {
	m[n] = c
}
func (m mongo) get(n int) contact {
	return m[n]
}

// set and get postgres database methods, not accessed directly
func (pg postgres) set(n int, c contact) {
	pg[n] = c
}
func (pg postgres) get(n int) contact {
	return pg[n]
}

// interact with databases via db interface methods
type db interface {
	set(n int, c contact)
	get(n int) contact
}

func set(db db, n int, c contact) {
	db.set(n, c)
}

func get(db db, n int) contact {
	return db.get(n)
}

func main() {
	// instantiate mock databases
	dbm := mongo{}
	dbp := postgres{}
	// create some contacts
	firstContact := contact{
		firstName: "John",
		lastName:  "Doe",
		phone:     "555-555-5555",
	}
	secondContact := contact{
		firstName: "Jane",
		lastName:  "Doe",
		phone:     "555-555-5551",
	}
	// set contacts in mock databases
	set(dbm, 1, firstContact)
	set(dbp, 1, secondContact)
	// get contacts from mock databases
	fmt.Println(get(dbm, 1))
	fmt.Printf("%T\n", dbm)
	fmt.Println(get(dbp, 1))
	fmt.Printf("%T\n", dbp)
}
