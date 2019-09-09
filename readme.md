Instalação
docker-compose -f "docker-compose.yml" up -d --build

ENDPOINTS

AUTOCOMPLETE API
curl -d '{"event": "com"}' -X POST http://localhost:8080/autocomplete_api

AUTOCOMPLETE hmtl
http://localhost:8080/get

GET(Obtem um json com eventos de compra e salva no banco)
curl -d '{"event": "comprou-produto"}' -X POST http://localhost:8080/get

timeline
curl -X POST http://localhost:8080/timeline