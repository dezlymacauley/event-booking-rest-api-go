// ABOUT: A REST API for an event booking booking system
//=============================================================================
// 1. It uses "net/http" from the Go Standard Library
// 2. And the Gin web framework
//=============================================================================

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// This is how to use code from the package
	"github.com/dezlymacauley/event-booking-rest-api-go/models"
)

func main() {

	// `gin.Default()` creates an Engine object,
	// which is essentially new web server that is created
	// using the Gin framework.
	//
	// The web server that is created already has
	// two built-in helpers (called middleware).
	//
	// 1. A Logger middleware for logging requests
	// 2. A Recovery middleware for handling errors so that the sever
	// does not crash if something goes wrong.
	//
	// `gin.Default()` returns `*gin.Engine` which is
	// a pointer to an instance of an Engine object (the web server),
	// and that pointer is assigned to the variable `server`.
	//
	// This makes it easy to control the web server by just using the
	// variable called `server`
	//
	// Also note that at this stage the web server is not running.
	server := gin.Default()

	// If someone makes a GET request to the server,
	// E.g. someone searches for http://localhost/events,
	// then the server should run the `getEvents` function to handle that
	// request.
	server.GET("/events", getEvents)

    server.POST("/events", createEvent)

	// This starts the server on port 8080
	// So that full address is http://localhost:8080
	// localhost is an alias for http://127.0.0.1
	server.Run(":8080")

}

// This is a function that is created to handle
// a GET request to the address http:localhost:8080/events
// The getEvents function receives a pointer to a Context object
// So basically when the client called this code
// server.GET("/events", getEvents), reference to the context
// of their request is given to the getEvents function for two reasons:
// 1. The getEvents function receives all of the information
// needed to process the request.
// 2. The getEvents function also gets access access to the `Context` object,
// which is essentially a toolbox of dot methods like `context.JSON`
// that can be used to easily read request information,
// and respond back to the client.
func getEvents(context *gin.Context) {

    events := models.GetAllEvents()
	// context.JSON is an easy way to send a formatted response back to the
	// client.
    // The http status code is set. http.StatusOK is an alias for 200
    // This tells the client (browser, frontend app, or API consumer) 
    // that everything went well.
    // The events slice is marshalled (converted) to a JSON object
    // Under the hood, Gin uses Go’s encoding/json package 
    // to do this transformation automatically.
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
    // The purpose of this function to process the data sent by the client
    // and use it to create a new event.
    // 1. The client data will be received as a JSON object
    // 2. The client data need to be deserialized (Converted from JSON
    // to a Go struct)

    // Step 1: Create an empty instance of the Event struct to store
    // the JSON data from the client, into Go code
    var event models.Event

    // This accepts a pointer to an instance of an empty struct
    // The Gin framework will convert the JSON to a Go Struct
    // and update the event struct
    err := context.ShouldBindJSON(&event)

    // In order for this to work the structure of the JSON object sent
    // by the client must match the the format of the GO struct:

    /*
        Go struct format:

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

    */

    // Error handling
    if err != nil {
        context.JSON(
            http.StatusBadRequest, 
            gin.H{"message": "Could not parse request data"},
        )
        return
    }

    // dummy values for now
    event.ID = 1
    event.UserID = 1

    context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}
