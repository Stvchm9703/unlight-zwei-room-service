package serverCtlNoRedis

import (
	pb "ULZRoomService/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetRoomInfo :
func (b *ULZRoomServiceBackend) QuickPair(ctx context.Context, request *pb.RoomCreateReq) (*pb.Room, error) {
	return nil, status.Error(codes.Unimplemented, "NOT_YET_IMPLEMENT")
}
