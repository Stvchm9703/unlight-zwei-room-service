FROM golang:1.13.7 AS build-env
WORKDIR /ULZRoomService
COPY go.mod /ULZRoomService/go.mod
COPY go.sum /ULZRoomService/go.sum
RUN go mod download
COPY . /ULZRoomService
RUN go build -o build_cli/room_status.go ./ULZRoomService

ENTRYPOINT ["/ULZRoomService"]

# CMD ["up", "--grpc-port=80"]
EXPOSE 11000

CMD [ "/ULZRoomService", "run", "-c", "config.yaml" ]