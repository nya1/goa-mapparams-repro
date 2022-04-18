package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("calc", func() {
	Title("Calculator Service")
	Description("HTTP service for multiplying numbers, a goa teaser")
	Server("calc", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

// Service describes a service
var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers")
	// Method describes a service method (endpoint)
	Method("multiply", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Attribute("type", Int, "type description")

			Attribute("metadata", MapOf(String, String), "this should be on openapi", func() {
				Example(map[string]string{"testKey": "val"})
			})
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(Int)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			Param("type") // <-- this is present in openapi

			MapParams("metadata") // <!-- this is NOT present in openapi

			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/multiply/")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})
})
