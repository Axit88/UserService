package main

import (
	"database/sql"
	"net/http"

	"github.com/Axit88/UserService/src/config"
	"github.com/Axit88/UserService/src/domain/userService/core/model"
	"github.com/Axit88/UserService/src/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func SetUpDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/UserService")
	return db, err
}

func GetUser(context *gin.Context) {
	db, err := SetUpDb()
	if err != nil {
		context.JSON(http.StatusInternalServerError, "DatabaseConnection Failed")
		return
	}

	id := context.Param("id")
	var user model.User
	err = db.QueryRow("SELECT UserId,UserName FROM USER WHERE UserId = ?", id).Scan(&user.UserId, &user.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User Not Found"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func AddUser(context *gin.Context) {
	db, err := SetUpDb()
	if err != nil {
		context.JSON(http.StatusInternalServerError, "DatabaseConnection Failed")
		return
	}

	var newUser model.User
	err = context.BindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json Payload"})
		return
	}

	insert, err := db.Query("INSERT INTO USER VALUES (?,?)", newUser.UserId, newUser.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Query Execution Failed"})
		return
	}
	defer insert.Close()

	context.JSON(http.StatusCreated, newUser) // (status , JSON)
}

func DeleteUser(context *gin.Context) {
	db, err := SetUpDb()
	if err != nil {
		context.JSON(http.StatusInternalServerError, "DatabaseConnection Failed")
		return
	}

	id := context.Param("id")

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", id).Scan(&count)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Query Execution Failed"})
		return
	}

	if count == 0 {
		context.JSON(404, gin.H{"error": "UserId Not Found"})
		return
	}

	_, err = db.Exec("DELETE FROM USER WHERE UserId = ?", id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Query Execution Failed"})
	}
	context.JSON(http.StatusOK, "Deleted Successfully")
}

func UpdateUser(context *gin.Context) {
	db, err := SetUpDb()
	if err != nil {
		context.JSON(http.StatusInternalServerError, "DatabaseConnection Failed")
		return
	}

	var newUser model.User
	err = context.BindJSON(&newUser)
	if err != nil {
		return
	}
	id := context.Param("id")

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", id).Scan(&count)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Query Execution Failed"})
		return
	}

	if count == 0 {
		context.JSON(404, gin.H{"error": "UserId Not Found"})
		return
	}

	res, err := db.Query("UPDATE USER SET UserName = ? WHERE UserId = ?", newUser.Username, id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Query Execution Failed"})
		return
	}
	defer res.Close()
	context.JSON(http.StatusOK, "Updated Successfully")
}

func RunRestServer(){
	router := gin.Default()
	router.GET("/User/:id", GetUser)
	router.POST("/User", AddUser)
	router.PUT("/User/:id", UpdateUser)
	router.DELETE("/User/:id", DeleteUser)

	var cfn, _ = config.NewConfig()
	url := cfn.UserServiceUrl.RestUrl
	router.Run(url)
}

func main() {
	utils.SetEnv()
	RunRestServer()
}
