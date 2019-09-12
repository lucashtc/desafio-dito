# Desafio DITO
## Instalação da API e do banco <br>

```bash
docker-compose -f "docker-compose.yml" up -d --build 
```
Banco é criado e alimentado durante criação do container
## ENDPOINTS

### AUTOCOMPLETE API
```bash
curl -d '{"event": "com"}' -X POST http://localhost:8080/autocomplete_api
```

### AUTOCOMPLETE Pagina
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