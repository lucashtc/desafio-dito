# Desafio DITO
## Instalação da API e do banco <br>


```bash
docker-compose -f "docker-compose.yml" up -d --build 
```
## Criar base e alimentar com dados fake
Será feito insert de 30.000 linhas na base
```bash
curl  -X POST http://192.168.99.100:8080/migration
```

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