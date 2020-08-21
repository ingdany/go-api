package connect

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"

	"../structures"
)

var connection *gorm.DB
var err error
var err2 error
//var connection *sql.DB

const engine_sql string = "mysql"

const username string = "root"
const password string = "M1P@$$w0rd"
const database string = "test1"

func InitializeDatabase() {
	ConnectORM(CreateString())
	log.Println("Connection has been successfully!")
}

func CloseConnection() {
	connection.Close()
	log.Println("Connection has been closed")
}

func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open( engine_sql, stringConnection )
	if err != nil {
		log.Fatal(err)
		//return nil
	}
	return connection
}

func GetUser(id string) structures.User {
	user := structures.User{}	
	var db *gorm.DB
	db = ConnectORM(CreateString())
	db.Where("id = ?", id).First(&user)
	return user
}

func CreateUser(user structures.User) structures.User {
	var db *gorm.DB
	db = ConnectORM(CreateString())
	db.Create(&user)
	return user
}

func CreateString() string {
	return username + ":" + password + "@/" + database
}

func UpdateUser(id string, user structures.User ) structures.User {
	currentUser := structures.User{}
	var db *gorm.DB
	db = ConnectORM(CreateString())
	db.Where("id = ?", id).First(&currentUser)
	currentUser.Username = user.Username
	currentUser.First_Name = user.First_Name
	currentUser.Last_Name = user.Last_Name
	db.Save(&currentUser)
	return currentUser
}

func DeleteUser(id string ) {
	currentUser := structures.User{}
	var db *gorm.DB
	db = ConnectORM(CreateString())
	db.Where("id = ?", id).First(&currentUser)
	db.Delete(&currentUser)
}
