package serverCtlNoRedis

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
	cm.PrintReqLog(ctx, "join-room", req)
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()

	tmp, ok := b.Roomlist[req.Key]
	if !ok {
		return nil, status.Error(codes.NotFound, "ROM_NOT_EXIST")
	}

	if tmp.Room.Dueler == nil && req.IsDuel {
		tmp.Room.Dueler = req.User
	}

	return &tmp.Room, nil
}
