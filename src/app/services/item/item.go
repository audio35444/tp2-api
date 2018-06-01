package item

import (
	elasticsearch "github.com/audio35444/tp2-api/src/local_libs/esDriver"
)

func GetItems() (*[]byte, error) {
	return elasticsearch.GetDocs("items")
}
func GetItem(docId string) (*[]byte, error) {
	return elasticsearch.GetDoc("items", docId)
}
func AddItem(element interface{}, docId string) (*[]byte, error) {
	return elasticsearch.AddDoc(element, "items", docId)
}
func DeleteItem(docId string) (*[]byte, error) {
	return elasticsearch.DeleteDoc("items", docId)
}
func UpdateItem(element interface{}, docId string) (*[]byte, error) {
	return elasticsearch.UpdateDoc(element, "items", docId)
}

// func GetOrder(orderId int64)(*orderDomain.Order,error){
//   return &orderDomain.Order{
//     Id:orderId,
//     Name:"emma",
//   },nil
//   //de esa forma evito usar el & al retornar
//   // return obj,nil   funcionQueDevuyelvedosValores()   orderInfo,err
// }

//func getOrderData() orderData
//func processOrderData() orderInformation
