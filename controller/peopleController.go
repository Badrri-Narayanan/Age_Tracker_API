package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListOfPeople(ctx *gin.Context, db *sql.DB) {

	result, err := db.Query("SELECT * FROM get_list_of_people('');")

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Unable to connect to DB")
	}

	peoples := []model.People{}

	for result.Next() {
		var newPeople model.People
		result.Scan(&newPeople.Id, &newPeople.Name, &newPeople.DateOfBirth)
		newPeople.CalculateAge()

		peoples = append(peoples, newPeople)
	}

	ctx.JSON(http.StatusOK, peoples)
}

func AddPeople(ctx *gin.Context, db *sql.DB) {

	body := ctx.Request.Body

	values, err := ioutil.ReadAll(body)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid request. Please provide username and passowrd")
	}

	newUser := model.UserInfo{}

	json.Unmarshal([]byte(values), &newUser)

	result, err := db.Query("SELECT * FROM add_people($1, $2);", newUser.Name, newUser.DateOfBirth)

	if err != nil {
		fmt.Printf("Error : %v", err)
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}

	var id int

	for result.Next() {
		result.Scan(&id)
	}

	ctx.String(http.StatusOK, "User created successfully. Newly created User Id %v", id)
}
