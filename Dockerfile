FROM golang:1.16.5 as development

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN go build -o spacelight ./internal

FROM alpine:3.14
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/spacelight ./
CMD ["./spacelight"]  