package token

import (
	"encoding/json"
	"fmt"
	"net/http"

	statusDomain "github.com/audio35444/tp2-api/src/app/domain/status"
	tokenDomain "github.com/audio35444/tp2-api/src/app/domain/token"
	tokenService "github.com/audio35444/tp2-api/src/app/services/token"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Name       string
	StatusCode int
}

// func DecrementToken(objToken string) bool {
//
// }
func ResetToken(c *gin.Context) {
	strToken := c.Param("token")
	fmt.Println(strToken)
	if len(strToken) <= 0 {
		c.JSON(http.StatusBadRequest, Error{Name: "Invalid Token", StatusCode: http.StatusBadRequest})
		return
	}
	data, err := tokenService.ResetToken(strToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	var objStatus statusDomain.Status
	json.Unmarshal(*data, &objStatus)
	c.JSON(http.StatusCreated, objStatus)
}
func AddToken(c *gin.Context) {
	data, err := tokenService.AddToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error()})
		return
	}
	var objStatus statusDomain.Status
	json.Unmarshal(*data, &objStatus)
	if objStatus.Result != "created" {
		c.JSON(http.StatusBadRequest, objStatus)
		return
	}
	c.JSON(http.StatusCreated, objStatus)
}
func GetToken(c *gin.Context) {
	docId := c.Param("docId")
	data, err := tokenService.GetToken(docId)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: "Bad Request", StatusCode: http.StatusBadRequest})
		return
	}
	type Source struct {
		Source tokenDomain.Token `json:"_source"`
	}
	var objToken Source
	json.Unmarshal(*data, &objToken)
	if len(objToken.Source.Id) <= 0 {
		var objStatus statusDomain.Status
		json.Unmarshal(*data, &objStatus)
		c.JSON(http.StatusNotFound, Error{Name: "Not Found", StatusCode: http.StatusNotFound})
		return
	}
	c.JSON(http.StatusOK, objToken.Source)
}
