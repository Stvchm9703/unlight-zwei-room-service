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

// GetRoomInfo :
func (b *ULZRoomServiceBackend) GetRoomInfo(ctx context.Context, req *pb.RoomReq) (*pb.Room, error) {
	cm.PrintReqLog(ctx, "get-room-info", req)

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

	if tmp.Password != "" && tmp.Password != req.Password {
		return nil, status.Error(codes.PermissionDenied, "ROOM_PASSWORD_INV")
	}

	return &tmp, nil
}
