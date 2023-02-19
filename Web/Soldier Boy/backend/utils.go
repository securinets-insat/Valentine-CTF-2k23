package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Task struct {
	gorm.Model
	Title       string
	Description string
	Url         string
	Flag        string
}

func GetDb() *gorm.DB {
	dsn := "host=REDACTED"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Task{})
	return db
}

// func CORSMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     }
// }