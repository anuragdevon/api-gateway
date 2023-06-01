FROM golang:alpine

WORKDIR /api_gateway

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./api_gateway ./main.go

EXPOSE 8000

CMD ["./api_gateway"]
