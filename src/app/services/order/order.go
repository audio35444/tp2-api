package order

import (
  orderDomain "github.com/audio35444/accelerator-api/src/app/domain/order"
)
func GetOrder(orderId int64)(*orderDomain.Order,error){
  return &orderDomain.Order{
    Id:orderId,
    Name:"emma",
  },nil
  //de esa forma evito usar el & al retornar
  // return obj,nil   funcionQueDevuyelvedosValores()   orderInfo,err
}

//func getOrderData() orderData
//func processOrderData() orderInformation
