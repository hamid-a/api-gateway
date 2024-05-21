FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway cmd/api-gateway/api-gateway.go

FROM alpine:edge

WORKDIR /app

COPY --from=build /app/api-gateway .

RUN apk --no-cache add ca-certificates tzdata

ENTRYPOINT ["/app/api-gateway"]
