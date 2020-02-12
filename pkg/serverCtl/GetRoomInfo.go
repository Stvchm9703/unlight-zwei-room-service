package serverCtl

import (
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"

	"github.com/gogo/status"
)

// GetRoomInfo :
func (b *ULZRoomServiceBackend) GetRoomInfo(ctx context.Context, req *pb.RoomReq) (*pb.RoomResp, error) {
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
	if _, err := wkbox.GetPara(&req.RoomKey, &tmp); err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(500, err.Error())
	}
	return &pb.RoomResp{
		Result:   &tmp,
		ErrorMsg: nil,
	}, nil
}
