curl -X POST http://localhost:8080/db/somekey -H "Content-Type:application/octet-stream" -d "value"

curl -X GET http://localhost:8080/db/somekey

curl -X POST http://localhost:8080/db/jsonkey -H "Content-Type:application/octet-stream" -d $'{"name":"AnuchitO"}'

curl -X GET "http://localhost:8080/db/jsonkey?format=json"