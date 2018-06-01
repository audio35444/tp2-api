package setting

import (
	"regexp"
)

var reIndexName = regexp.MustCompile(`{index_name}`)
var reId = regexp.MustCompile(`/{id}`)

type Setting struct {
	RootEndpoint string `json:"root-endpoint"`
	Port         string `json:"port"`
	IndicesShow  string `json:"indices-show"`
	IndexName    string `json:"index-name"`
	DocIndex     string `json:"doc-index"`
	DocUpdate    string `json:"doc-update"`
	DocSearch    string `json:"doc-search"`
}

func (p *Setting) GetEndpointDocSearch(indexName string) string {
	return p.RootEndpoint + p.Port + reIndexName.ReplaceAllString(p.DocSearch, indexName)
}
func (p *Setting) GetEndpointDocUpdate(indexName string, docId string) string {
	return p.RootEndpoint + p.Port + reId.ReplaceAllString(reIndexName.ReplaceAllString(p.DocUpdate, indexName), "/"+docId)
}
func (p *Setting) GetEndpointDocDelete(indexName string, docId string) string {
	return p.RootEndpoint + p.Port + reId.ReplaceAllString(reIndexName.ReplaceAllString(p.DocIndex, indexName), "/"+docId)
}
func (p *Setting) GetEndpointDocIndex(indexName string) string {
	return p.RootEndpoint + p.Port + reId.ReplaceAllString(reIndexName.ReplaceAllString(p.DocIndex, indexName), "")
}
func (p *Setting) GetEndpointDocFromIndex(indexName string, docId string) string {
	//BBxKp2MBedOyFXZO2eFz
	return p.RootEndpoint + p.Port + reId.ReplaceAllString(reIndexName.ReplaceAllString(p.DocIndex, indexName), "/"+docId)
}
func (p *Setting) GetEndpointDocIndexWithId(indexName string, docId string) string {
	return p.RootEndpoint + p.Port + reId.ReplaceAllString(reIndexName.ReplaceAllString(p.DocIndex, indexName), "/"+docId)
}
func (p *Setting) GetEndpointNewIndex(indexName string) string {
	return p.RootEndpoint + p.Port + reIndexName.ReplaceAllString(p.IndexName, indexName)
}

//devuelve la URL de sitios
func (p *Setting) GetEndpointIndices() string {
	return p.RootEndpoint + p.Port + p.IndicesShow
}

func (p *Setting) GetEndpointPing() string {
	return p.RootEndpoint + p.Port
}

func LoadSetting() *Setting {
	// p.RootEndpoint = "http://localhost"
	// p.DocUpdate = "{index_name}/_doc/{id}/_update?pretty"
	// p.DocSearch = "{index_name}/_search?pretty"
	// p.IndexName = "{index_name}?pretty"
	// p.DocIndex = "{index_name}/_doc/{id}?pretty"
	// p.IndicesShow = "_cat/indices?v"
	// p.Port = ":9200/"
	return &Setting{
		RootEndpoint: "http://localhost",
		DocUpdate:    "{index_name}/_doc/{id}/_update?pretty",
		DocSearch:    "{index_name}/_search?pretty",
		IndexName:    "{index_name}?pretty",
		DocIndex:     "{index_name}/_doc/{id}?pretty",
		IndicesShow:  "_cat/indices?v",
		Port:         ":9200/",
	}
	// p = &Setting{
	// 	RootEndpoint: "http://localhost",
	// 	,
	// 	,
	// 	,
	// 	,
	// 	,
	// 	DocSearch:    ,
	// }
	// ex, err := os.Executable()
	// if err != nil {
	// 	fmt.Println("entro en ele erro")
	// 	panic(err)
	// }
	// exPath := filepath.Dir(ex)
	// data, err := ioutil.ReadFile(exPath + "./esDriver/.driver-elasticsearch-setting.json")
	// if err == nil {
	// 	json.Unmarshal(data, &p)
	// } else {
	// 	data, err := ioutil.ReadFile("./esDriver/.driver-elasticsearch-setting.json")
	// 	// fmt.Println(err.Error())
	// 	if err == nil {
	// 		json.Unmarshal(data, &p)
	// 	}
	// }
}
