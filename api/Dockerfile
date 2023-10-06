FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/cosmtrek/air

COPY . .

RUN go build -o main .

EXPOSE 8888

CMD ["air"]