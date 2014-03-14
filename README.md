ltc-server
==========

An HTTP server to handle requests for content. 

This server is meant to handle requests from ltc-client.

##Routes
Here are the routes that are handled by this server:

### /status
The status route will return the status code set by the /failure route. The default is to return "200 OK". 
This route is meant to be a "health" check route for HAPRoxy.

### /failure
The failure route will change the default status code returned by the status route. This method takes a query string "status" where the value will be the new status code returned by the status route.

### /small
Returns the content of small.txt

### /med
Returns the content of med.txt

### /large
Returns the content of large.txt
