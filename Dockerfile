FROM golang:1.13 AS builder
WORKDIR /app
COPY ./room_status ./

COPY ./insecure ./insecure

COPY ./config.test_server.yaml ./config.yaml

EXPOSE 11000

CMD [ "./roomserver", "start", "-c", "config.yaml" ]



