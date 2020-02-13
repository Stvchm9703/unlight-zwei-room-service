package serverCtl

import (
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
	defer func() {
		wkbox.Preserve(false)
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()
	var room pb.Room
	_, err := wkbox.GetPara(&req.Key, &room)
	if err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if room == nil {
		return nil, status.Error(codes.NotFound, "ROOM_NOT_FOUND")
	}

	room.Password = req.Password
	room.CostLimitMax = req.CostLimitMax
	room.CostLimitMin = req.CostLimitMin
	room.CharCardNvn = req.CharCardNvn
	room.CharCardLimitMax = req.CharCardLimitMax
	room.CharCardLimitMin = req.CharCardLimitMin

	wkbox.UpdatePara(&room.Key, room)

	// log.Println("b.RoomList", b.Roomlist)
	return nil, nil

}
