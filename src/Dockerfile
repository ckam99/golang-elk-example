FROM golang:1.22


WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o server main.go

CMD ["/app/server"]