FROM golang:alpine AS builder

WORKDIR /build
COPY server/main.go .

RUN go build -o order main.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/order .

EXPOSE 8090

CMD ["./hello"]