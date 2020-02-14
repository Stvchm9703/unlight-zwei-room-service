package serverCtlNoRedis

import (
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// GetRoomInfo :
func (b *ULZRoomServiceBackend) GetRoomInfo(ctx context.Context, req *pb.RoomReq) (*pb.Room, error) {
	start := time.Now()
	b.mu.Lock()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()

	roommgt, ok := b.Roomlist[req.Key]
	if !ok {
		return nil, status.Error(codes.NotFound, "ROOM_NOT_FOUND")
	}

	if roommgt.Password != "" && roommgt.Password != req.Password {
		return nil, status.Error(codes.PermissionDenied, "ROOM_PASSWORD_INV")
	}

	return &roommgt.Room, nil
}
