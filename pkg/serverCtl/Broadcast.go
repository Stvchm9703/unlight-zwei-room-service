package serverCtl

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"context"

	"log"

	"github.com/gogo/status"

	"google.golang.org/grpc/codes"
	// "time"
)

func (this *ULZRoomServiceBackend) ServerBroadcast(rReq *pb.RoomReq, stream pb.RoomService_ServerBroadcastServer) error {
	_, err := this.AddStream(&rReq.Key, &rReq.User.Id, &stream)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	go func() {
		<-stream.Context().Done()
		log.Println("close done")
		_, err := this.DelStream(&rReq.Key, &rReq.User.Id)
		if err != nil {
			log.Println(err)
		}
		this.BroadCast(&rReq.Key, &rReq.User.Id,
			cm.MsgUserQuitRoom(&rReq.Key, &rReq.User.Id, &rReq.User.Name))
	}()
	for {
	}
}

func (this *ULZRoomServiceBackend) SendMessage(ctx context.Context, msg *pb.RoomMsg) (*pb.Empty, error) {
	this.BroadCast(&msg.Key, &msg.FormId, msg)
	return nil, nil
}
