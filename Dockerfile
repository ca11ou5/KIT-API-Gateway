FROM golang:1.21-alpine AS builder
WORKDIR /usr/local/src

RUN apk --no-cache add git make musl-dev gcc

#dependecies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#build
COPY ./ ./
RUN go build -o ./bin/main ./cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/main /
COPY configs/envs/dev.env /dev.env

CMD ["/main"]