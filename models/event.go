package models

import "time"

// The structure of data for each event
type Event struct {
     ID int
     Name string
     Description string
     Location string
     DateTime time.Time 
     // An id that links the user who created the event 
     // to the id ID of the event
     UserID int
}

// A slice of events
// Basically a data structure that will be a list of structs on an event
// The list is currently empty
var events = []Event{}

// Methods that will be availabl to each instance 
// that is created from a struct

// This will save an event to the database 
// (e Event) means that this is not a standalone function but a method that
// is available to any instances created from the struct
// `e` is just a naming convention in Go that means an instance of a struct
// The letter chosen is the first letter of the struct name, in lowercase.
func (e Event) Save() {
    events = append(events, e)
}

//_____________________________________________________________________________

// NOTE: Make sure that the function name here starts with a Capital letter

// The function will return the variable events
func GetAllEvents() []Event {
    return events
}

//_____________________________________________________________________________
