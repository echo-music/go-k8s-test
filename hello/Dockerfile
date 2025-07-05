FROM golang:alpine AS builder

WORKDIR /build
COPY ../hello .
#COPY ../go.mod .


RUN go build -o hello hello/hello.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/hello .

EXPOSE 8090

CMD ["./hello"]