
::Instalação
Para banco de dados use mysql usando o arquivo dito.sql para instalação do banco
Configure o arquivo autocomplete/database.go informando o endereço ip do banco,nome e usuário. 
e execute o segunte comando para executar a API

docker-compose -f "docker-compose.yml" up -d --build
*apenas para executar codigo web server em GO/Golang

::ENDPOINTS

::AUTOCOMPLETE API
curl -d '{"event": "com"}' -X POST http://localhost:8080/autocomplete_api

::AUTOCOMPLETE hmtl
http://localhost:8080/get

::GET(Obtem um json com eventos de compra e salva no banco)
curl -d '{"event": "comprou-produto"}' -X POST http://localhost:8080/get

::TIMELINE
curl -X POST http://localhost:8080/timeline