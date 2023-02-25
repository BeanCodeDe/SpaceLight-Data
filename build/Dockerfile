FROM golang:1.20.1 as development

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN make app.build

FROM alpine:3.14
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/spacelight-data ./
CMD ["./spacelight-data"]  