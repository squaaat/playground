FROM golang:1.16.5 as BUILDER
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV CGO_ENABLED=0
RUN go build  -o ./playground ./main.go


FROM alpine:3.12

WORKDIR /

COPY --from=BUILDER /build/playgroud /playground

ENTRYPOINT ["echo", "Must specify a binary for entrypoint, /search-server"]