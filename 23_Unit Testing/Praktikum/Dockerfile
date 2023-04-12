from golang:1.18-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /docker-ping

EXPOSE 8080

CMD [ "/docker-ping" ]