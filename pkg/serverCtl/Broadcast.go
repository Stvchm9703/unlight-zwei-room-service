package serverCtl

import (
	pb "ULZRoomService/proto"
	"context"

	"log"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// "time"
)

func (this *ULZRoomServiceBackend) ServerBroadcast(rReq *pb.RoomReq, stream pb.RoomService_ServerBroadcastServer) error {
	_, err := this.AddStream(&rReq.Key, &rReq.UserId, &stream)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	go func() {
		<-stream.Context().Done()
		log.Println("close done")
		_, err := this.DelStream(&rReq.Key, &rReq.UserId)
		if err != nil {
			log.Println(err)
		}
		this.BroadCast(&rReq.Key, &rReq.UserId,
			&pb.RoomMsg{
				Key:     rReq.Key,
				FormId:  rReq.UserId,
				ToId:    "ALL_USER",
				MsgType: pb.RoomMsg_SYSTEM_INFO,
				Message: rReq.UserName + " is offline",
			})
	}()
	for {
	}
}

func (this *ULZRoomServiceBackend) SendMessage(ctx context.Context, msg *pb.RoomMsg) (*pb.Empty, error) {
	this.BroadCast(&msg.Key, &msg.FormId, msg)
	return nil, nil
}
