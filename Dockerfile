FROM golang:1.23.4
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8000
CMD ["./main"]