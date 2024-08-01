// main.go
package main

import (
    "fmt"
	"api-go/databases"
)

func main() {
    // Initialize the database
    databases.InitDB()

    // Migrate the schema
    databases.DB.AutoMigrate(&databases.User{})

    // Create a new user
    user := databases.User{Name: "John Doe", Email: "johndoe@example.com", MasterKey: "123456"}
    databases.DB.Create(&user)

    // Read the user back
    var readUser databases.User
    databases.DB.First(&readUser, user.ID)
    fmt.Println("User:", readUser.Name, readUser.Email)
}
