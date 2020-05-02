package common

import (
	pb "ULZRoomService/proto"
	"encoding/json"
	"fmt"
)

// MsgSystShutdown : Message-System-Shutdown
func MsgSystShutdown(key *string) *pb.RoomMsg {
	return &pb.RoomMsg{
		Key:     *key,
		FromId:  "SYSTEM",
		ToId:    "ALL_USER",
		Message: "System is going shutdown, this room is going abort",
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	}
}

func MsgHostQuitRoom(key *string, username *string) *pb.RoomMsg {
	return &pb.RoomMsg{
		Key:     *key,
		FromId:  "SYSTEM",
		ToId:    "ALL_USER",
		Message: fmt.Sprintf("Host Player <%v> is quited room, this room is going abort", *username),
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	}
}

func MsgUserQuitRoom(key *string, userId *string, username *string) *pb.RoomMsg {
	return &pb.RoomMsg{
		Key:     *key,
		FromId:  *userId,
		ToId:    "ALL_USER",
		Message: fmt.Sprintf("User <%v> is quited room", *username),
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	}
}

func MsgHostUpdateRoom(key *string, pw *pb.Room) *pb.RoomMsg {
	jsonstr, _ := json.Marshal(pw)
	return &pb.RoomMsg{
		Key:     *key,
		FromId:  "SYSTEM",
		ToId:    "ALL_USER",
		Message: fmt.Sprintf("UPDATE_ROOM:pw::%s", string(jsonstr)),
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	}
}
