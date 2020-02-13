package common

import (
	pb "ULZRoomService/proto"
)

// MsgSystShutdown : Message-System-Shutdown
func MsgSystShutdown(key *string) *pb.RoomMsg {
	return &pb.RoomMsg{
		Key:     *key,
		FormId:  "SYSTEM",
		ToId:    "ALL_USER",
		Message: "host is quited room, this room is going abort",
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	}
}

func MsgHostQuitRoom(key *string) *pb.RoomMsg {
	return &pb.RoomMsg{
		Key:     *key,
		FormId:  "SYSTEM",
		ToId:    "ALL_USER",
		Message: "host is quited room, this room is going abort",
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	}
}
