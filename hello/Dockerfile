FROM golang:alpine AS builder

WORKDIR /build
COPY . .


RUN go build -o hello hello.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/hello .

EXPOSE 8090

CMD ["./hello"]