FROM golang:1.17 as build

ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

RUN CGO_ENABLED=0 go build -o /opt/app/app ./cmd/app
        
FROM alpine:3.9

WORKDIR /app

COPY --from=build /opt/app/ .

COPY ./configs/config.toml .

CMD ./app -config ./config.toml