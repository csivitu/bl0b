FROM golang:1.15.5-alpine3.12 AS prepare
RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

RUN wget -q https://github.com/roerohan/wait-for-it/releases/download/v0.2.2/wait-for-it
RUN chmod +x ./wait-for-it

FROM prepare AS build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/bl0b

FROM scratch

COPY --from=build /app/bin/bl0b /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/wait-for-it .

ENTRYPOINT [ "/app/bl0b" ]
