FROM golang:1.25-alpine AS builder 
RUN apk add --no-cache git

WORKDIR /app
COPY go.mod ./ 
RUN go mod download 

RUN go install github.com/air-verse/air@latest
ENV PATH="/root/go/bin:${PATH}"

FROM builder AS final
WORKDIR /app
COPY . .