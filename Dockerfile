FROM golang:1.22.1-alpine3.19
WORKDIR /app
COPY go.mod .
COPY main.go .
RUN go mod download
RUN go build -v -o ./build/myapp
CMD ["./build/myapp"]
