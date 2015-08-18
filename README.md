# Dima-Kondravotych-Exercise
short coding exercise to help evaluate problem solving skills

Additional question may be provided via issues to this repo. Good luck and have fun! :)

###Create HTTP Rest API:
1. Use echo or gin for web handler 
2. Implement login endpoint with JWT token, and simple middleware that checks header for 'Authorization: Bearer %jwt_token%' in each request. Otherwise return 403 and json struct with error 
3. Implement endpoint that will use oAuth2 authorization for FB, to login and issue access_token
3. Log each request including status code 
4. Implement persistence with MySQL and Gorm (https://github.com/jinzhu/gorm) 
5. Use Goose or other tool of choice for DB migration (https://bitbucket.org/liamstask/goose) 
6. Implement save endpoint for Task object 
7. Implement update endpoint for Task object 
8. Implement get endpoint for Task object 
9. Implement delete endpoint for Task object (just update IsDeleted field)  
10. Use CORS (reply with header Access-Control-Allow-Origin: *) 
11. Add support for OPTION HTTP method for each endpoints  
12. Configure daemon over simple JSON config. Specify path as process flag for daemon. Required params: ListenAddress, DatabaseUri. 
13. Put in comments below description of taken architectural decisions and


###Task:
```
type Task struct {
    Id          int64
    Title       string
    Description string
    Priority    int
    CreatedAt   *time.Time
    UpdatedAt   *time.Time
    CompletedAt *time.Time
    IsDeleted   bool
    IsCompleted bool
}
```

QUICK START
1) Create database with permissions
CREATE DATABASE dkondratovych_task;

CREATE USER 'gopher'@'localhost' IDENTIFIED BY 'go';
GRANT ALL PRIVILEGES ON dkondratovych_task . * TO 'gopher'@'localhost';
FLUSH PRIVILEGES;

2) Install dependencies. Run install.sh

3) Run database migrations. Run goose up

4) Build and install
go build
go install

5) Try next API calls. :)

6) Full list routes with OPTIONS routes you can find in github.com/seesawlabs/Dima-Kondravotych-Exercise/server/router.go

7) Run application with flag -config /path/to/task_config.json

Task requests information

1) DELETE http://localhost:8080/api/v1/tasks

Headers:
Content type: application/json
Authorization: Bearer %token%

Body:
{
  "id": 1
}

Response:
204 - Delete was successful. No body content.
404 - Internal server error.


2) POST http://localhost:8080/api/v1/tasks

Headers:
Content type: application/json
Authorization: Bearer %token%

Body:
{
	"title" : "Dima Test",
	"description" : "Description",
	"priority" : 1,
	"completedAt" : "2015-08-16T18:50:26+07:00",
	"isDeleted": false,
	"isCompleted": true
}

"completedAt" should be formatted according to RFC3339

Response:
201 - Task was added successfully.
Response Body:
{"data":{"url":"/api/v1/tasks/5"},"metadata":null}


404 - Internal server error.


3) GET http://localhost:8080/api/v1/tasks/1

Headers:
Content type: application/json
Authorization: Bearer %token%

Responses
404 - Task is not found
500 - Internal server error

200 - Task was found
{
    data: {
        id: 1,
        Title: "test",
        Description: "",
        Priority: null,
        createdAt: "0001-01-01T00:00:00Z",
        updatedAt: "2015-08-16T18:05:04+07:00",
        completedAt: "0001-01-01T00:00:00Z",
        IsDeleted: true,
        IsCompleted: false
    },
    metadata: null
}

4) PUT http://localhost:8080/api/v1/tasks

Headers:
Content type: application/json
Authorization: Bearer %token%

Body Example 1:
{
    "taskId": 1,
    "task" : {
    	"title" : "Test",
    	"description" : "Test desc",
    	"priority" : 2,
    	"completedAt" : "2015-08-16T18:50:26+07:00",
    	"isDeleted": true,
    	"isCompleted": false
    }
}

Body Example 2:
{
    "taskId": 1,
    "task" : {
    	"title" : "Test"
    }
}

500 - Internal server error
404 - Task in not found
200 - Task was updated


5) POST http://localhost:8080/auth/jwt

Content type: application/json

Request body
{
    "username": "Bender",
    "password": "molly"
}

Response body
{
    "token":"token_string"
}

6) Facebook auth.
GET http://localhost:8080/auth/facebook

Facebook Redirect URL http://localhost:8080/auth/facebook/callback