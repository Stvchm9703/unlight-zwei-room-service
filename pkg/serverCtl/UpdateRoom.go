package serverCtl

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"

	// types "github.com/gogo/protobuf/types"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// UpdateRoom :
func (b *ULZRoomServiceBackend) UpdateRoom(ctx context.Context, req *pb.RoomCreateReq) (*pb.Room, error) {
	start := time.Now()
	b.mu.Lock()
	wkbox := b.searchAliveClient()
	cm.PrintReqLog(ctx, "update-room", req)
	defer func() {
		wkbox.Preserve(false)
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("update-Room took %s", elapsed)
	}()
	var room pb.Room
	_, err := wkbox.GetPara(&req.Key, &room)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	room.Password = req.Password
	room.CostLimitMax = req.CostLimitMax
	room.CostLimitMin = req.CostLimitMin
	room.CharCardNvn = req.CharCardNvn
	room.CharCardLimitMax = req.CharCardLimitMax
	room.CharCardLimitMin = req.CharCardLimitMin
	b.BroadCast(
		cm.MsgHostUpdateRoom(&room.Key, &room.Password))

	wkbox.UpdatePara(&room.Key, room)

	return &room, nil
}
