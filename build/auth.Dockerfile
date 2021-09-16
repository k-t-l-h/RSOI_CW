FROM golang:latest as builder
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY .env .
COPY . .
RUN  go test internal/...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth cmd/auth/main.go

EXPOSE 8010
CMD sleep 15 && ./auth
