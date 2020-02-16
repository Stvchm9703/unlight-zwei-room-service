FROM golang:1.13.7 AS build-env
WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download
COPY . /app
RUN go build -o ./room_service build_cli/room_status.go 

EXPOSE 11000

CMD [ "/room_service", "run", "-c", "config.yaml" ]