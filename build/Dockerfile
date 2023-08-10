FROM golang:1.20-alpine AS builder

ARG NAME=producer

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o ${NAME} ./cmd/producer/producer.go


FROM alpine:3 AS final

WORKDIR /app
COPY --from=builder /app/${NAME} .

VOLUME [ "./.env" ]

CMD ["./producer"]