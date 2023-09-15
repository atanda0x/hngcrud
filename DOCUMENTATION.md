#  REST API Documentation
Introduction
This is the documentation for the Person Management API, a simple API for managing person records. This API allows you to perform CRUD (Create, Read, Update, Delete) operations on person data.

## UML
![UML Diagram](./uml.png)

## Base URL
The base URL for all API endpoints is: https://crud-x8cw.onrender.com/api/persons/2

Endpoints
Create a New Person
URL: /
Method: POST
Request Body:
json
`
{
  "id": 2,
  "first_name": "0x",
  "last_name": "atanda"
}`
Response:
Status Code: 201 Created
Response Body:
json
`{
  "id": 2,
  "first_name": "0x",
  "last_name": "atanda"
}`

Get Details of a Person
URL: /{id}
Method: GET
Response:
Status Code: 200 OK
Response Body:
json
`{
  "id": 2,
  "first_name": "0x",
  "last_name": "atanda"
}`

Update Details of an Existing Person
URL: /{id}
Method: PUT
Request Body:
json
`{
  "id": 2,
  "first_name": "0x",
  "last_name": "atanda"
}`

Response:
Status Code: 204 No Content
Remove a Person
URL: /{id}
Method: DELETE
Response:
Status Code: 204 No Content
Dynamic Parameter Handling
The API can dynamically handle parameters for first names and last names. You can perform CRUD operations by specifying the first and last names in the request.

## Sample Usage
Here are some sample API requests and expected responses:

## Create a New Person
Request:

bash
Copy code
curl -X POST -H "Content-Type: application/json" -d '{"first_name": "Alice", "last_name": "Johnson"}' http://localhost:8080/api/persons
Response:

json
Copy code
`{
  "id": 2,
  "first_name": "0x",
  "last_name": "atanda"
}`
Get Details of a Person
Request:

bash
Copy code
curl https://crud-x8cw.onrender.com/api/persons/2
Response:

json
Copy code
`{
  "id": 2,
  "first_name": "0x",
  "last_name": "atanda"
}`
Update Details of an Existing Person
Request:

bash
Copy code
curl -X PUT -H "Content-Type: application/json" -d '{"first_name": "Updated Alice", "last_name": "Updated Johnson"}' https://crud-x8cw.onrender.com/api/persons/2

Response:

bash
Copy code
HTTP/1.1 204 No Content
Remove a Person
Request:

bash
Copy code
curl -X DELETE https://crud-x8cw.onrender.com/api/persons/2
Response:

bash
Copy code
HTTP/1.1 204 No Content
Known Limitations
This API does not include authentication or authorization mechanisms. It's for demonstration purposes and should be enhanced with security features for production use.
Error handling is minimal and can be improved for better error messages.
Setting Up and Running the API
Clone the repository.
Install the required Go packages: go get github.com/gorilla/mux and go get github.com/lib/pq.
Modify the database connection settings in the code (DBHost, DBPort, DBUser, DBPassword, DBName) to match your PostgreSQL configuration.
Run the API: go run main.go
The API will be accessible at https://crud-x8cw.onrender.com/api/persons/2.