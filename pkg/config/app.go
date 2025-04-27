// // package config

// // import(
// // 	"github.com/jinzhu/gorm"
// // 	_ "github.com/jinzhu/gorm/dialects/mysql"
// // )
// // var(
// // 	db * gorm.DB
// // )

// // func Connect(){ // helps to create a connection with mysql
// // 	d, err:= gorm.Open("","")
// // 	if err!= nil{
// // 		panic(err)
// // 	}
// // 	db=d
// // }

// // func GetDB() *gorm.DB{
// // 	return db
// // }

// package config

// import (
//     "fmt"
//     "github.com/jinzhu/gorm"
//     _ "github.com/jinzhu/gorm/dialects/mysql"
// )

// var db *gorm.DB

// func Connect() {
//     // user:password@tcp(host:port)/database?charset=utf8&parseTime=True&loc=Local
//     dsn := "root:@tcp(127.0.0.1:3306)/bookstore_db?charset=utf8&parseTime=True&loc=Local"
//     d, err := gorm.Open("mysql", dsn)
//     if err != nil {
//         panic(fmt.Sprintf("Failed to connect to database: %v", err))
//     }
//     db = d
// }

// func GetDB() *gorm.DB {
//     return db
// }

package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
    // Load .env file (silently ignores if it doesn’t exist)
    _ = godotenv.Load()

    user := os.Getenv("DB_USER")       // e.g. "root"
    pass := os.Getenv("DB_PASSWORD")   // e.g. "s3cr3t"
    host := os.Getenv("DB_HOST")       // e.g. "127.0.0.1"
    port := os.Getenv("DB_PORT")       // e.g. "3306"
    name := os.Getenv("DB_NAME")       // e.g. "bookstore_db"

    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        user, pass, host, port, name,
    )

    conn, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }
    db = conn
}

func GetDB() *gorm.DB {
    return db
}