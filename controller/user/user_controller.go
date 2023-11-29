package user

import (
	dbConfig "project-pertama-golang/config"
	"project-pertama-golang/model/user"

	"github.com/gin-gonic/gin"
)

func GetAllUser(context *gin.Context) {
	var users []user.User

	var db = dbConfig.OpenConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_user")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var user user.User

		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.PhoneNumber)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	context.JSON(200, gin.H{
		"status": "success",
		"data":   users,
	})
}

func AddUser(context *gin.Context) {
	var db = dbConfig.OpenConnection()
	defer db.Close()

	var user user.User
	context.Bind(&user)

	_, err := db.Exec("INSERT INTO tbl_user (first_name, last_name, phone_number) VALUES ($1, $2, $3)", user.FirstName, user.LastName, user.PhoneNumber)
	if err != nil {
		panic(err)
	}

	context.JSON(200, gin.H{
		"status": "success",
		"data":   user,
	})
}
