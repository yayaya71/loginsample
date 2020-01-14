package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
}

func main() {

	db, err := sql.Open("mysql", "ya_k/unit2_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
 
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
 
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.ID, user.Name)
	}
 
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "login.html", gin.H{})

	})

	router.Run(":8080")

}
