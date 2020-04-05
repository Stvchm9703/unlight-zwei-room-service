package serverCtl

import (
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
	wkbox := b.searchAliveClient()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Get-Room took %s", elapsed)
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
		byts, _ := proto.Marshal(req)
		b.BroadCast(&pb.RoomMsg{
			Key:     req.Key,
			FromId:  req.Side.String(),
			FmName:  req.Side.String(),
			ToId:    "All",
			ToName:  "All",
			Message: fmt.Sprintf("CardChange::%s", string(byts)),
		})
	}()
	wkbox.SetPara(&req.Key, &tmp)
	// if tmp.Password != "" && tmp.Password != req.Password {
	// 	return nil, status.Error(codes.PermissionDenied, "ROOM_PASSWORD_INV")
	// }

	return nil, status.Errorf(codes.Unimplemented, "method UpdateCard not implemented")
}
