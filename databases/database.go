package databases

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GenerateUUID() string {
    return uuid.New().String()
}

func InitDB() {
    dsn := "root:mysql123@tcp(127.0.0.1:3306)/test_api_go?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("[MYSQL] Failed to connect to database:", err)
    }
    db.AutoMigrate(&User{}, &Folder{}, &File{}, &Storage{}) // Migrate the schema
    DB = db
    fmt.Println("[MYSQL] Database connected")
}

// CRUD operations for User
func CreateUser(tmp *TempUser) (User, error) {
    newUser := User{
        UUID:     GenerateUUID(),
        Username:  tmp.Username,
        Email:     tmp.Email,
        Masterkey: tmp.Masterkey,
        Storage:   Storage{UUID: GenerateUUID(), UsedSpace: 0, TotalSpace: tmp.Storage, BonusSpace: 0},
    }

    result := DB.Create(&newUser)
    return newUser, result.Error
}

func ReadUser(user *User, id string) error {
    result := DB.Where("uuid = ?", id).First(user)
    return result.Error
}

// Change the user's email
func UpdateUser(user *User, email string) error {
    result := DB.Model(user).Where("uuid = ?", user.UUID).Update("email", email)
    return result.Error
}

func DeleteUser(user *User) error {
    result := DB.Delete(user)
    return result.Error
}



