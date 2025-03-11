create a scalable smart home control system with the following components:
	1. a webhook-based golang backend that:
	-	processes voice commands from an alexa custom skill
	-	provides api endpoints for an admin control panel
	-	executes commands on various smart home devices
	-	uses shared business logic for both interfaces
	2.	design the system architecture focusing on:
	-	separation of concerns
	-	code reusability
	-	easy extensibility for new device types
	-	optional state persistence
	3.	implement the http server with:
	-	an alexa webhook endpoint that receives json requests
	-	rest api endpoints for the admin panel
	-	a shared command processing layer
	-	device control abstraction
	4.	consider scalability through:
	-	stateless design
	-	clear separation of command and query operations
	-	adapter pattern for device interfaces
	-	proper error handling and logging
	5.	focus on the backend implementation without requiring a database initially, but design to accommodate one if needed later.
provide code examples for the core components, explaining the key architectural decisions and implementation details.