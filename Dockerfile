FROM golang:1.13.1 AS build-env
WORKDIR /ULZRoomService
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY go.mod /ULZRoomService/go.mod
COPY go.sum /ULZRoomService/go.sum
RUN go mod download
COPY . /ULZRoomService
RUN CGO_ENABLED=0 GOOS=linux go build -o build/ULZRoomService ./ULZRoomService


FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /ULZRoomService/build/ULZRoomService /
ENTRYPOINT ["/ULZRoomService"]
CMD ["up", "--grpc-port=80"]
