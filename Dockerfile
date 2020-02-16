FROM golang:1.13.7 AS build-env
WORKDIR /ULZRoomService
COPY go.mod /ULZRoomService/go.mod
COPY go.sum /ULZRoomService/go.sum
RUN go mod download
COPY . /ULZRoomService
RUN go build -o ./room_service build_cli/room_status.go 

ENTRYPOINT ["/ULZRoomService"]

EXPOSE 11000

CMD [ "/room_service", "run", "-c", "config.yaml" ]