FROM golang:alpine

WORKDIR /app/url-shortener-service

COPY go.mod .
COPY go.sum .
ENV GOPATH=/
RUN go mod download

#build appliction
COPY . .
RUN go build -o url-shortener-service ./cmd/main/app.go

CMD ["./url-shortener-service"]