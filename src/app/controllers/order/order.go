package order

import (
  "net/http"
  "strconv"
  "fmt"
  "github.com/gin-gonic/gin"
  orderService "github.com/audio35444/accelerator-api/src/app/services/order"
)
//http://localhost:8000/order?orderId=21432   => c.Query("orderId")
//http://localhost:8000/order/23432324 => c.Param("orderId")
func GetOrder(c *gin.Context){
  // c.Query("orderId")
  // num,_ :=strconv.ParseInt("-42", 10, 64)
  num,err:= strconv.ParseInt(c.Param("orderId"),10,64)
  if err != nil{
    //lo que quiero devolver en el caso de error
    //400 mal enviada la info
    //404 por que no la encontre
    c.String(http.StatusBadRequest,err.Error())
    return
  }
  fmt.Println(num)
  // str:= c.Param("orderId")
  orderInfo,err := orderService.GetOrder(num)
  //el nil solo es adminitdo por estructuras, no por datos primitivos como string int estructuras
  //erros.New("B no puede ser 0")
  // a,b=b,a   => intercambio de a por b y b por a
  c.JSON(http.StatusOK,orderInfo)//espera una estructura o una referencia a una estructura
}
