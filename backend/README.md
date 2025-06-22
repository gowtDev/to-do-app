# readme
# Todo List API
This is a practice project for [Todo List API](https://roadmap.sh/projects/todo-list-api).

## Usage

### Get Started

```bash
git clone https://github.com/gowtDev/to-do-app.git
cd to-do-app
go build -o main
go run main.go  # and the server will run on localhost:8080
```

# apis
1.register
2.login

1.createtodo
2.updatetodo
3.deletetodo
4.gettodo

### User Registration

# Request:
curl -X POST http://localhost:8080/iam/register -H "Content-Type: application/json" -d "{\"name\":\"David\",\"email\":\"david@gmail.com\",\"password\":\"password\"}"

# Response:
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhdmlkQGdtYWlsLmNvbSIsImV4cCI6MTc1MDI1OTI2MCwiaXNzdWVkX2F0IjoiMjAyNS0wNi0xN1QyMDozNzo0MC4wODU5NzU5KzA1OjMwIn0.2CHS9VdmH7bZ8yYwXfOCnsDhaHZ5IUGn7PFTCKwmMpg"}

### User Login

# Request:
curl -X POST http://localhost:8080/iam/login -H "Content-Type: application/json" -d "{\"email\":\"david@gmail.com\",\"password\":\"password\"}"

# Response:
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhdmlkQGdtYWlsLmNvbSIsImV4cCI6MTc1MDI1OTMyMywiaXNzdWVkX2F0IjoiMjAyNS0wNi0xN1QyMDozODo0My41NTU1Nzk2KzA1OjMwIn0.3zIRJODnu5KOAiZNu_UZaVfuQt-mDn7wN1pdgfzkPXU"}

### Create a Todo Item

# Request:
curl -X POST http://localhost:8080/iam/todos -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhdmlkQGdtYWlsLmNvbSIsImV4cCI6MTc1MDI1OTMyMywiaXNzdWVkX2F0IjoiMjAyNS0wNi0xN1QyMDozODo0My41NTU1Nzk2KzA1OjMwIn0.3zIRJODnu5KOAiZNu_UZaVfuQt-mDn7wN1pdgfzkPXU" -H "Content-Type: application/json" -d "{\"title\":\"Buy snacks\",\"description\":\"Buy fruits, cholocates, and bread\"}"

# Response:
{"ID":1,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:39:31.7127691+05:30"}

### Update a Todo Item

# Request:
curl -X PUT http://localhost:8080/iam/todos/1 -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhdmlkQGdtYWlsLmNvbSIsImV4cCI6MTc1MDI1OTMyMywiaXNzdWVkX2F0IjoiMjAyNS0wNi0xN1QyMDozODo0My41NTU1Nzk2KzA1OjMwIn0.3zIRJODnu5KOAiZNu_UZaVfuQt-mDn7wN1pdgfzkPXU" -H "Content-Type: application/json" -d "{\"title\":\"Buy snacks\",\"description\":\"Buy fruits only\"}"

# Response:
{"ID":1,"Title":"Buy snacks","Description":"Buy fruits only","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:39:31.7127691+05:30"}


### Delete a Todo Item


# Request:
curl -X DELETE http://localhost:8080/iam/todos/1 -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhdmlkQGdtYWlsLmNvbSIsImV4cCI6MTc1MDI1OTMyMywiaXNzdWVkX2F0IjoiMjAyNS0wNi0xN1QyMDozODo0My41NTU1Nzk2KzA1OjMwIn0.3zIRJODnu5KOAiZNu_UZaVfuQt-mDn7wN1pdgfzkPXU"

# Response:

### Get Todo Items


# Request:
curl -X GET "http://localhost:8080/iam/todos?page=1&limit=10" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhdmlkQGdtYWlsLmNvbSIsImV4cCI6MTc1MDI1OTMyMywiaXNzdWVkX2F0IjoiMjAyNS0wNi0xN1QyMDozODo0My41NTU1Nzk2KzA1OjMwIn0.3zIRJODnu5KOAiZNu_UZaVfuQt-mDn7wN1pdgfzkPXU"

# Response:
{"data":[
{"ID":2,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:46.9318207+05:30"},
{"ID":3,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:49.6535163+05:30"},
{"ID":4,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:50.6288509+05:30"},
{"ID":5,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:51.3305418+05:30"},
{"ID":6,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:51.9841687+05:30"},
{"ID":7,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:52.6716424+05:30"},
{"ID":8,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:53.3248869+05:30"},
{"ID":9,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:53.9060129+05:30"},
{"ID":10,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:54.478566+05:30"},
{"ID":11,"Title":"Buy snacks","Description":"Buy fruits, cholocates, and bread","CreatedBy":"david@gmail.com","CreatedAt":"2025-06-17T20:42:55.1322153+05:30"}],
"limit":10,"page":1,"total":10}


# Sample application logs:

go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /todos                    --> github.com/gowtDev/to-do-app/handlers.createTodos (4 handlers)
[GIN-debug] PUT    /todos/:id                --> github.com/gowtDev/to-do-app/handlers.updateTodos (4 handlers)
[GIN-debug] DELETE /todos/:id                --> github.com/gowtDev/to-do-app/handlers.deleteTodos (4 handlers)
[GIN-debug] GET    /todos                    --> github.com/gowtDev/to-do-app/handlers.getTodos (4 handlers)
[GIN-debug] POST   /register                 --> github.com/gowtDev/to-do-app/handlers.createUser (3 handlers)
[GIN-debug] POST   /login                    --> github.com/gowtDev/to-do-app/handlers.login (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2025/06/17 - 20:36:37 | 400 |      3.5409ms |             ::1 | POST     "/register"
[GIN] 2025/06/17 - 20:37:40 | 201 |      9.6938ms |             ::1 | POST     "/register"
[GIN] 2025/06/17 - 20:38:43 | 200 |      2.2184ms |             ::1 | POST     "/login"
[GIN] 2025/06/17 - 20:39:31 | 201 |     14.6868ms |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:40:12 | 401 |      9.8717ms |             ::1 | PUT      "/todos/1"
[GIN] 2025/06/17 - 20:40:20 | 401 |            0s |             ::1 | PUT      "/todos/1"
[GIN] 2025/06/17 - 20:40:57 | 200 |      1.1439ms |             ::1 | PUT      "/todos/1"
[GIN] 2025/06/17 - 20:41:52 | 204 |     12.9132ms |             ::1 | DELETE   "/todos/1"
[GIN] 2025/06/17 - 20:42:46 | 201 |      1.0531ms |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:49 | 201 |            0s |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:50 | 201 |            0s |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:51 | 201 |       583.1µs |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:51 | 201 |            0s |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:52 | 201 |       622.8µs |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:53 | 201 |            0s |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:53 | 201 |            0s |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:54 | 201 |       698.5µs |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:55 | 201 |       532.3µs |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:55 | 201 |      1.0405ms |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:56 | 201 |      1.1146ms |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:56 | 201 |       558.3µs |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:42:57 | 201 |            0s |             ::1 | POST     "/todos"
[GIN] 2025/06/17 - 20:43:05 | 200 |       510.3µs |             ::1 | GET      "/todos?page=1&limit=10"
