package serverCtlNoRedis

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateRoom :
func (b *ULZRoomServiceBackend) UpdateRoom(ctx context.Context, req *pb.RoomCreateReq) (*pb.Room, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	start := time.Now()
	b.mu.Lock()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()

	rmg, ok := b.RoomList[req.Key]
	if !ok {
		return nil, status.Error(codes.NotFound, "ROOM_NOT_FOUND")
	}

	rmg.Room.Password = req.Password
	rmg.Room.CostLimitMax = req.CostLimitMax
	rmg.Room.CostLimitMin = req.CostLimitMin
	rmg.Room.CharCardNvn = req.CharCardNvn
	rmg.Room.CharCardLimitMax = req.CharCardLimitMax
	rmg.Room.CharCardLimitMin = req.CharCardLimitMin

	return rmg.Room, nil
}
