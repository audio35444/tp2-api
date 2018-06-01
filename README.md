# TP 2: API Restful

## Arquitectura

Tiene la estructura => Controllers / Domain / Sservices
Tambien utiliza una libreria propia, hecha en golang para la utilizacion de Elasticsearch, se encuentra en la carpeta local_libs

## Como Correr la aplicación ?

```bash
git clone repo
cd repo
go run main.go
or
go build main.go
./main
```

API restful => http://localhost:8080
elasticsearch => http://localhost:9200

## Como utilizar la API

### Gestión del Token
- Pedido Token => curl -X POST http://localhost:8080/tokens  genera un nuevo token con 100 intentos
- Reset Token => curl -X PUT http://localhost:8080/tokens/:accessToken le devuelve los 100 intentos al token
- Query Token => curl -X GET http://localhost:8080/tokens/:accessToken devuelve los datos del token

### Consutla de Items

- Get Items => curl http://localhost:8080/items?token=(accessToken)

- Get Item => curl http://localhost:8080/items/:item?token=(accessToken)
- New item => curl -X POST http://localhost:8080/items?token=(accessToken) -d '{"id":"MLA'$(( ( RANDOM % 100000 )  + 100000 ))'Curl","title":"TitleCurl","category_id":"MLA33521curl","price":56.4,"currency_id":"MLAy332curl","available_quantity":332,"buying_mode":"only curl","descriptions":[{"id":"unIdCurl"}],"listing_type_id":"kasdfjhh33","condition":"lalalalallsjddshdshsd","video_id":"http:youtube.com/23423esdffsdFAS","pictures":[{"source":"https://algo.com.curl"}]}'

- Update item => curl -X PUT http://localhost:8080/items?token=(accessToken) -d '{"id":"'$1'","title":"TitleCurl_updated","category_id":"MLA33521curl","price":56.4,"currency_id":"MLAy332curl","available_quantity":332,"buying_mode":"only curl","descriptions":[{"id":"unIdCurl"}],"listing_type_id":"kasdfjhh33","condition":"lalalalallsjddshdshsd","video_id":"http:youtube.com/23423esdffsdFAS","pictures":[{"source":"https://algo.com.curl"}]}'

- Delete item => curl -X DELETE http://localhost:8080/items/:item/?token=(accessToken)


La dase de datos de Elasticsearch se puede cargar con una libreria propia realizada en golang, que consulta items reales de la API de Meli  [SDK-Meli](https://github.com/audio35444/sdk-meli)
