package serverCtl

import (
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// GetRoomInfo :
func (b *ULZRoomServiceBackend) JoinRoom(ctx context.Context, req *pb.RoomReq) (*pb.RoomResp, error) {
	start := time.Now()
	b.mu.Lock()
	wkbox := b.searchAliveClient()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
		(wkbox).Preserve(false)
	}()

	var tmp pb.Room
	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		return nil, status.Errorf( codes.NotFound, err.Error())
	}
	if tmp.Dueler == nil  && req.IsDuel {
		tmp.Dueler = req.User
	}

	return &pb.RoomResp{
		Result:   &tmp,
		ErrorMsg: nil,
	}, nil
}
