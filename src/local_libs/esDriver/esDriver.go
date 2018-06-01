package esDriver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/audio35444/tp2-api/src/local_libs/esDriver/setting"
)

//el pretty que s emanda por get le indica a elasticsearch que devuelva en formato json
//curl -X GET "http://localhost:9200/_cat/indices?v"  consultar indices
//curl -X PUT "localhost:9200/client?pretty"
//curl -X GET "localhost:9200/client/_doc/1?pretty" consultar el doc 1 del index client
//curl -X DELETE "localhost:9200/client?pretty" eliminar el index client
//curl -X POST "localhost:9200/client/_doc/1/_update?pretty" -H 'Content-Type: application/json' -d'  update sobre el doc 1 index client
// {
//   "script" : "ctx._source.age += 5"
// }
// '

//curl -X POST "localhost:9200/client/_doc?pretty" -H 'Content-Type: application/json' -d agregar un doc nuevo id=autogenerado al index client
//{
//   "name": "John Doe"
//}
//curl -X GET "localhost:9200/{index_name}/_search?pretty"
// curl -X GET "localhost:9200/{index_name}/_search?Q=*&sort={attribute_name}:(asc|dsc)&pretty"

var obj setting.Setting = *setting.LoadSetting()

//curl -X PUT "localhost:9200/client/_doc/1?pretty" -H 'Content-Type: application/json' -d' agregar un doc nuevo id=1 al index client
// {
//   "name": "John Doe"
// }
// '
type Prueba struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type SP struct {
	Script string `json:"script"`
}
type Status struct {
	Index  string `json:"_index"`
	Id     string `json:"_id"`
	Result string `json:"result"`
	Found  bool   `json:"found"`
}

func GetDocs(indexName string) (*[]byte, error) {
	return genericRequest("GET",
		obj.GetEndpointDocSearch(indexName),
		nil,
		false)
}
func DeleteDoc(indexName string, docId string) (*[]byte, error) {
	return genericRequest("DELETE",
		obj.GetEndpointDocDelete(indexName, docId),
		nil,
		false)
}
func DeleteIndex() (*[]byte, error) {
	if len(os.Args) >= 2 {
		return genericRequest("DELETE",
			obj.GetEndpointNewIndex(os.Args[1]),
			nil,
			false)
	}
	return nil, errors.New("Not Args")
}
func UpdateDoc(element interface{}, indexName string, docId string) (*[]byte, error) {
	result, err := json.Marshal(element)
	if err != nil {
		return nil, err
	}
	flag, err := IsExistDoc(docId, indexName)
	if err != nil {
		return nil, err
	}
	if !flag {
		return nil, errors.New("Doc Not Exist")
	}
	return genericRequest("PUT",
		obj.GetEndpointDocIndexWithId(indexName, docId),
		strings.NewReader(string(result)),
		true)
}
func AddDocToIndex(element interface{}, indexName string) (*[]byte, error) {
	result, _ := json.Marshal(element)
	return genericRequest("POST",
		obj.GetEndpointDocIndex(indexName),
		strings.NewReader(string(result)),
		true)
}

// {
//   "query": {
//     "bool": {
//       "must": [
//         { "match_phrase": { "id": "MLA66630074" } }
//       ]
//     }
//   }
// }
func IsExistDoc(docId string, indexName string) (bool, error) {
	data, err := GetDoc(indexName, docId)
	if err != nil {
		return false, err
	}
	var objStatus Status
	json.Unmarshal(*data, &objStatus)
	return objStatus.Found, nil
}
func AddDoc(element interface{}, indexName string, docId string) (*[]byte, error) {
	result, err := json.Marshal(element)
	if err != nil {
		return nil, err
	}
	flag, err := IsExistDoc(docId, indexName)
	if err != nil {
		return nil, err
	}
	if flag {
		return nil, errors.New("Doc Exist")
	}
	return genericRequest("PUT",
		obj.GetEndpointDocIndexWithId(indexName, docId),
		strings.NewReader(string(result)),
		true)
}
func AddDocToIndexWithId(element interface{}, indexName string, docId string) (*[]byte, error) {
	result, _ := json.Marshal(element)
	return genericRequest("PUT",
		obj.GetEndpointDocIndexWithId(indexName, docId),
		strings.NewReader(string(result)),
		true)
}
func GetDoc(indexName string, docId string) (*[]byte, error) {
	return genericRequest("GET",
		obj.GetEndpointDocFromIndex(indexName, docId),
		nil,
		false)
}
func NewIndex(indexName string) {
	data, err := genericRequest("PUT",
		obj.GetEndpointNewIndex(indexName),
		nil,
		false)
	if err == nil {
		fmt.Println(string(*data))
	} else {
		fmt.Println(error(err).Error())
	}
}
func GetIndices() {
	data, err := genericRequest("GET",
		obj.GetEndpointIndices(),
		nil,
		false)
	if err == nil {
		fmt.Println(string(*data))
	} else {
		fmt.Println(error(err).Error())
	}
}
func genericRequest(method string, fullPath string, body io.Reader, isJson bool) (dataResult *[]byte, errResult error) {
	err := isOn()
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, fullPath, body)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	if isJson {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := client.Do(req)
	if err == nil {
		defer res.Body.Close()
		data, err1 := ioutil.ReadAll(res.Body)
		dataResult = &data
		errResult = err1
		return
	} else {
		errResult = err
		return
	}
}

//isOn verifica si hay conexion al servidor de elasticsearch
func isOn() error {
	_, err := http.Get(obj.GetEndpointPing())
	if err != nil {
		return errors.New("elastic-off")
	}
	return nil
}
func PruebaUpdate() {
	var script SP
	script.Script = "ctx._source.age += 5"
	result, _ := json.Marshal(script)
	data, err := genericRequest("POST",
		obj.GetEndpointDocUpdate("nuevoindice", "BBxKp2MBedOyFXZO2eFz"),
		strings.NewReader(string(result)),
		true)
	if err == nil {
		fmt.Println(string(*data))
	} else {
		fmt.Println(error(err).Error())
	}
}

// func genericShow(res *http.Response, err error) {
// 	if err == nil {
// 		defer res.Body.Close()
// 		data, _ := ioutil.ReadAll(res.Body)
// 		fmt.Println("\n---------------- Elasticsearch Indices ----------------")
// 		fmt.Println(string(data))
// 		fmt.Println("--------------------------------")
// 	} else {
// 		fmt.Println(err.Error())
// 	}
// }
