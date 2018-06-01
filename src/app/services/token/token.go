package token

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	tokenDomain "github.com/audio35444/tp2-api/src/app/domain/token"
	elasticsearch "github.com/audio35444/tp2-api/src/local_libs/esDriver"
)

func GetTokens() (*[]byte, error) {
	return elasticsearch.GetDocs("tokens")
}
func GetToken(docId string) (*[]byte, error) {
	return elasticsearch.GetDoc("tokens", docId)
}
func ResetToken(docId string) (*[]byte, error) {
	type Source struct {
		Source tokenDomain.Token `json:"_source"`
	}
	var objToken Source
	data, err := elasticsearch.GetDoc("tokens", docId)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(*data, &objToken)
	objToken.Source.Cant = 100
	return UpdateToken(objToken.Source, objToken.Source.Id)
}
func DecrementToken(docId string) error {
	type Source struct {
		Source tokenDomain.Token `json:"_source"`
	}
	var objToken Source
	data, err := elasticsearch.GetDoc("tokens", docId)
	if err != nil {
		return err
	}
	json.Unmarshal(*data, &objToken)
	if objToken.Source.Cant <= 0 {
		return errors.New("Invalid Token")
	}
	objToken.Source.Cant = objToken.Source.Cant - 1
	UpdateToken(objToken.Source, objToken.Source.Id)
	return nil
}
func AddToken() (*[]byte, error) {
	t := time.Now()
	hex := getMD5HashHex(string(rand.Intn(1000)) + t.String() + string(rand.Intn(1000)))
	// b64 := getMD5HashBase64(string(rand.Intn(1000)) + t.String() + string(rand.Intn(1000)))
	objToken := tokenDomain.Token{
		Id:          hex,
		Cant:        100,
		CreateToken: t,
	}
	return elasticsearch.AddDoc(objToken, "tokens", hex)
}
func DeleteToken(docId string) (*[]byte, error) {
	return elasticsearch.DeleteDoc("tokens", docId)
}
func UpdateToken(element interface{}, docId string) (*[]byte, error) {
	return elasticsearch.UpdateDoc(element, "tokens", docId)
}

func getMD5HashHex(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func getMD5HashBase64(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
