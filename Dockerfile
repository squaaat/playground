FROM golang:1.16.5 as BUILDER
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV CGO_ENABLED=0
RUN go build -o ./playground-server ./main.go

# FROM alpine:3.12

# WORKDIR /

# COPY --from=BUILDER /build/playgroud-server /playground-server

# ENTRYPOINT ["echo", "Must specify a binary for entrypoint, /playground-server"]