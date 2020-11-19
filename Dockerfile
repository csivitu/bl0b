FROM golang:1.15.5-alpine3.12 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/bl0b

FROM scratch

COPY --from=build /app/bin/bl0b /app/

ENTRYPOINT [ "/app/bl0b" ]
