package serverCtl

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// JoinRoom :
func (b *ULZRoomServiceBackend) JoinRoom(ctx context.Context, req *pb.RoomReq) (*pb.Room, error) {
	start := time.Now()
	b.mu.Lock()
	wkbox := b.searchAliveClient()
	cm.PrintReqLog(ctx, "join-room", req)
	defer func() {
		(wkbox).Preserve(false)
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("join-Room took %s", elapsed)
	}()

	var tmp pb.Room
	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if tmp.Dueler == nil && req.IsDuel {
		tmp.Dueler = req.User
	}

	wkbox.UpdatePara(&tmp.Key, tmp)

	return &tmp, nil
}
