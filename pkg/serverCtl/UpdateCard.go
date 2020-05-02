package serverCtl

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (b *ULZRoomServiceBackend) UpdateCard(ctx context.Context, req *pb.RoomUpdateCardReq) (*pb.Empty, error) {

	start := time.Now()
	b.mu.Lock()
	cm.PrintReqLog(ctx, "update-card", req)
	wkbox := b.searchAliveClient()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Update-Card took %s", elapsed)
		(wkbox).Preserve(false)
	}()

	var tmp pb.Room
	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if req.Side == pb.RoomUpdateCardReq_HOST {
		tmp.HostCharcardId = req.CharcardId
		tmp.HostCardsetId = req.CardsetId
		tmp.HostCardlevel = req.Level
	} else {
		tmp.HostCharcardId = req.CharcardId
		tmp.HostCardsetId = req.CardsetId
		tmp.HostCardlevel = req.Level
	}

	go func() {
		// byts, _ := proto.Marshal(req)
		b.BroadCast(&pb.RoomMsg{
			Key:     req.Key,
			FromId:  req.Side.String(),
			FmName:  req.Side.String(),
			ToId:    "All",
			ToName:  "All",
			MsgType: pb.RoomMsg_SYSTEM_INFO,
			Message: fmt.Sprintf("CardChange::%s", proto.MarshalTextString(req)),
		})
	}()
	wkbox.SetPara(&req.Key, &tmp)
	return &pb.Empty{}, nil
}
