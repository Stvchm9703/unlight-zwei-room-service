package serverCtl

import (
	pb "ULZRoomService/proto"
	"context"
	// "log"
	// "time"
	// "github.com/gogo/status"
)

func (this *ULZRoomServiceBackend) ServerBroadcast(rReq *pb.RoomReq, stream pb.RoomService_ServerBroadcastServer) error {
	// this.bc_stream
	// if _,err:=this.GetGS(rReq)
	return nil
}

func (this *ULZRoomServiceBackend) SendMessage(context.Context, *pb.RoomMsg) (*pb.Empty, error) {
	return nil, nil
}
