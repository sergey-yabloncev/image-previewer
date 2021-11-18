FROM golang:latest as build

ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY .. ${CODE_DIR}
RUN go mod download

RUN CGO_ENABLED=0 go build  -v -o /opt/app/app ./cmd/app
        
FROM alpine:3.9

WORKDIR /app

COPY --from=build /opt/app/ .

COPY .env .
COPY Makefile .

CMD ./app
