FROM golang:alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download \
&& go build -o order server/main.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/order .

EXPOSE 8090

CMD ["./order"]