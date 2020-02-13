package common

import (
	pb "ULZRoomService/proto"
	"fmt"
)

// MsgSystShutdown : Message-System-Shutdown
func MsgSystShutdown(key *string) *pb.RoomMsg {
	return &pb.RoomMsg{
		Key:     *key,
		FormId:  "SYSTEM",
		ToId:    "ALL_USER",
		Message: "System is going shutdown, this room is going abort",
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	}
}

func MsgHostQuitRoom(key *string, username *string) *pb.RoomMsg {
	return &pb.RoomMsg{
		Key:     *key,
		FormId:  "SYSTEM",
		ToId:    "ALL_USER",
		Message: fmt.Sprintf("Host Player <%v> is quited room, this room is going abort", *username),
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	}
}

func MsgUserQuitRoom(key *string, userId *string, username *string) *pb.RoomMsg {
	return &pb.RoomMsg{
		Key:     *key,
		FormId:  *userId,
		ToId:    "ALL_USER",
		Message: fmt.Sprintf("User <%v> is quited room", *username),
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	}
}
