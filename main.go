// ABOUT: A REST API for an event booking booking system
//=============================================================================
// 1. It uses "net/http" from the Go Standard Library
// 2. And the Gin web framework
//=============================================================================

package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
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
    
    // context.JSON is an easy way to send a formatted response back to the
    // client.
    // It returns two things:
    // 1. A status code to inform the client that the request was received
    // successfully.
    // 2. A JSON object  
    context.JSON(
        http.StatusOK, 
        gin.H{
            "message": "Hello!",
        },
    )
}
