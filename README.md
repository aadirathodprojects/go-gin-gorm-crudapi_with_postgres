# go-gin-gorm-crudapi_with_postgres

==============================
CRUD API with Gin and GORM (Go)
==============================

This document explains, step-by-step, the concepts and functions used to build a simple CRUD API using the Gin web framework and GORM ORM in Golang.

-----------------------------
1. Project Structure Overview
-----------------------------

go-gin-gorm/
├── config/        -> Database connection config (db.go)
├── controller/    -> CRUD logic (user handler functions)
├── models/        -> Structs (models) that map to DB tables
├── main.go        -> Application entry point, routes defined here

---------------------------
2. Required Go Libraries
---------------------------

- "github.com/gin-gonic/gin"  -> Web framework
- "gorm.io/gorm"              -> ORM for DB handling
- "gorm.io/driver/postgres"   -> PostgreSQL driver

-------------------------
3. Model (models/user.go)
-------------------------

This defines the `User` struct mapped to the DB table.

```
type User struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}
```

- `json:"..."` => Ensures the fields bind with JSON keys from request/response.

----------------------------------
4. Database Setup (config/db.go)
----------------------------------

Connect to PostgreSQL using GORM:

```
dsn := "host=localhost user=postgres password=root dbname=Practice port=5432 sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

Auto migrate the model:
```
db.AutoMigrate(&models.User{})
```

-------------------------------------------
5. Routes Setup (main.go or routes.go file)
-------------------------------------------

```
r := gin.Default()
r.POST("/user", controller.CreateUser)
r.GET("/user", controller.GetUsers)
r.PUT("/user/:id", controller.UpdateUser)
r.DELETE("/user/:id", controller.DeleteUser)
r.Run(":8082")
```

- `POST`    => Create a new user.
- `GET`     => Fetch all users.
- `PUT`     => Update a specific user.
- `DELETE`  => Delete a user by ID.

----------------------------------------
6. CRUD Controller Logic (controller/)
----------------------------------------

### CreateUser

```go
var user models.User
if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}
config.DB.Create(&user)
c.JSON(http.StatusOK, user)
```

- `ShouldBindJSON()` binds incoming JSON to Go struct.
- `config.DB.Create()` inserts into DB.

### GetUsers

```go
var users []models.User
config.DB.Find(&users)
c.JSON(http.StatusOK, users)
```

- `Find()` fetches all records.

### UpdateUser

```go
id := c.Param("id")
var user models.User
config.DB.First(&user, id)

var updatedData models.User
c.BindJSON(&updatedData)
user.Name = updatedData.Name
...
config.DB.Save(&user)
```

- `First()` gets the user by ID.
- `Save()` updates the DB record.

### DeleteUser

```go
id := c.Param("id")
config.DB.First(&user, id)
config.DB.Delete(&user)
c.JSON(http.StatusOK, gin.H{"message": "successful deletion"})
```

----------------------------
7. Gin Handler Functions
----------------------------

| Function                | Purpose |
|------------------------|---------|
| c.Param("id")          | Gets path parameter like /user/:id |
| c.ShouldBindJSON(&obj) | Safely bind JSON request body |
| c.JSON(status, data)   | Send a JSON response |
| c.String(...)          | Return plain string response |
| c.Query("q")           | Get query string parameter |
| c.PostForm("key")      | Get form field value |

--------------------------
8. GORM DB Methods Used
--------------------------

| Method        | Purpose |
|---------------|---------|
| Create(&obj)  | Insert record |
| Find(&slice)  | Get all records |
| First(&obj, id)| Get single record by ID |
| Save(&obj)    | Update existing record |
| Delete(&obj)  | Remove record from DB |

---------------------------
9. Response Best Practices
---------------------------

- Use proper HTTP status codes (200 OK, 400 Bad Request, 404 Not Found, 500 Internal Server Error)
- Always send JSON with clear message or error keys

-----------------------------------
10. Build & Run the Application
-----------------------------------

```bash
go run main.go
```

Test with Thunder Client, Postman, or curl:
- `POST /user` with JSON body
- `GET /user`
- `PUT /user/:id`
- `DELETE /user/:id`


Important Gin Functions You Should Know
Here’s a list of commonly used Gin functions and methods, grouped by category and with comments:

🔹 1. Request Handling
Function	Description
c.Param("name")	Gets a route/path parameter like /user/:id.
c.Query("name")	Gets query string value like /user?name=xyz.
c.PostForm("key")	Reads form field from a form submission.
c.ShouldBindJSON(&obj)	Safely binds JSON payload to a struct.
c.BindJSON(&obj)	Similar to above but returns 400 on error automatically.
c.Bind(&obj)	Automatically detects content type (JSON, form, etc.).

🔹 2. Response Handling
Function	Description
c.JSON(statusCode, data)	Sends JSON response.
c.String(statusCode, "text")	Sends plain text.
c.HTML(statusCode, "template.html", gin.H{})	Renders HTML templates (when using HTML rendering).
c.Redirect(code, "/newpath")	Redirects to another route.

🔹 3. Middleware & Context
Function	Description
c.Next()	Executes the next middleware in the chain.
c.Abort()	Stops execution and doesn’t call remaining middleware.
c.Set("key", value)	Stores data in context (can be used between middleware and handlers).
c.Get("key")	Retrieves stored data from context.

🔹 4. Router Setup (usually in main.go or routes/user.go)
Syntax	Purpose
router.GET("/path", handlerFunc)	Handle HTTP GET
router.POST("/path", handlerFunc)	Handle HTTP POST
router.PUT("/path", handlerFunc)	Handle HTTP PUT
router.DELETE("/path", handlerFunc)	Handle HTTP DELETE
router.Group("/api")	Grouping routes under a prefix
router.Use(MiddlewareFunc)	Applying global middleware

🔹 5. Error Handling
go
Copy
Edit
if err := c.ShouldBindJSON(&user); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}
Gin promotes graceful error handling and helps you return appropriate HTTP status codes.

🔹 6. Advanced: Custom Middleware (Optional)
You can create your own middleware, e.g. for logging or authentication:

go
Copy
Edit
func LoggerMiddleware(c *gin.Context) {
	fmt.Println("Request:", c.Request.URL.Path)
	c.Next() // Call the next handler
}
Register it like:

go
Copy
Edit
router.Use(LoggerMiddleware)
✅ You're Currently Using These (from your code):
Function	You're Using
c.ShouldBindJSON()	✅
c.JSON()	✅
c.Param()	✅
c.BindJSON()	✅
c.Status()	❌ (Optional if you just want to set status without data)

==========================
End of Documentation ✅
==========================
