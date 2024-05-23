FROM golang:1.21.6 as builder

WORKDIR /jump_server

COPY . .

RUN go mod download

RUN go install ./cmd/jump_server

FROM debian:bookworm-slim

WORKDIR /jump_server

COPY --from=builder /go/bin/jump_server .

CMD ["/jump_server/jump_server"]
