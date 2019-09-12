FROM golang:1.13 as golang
COPY . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLE=0 go build -ldflags="-w -s"
CMD ./desafio-dito
EXPOSE 8080