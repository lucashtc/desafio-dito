FROM golang:1.13
COPY . /go/src/lucas
WORKDIR /go/src/lucas
RUN go build -o dito
CMD ./dito
EXPOSE 8080