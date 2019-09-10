# Desafio DITO
##Instalação do banco <br>
Para banco de dados use mysql <br>
O arquivo dito.sql possui script para criação e aliemtnação do banco de dados.<br>
Configure o arquivo autocomplete/database.go informando o endereço ip(ou hostname) do banco,nome e usuário. <br>
Use o seguinte comando para copilar o codigo GO e executar.

```bash
docker-compose -f "docker-compose.yml" up -d --build 
```
*apenas para executar codigo web server em GO/Golang <br>
*banco por enquanto precisa ser instalado manualmente

## ENDPOINTS

### AUTOCOMPLETE API
```bash
curl -d '{"event": "com"}' -X POST http://localhost:8080/autocomplete_api
```

### AUTOCOMPLETE hmtl
```bash
http://localhost:8080/get
```

### GET(Obtem um json com eventos de compra e salva no banco)
```bash
curl -d '{"event": "comprou-produto"}' -X POST http://localhost:8080/get
```

### TIMELINE
```bash
curl -X POST http://localhost:8080/timeline
```