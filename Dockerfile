FROM golang:1.22-alpine

WORKDIR /app1

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

CMD [ "./main" ]