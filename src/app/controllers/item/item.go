package item

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	itemDomain "github.com/audio35444/tp2-api/src/app/domain/item"
	searchDomain "github.com/audio35444/tp2-api/src/app/domain/search"
	statusDomain "github.com/audio35444/tp2-api/src/app/domain/status"
	itemService "github.com/audio35444/tp2-api/src/app/services/item"
	tokenService "github.com/audio35444/tp2-api/src/app/services/token"
	"github.com/gin-gonic/gin"
)

//http://localhost:8000/order?orderId=21432   => c.Query("orderId")
//http://localhost:8000/order/23432324 => c.Param("orderId")
type Error struct {
	Name       string
	StatusCode int
}

func UpdateItem(c *gin.Context) {
	//startTokenValidation
	strToken := c.Query("token")
	if len(strToken) <= 0 {
		c.JSON(http.StatusBadRequest, Error{Name: "Invalid Token", StatusCode: http.StatusBadRequest})
		return
	}
	err := tokenService.DecrementToken(strToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	//endTokenValidation
	defer c.Request.Body.Close()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error()})
		return
	}
	var objItem itemDomain.Item
	json.Unmarshal(body, &objItem)
	if len(objItem.Id) <= 0 {
		c.JSON(http.StatusBadRequest, Error{Name: "empty Id"})
		return
	}
	// data, err := json.Marshal(string(body))
	data, err := itemService.UpdateItem(objItem, objItem.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error()})
		return
	}
	var objStatus statusDomain.Status
	json.Unmarshal(*data, &objStatus)
	if objStatus.Result != "updated" {
		c.JSON(http.StatusBadRequest, objStatus)
		return
	}
	c.JSON(http.StatusCreated, objStatus)
}
func DeleteItem(c *gin.Context) {
	//startTokenValidation
	strToken := c.Query("token")
	if len(strToken) <= 0 {
		c.JSON(http.StatusBadRequest, Error{Name: "Invalid Token", StatusCode: http.StatusBadRequest})
		return
	}
	err := tokenService.DecrementToken(strToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	//endTokenValidation
	docId := c.Param("docId")
	data, err := itemService.DeleteItem(docId)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: "Bad Request", StatusCode: http.StatusBadRequest})
		return
	}
	var objStatus statusDomain.Status
	json.Unmarshal(*data, &objStatus)
	if objStatus.Result != "deleted" {
		c.JSON(http.StatusBadRequest, objStatus)
		return
	}
	c.JSON(http.StatusOK, objStatus)
}

func AddItem(c *gin.Context) {
	//startTokenValidation
	strToken := c.Query("token")
	if len(strToken) <= 0 {
		c.JSON(http.StatusBadRequest, Error{Name: "Invalid Token", StatusCode: http.StatusBadRequest})
		return
	}
	err := tokenService.DecrementToken(strToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	//endTokenValidation
	defer c.Request.Body.Close()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error()})
		return
	}
	var objItem itemDomain.Item
	json.Unmarshal(body, &objItem)
	if len(objItem.Id) <= 0 {
		c.JSON(http.StatusBadRequest, Error{Name: "empty Id"})
		return
	}
	// data, err := json.Marshal(string(body))
	data, err := itemService.AddItem(objItem, objItem.Id)
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

func GetItem(c *gin.Context) {
	//startTokenValidation
	strToken := c.Query("token")
	if len(strToken) <= 0 {
		c.JSON(http.StatusBadRequest, Error{Name: "Invalid Token", StatusCode: http.StatusBadRequest})
		return
	}
	err := tokenService.DecrementToken(strToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	//endTokenValidation
	docId := c.Param("docId")
	data, err := itemService.GetItem(docId)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: "Bad Request", StatusCode: http.StatusBadRequest})
		return
	}
	type Source struct {
		Source itemDomain.Item `json:"_source"`
	}
	var objItem Source
	json.Unmarshal(*data, &objItem)
	if len(objItem.Source.Id) <= 0 {
		var objStatus statusDomain.Status
		json.Unmarshal(*data, &objStatus)
		c.JSON(http.StatusNotFound, Error{Name: "Not Found", StatusCode: http.StatusNotFound})
		return
	}
	c.JSON(http.StatusOK, objItem.Source)
}
func GetItems(c *gin.Context) {
	//startTokenValidation
	strToken := c.Query("token")
	if len(strToken) <= 0 {
		c.JSON(http.StatusBadRequest, Error{Name: "Invalid Token", StatusCode: http.StatusBadRequest})
		return
	}
	err := tokenService.DecrementToken(strToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	//endTokenValidation
	data, err := itemService.GetItems()
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Name: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	var objSearch searchDomain.Search
	json.Unmarshal(*data, &objSearch)
	c.JSON(http.StatusOK, objSearch)
}

// func GetItem(c *gin.Context){
//   // c.Query("orderId")
//   // num,_ :=strconv.ParseInt("-42", 10, 64)
//   num,err:= strconv.ParseInt(c.Param("orderId"),10,64)
//   if err != nil{
//     //lo que quiero devolver en el caso de error
//     //400 mal enviada la info
//     //404 por que no la encontre
//     c.String(http.StatusBadRequest,err.Error())
//     return
//   }
//   fmt.Println(num)
//   // str:= c.Param("orderId")
//   orderInfo,err := orderService.GetOrder(num)
//   //el nil solo es adminitdo por estructuras, no por datos primitivos como string int estructuras
//   //erros.New("B no puede ser 0")
//   // a,b=b,a   => intercambio de a por b y b por a
//   c.JSON(http.StatusOK,orderInfo)//espera una estructura o una referencia a una estructura
// }
