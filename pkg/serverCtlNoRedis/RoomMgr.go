package serverCtlNoRedis

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"errors"
	"log"
)

// ======================================================================================================
// RoomMgr : Room Manager
type RoomMgr struct {
	pb.Room
	clientConn map[string]*pb.RoomService_ServerBroadcastServer
}

// ----------------------------------------------------------------------------------------------------
// roommgr.clientConn

func (rm *ULZRoomServiceBackend) GetStream(roomKey *string, userId *string) *pb.RoomService_ServerBroadcastServer {
	a, ok := rm.Roomlist[*roomKey]
	b, ok := a.clientConn[*userId]
	if ok {
		return b
	}
	return nil
}

func (rm *ULZRoomServiceBackend) AddStream(roomKey *string, userId *string, stream *pb.RoomService_ServerBroadcastServer) (bool, error) {
	// _, ok := rm.bc_stream[user_id]

	if rm.Roomlist == nil {
		log.Println("help this")
	} else {
		log.Println("this-roomlist : ", rm.Roomlist)
	}

	a, ok := rm.Roomlist[*roomKey]
	if !ok {
		return false, errors.New("ROOM_NOT_EXIST")
	}

	_, ok = a.clientConn[*userId]
	if ok {
		return false, errors.New("USER_EXIST")
	}
	a.clientConn[*userId] = stream
	return true, nil
}

func (rm *ULZRoomServiceBackend) DelStream(roomKey *string, userId *string) (bool, error) {
	log.Println("Del Stream:", userId)
	a, ok := rm.Roomlist[*roomKey]
	if !ok {
		return false, errors.New("ROOM_NOT_EXIST")
	}
	if a.clientConn[*userId] != nil {
		*(a.clientConn[*userId]) = nil
		delete(a.clientConn, *userId)
		return true, nil
	}
	return false, errors.New("StreamNotExist")
}

// func (rm *ULZRoomServiceBackend) BroadCast(roomkey *string, from *string, message *pb.RoomMsg) error {
// 	log.Println("BS!", message)
// 	log.Println(rm.Roomlist[*roomkey])
// 	if rm.Roomlist == nil {
// 		log.Println("help this")
// 	} else {
// 		log.Println("this-roomlist : ", rm.Roomlist)
// 	}
// 	rmb, ok := rm.Roomlist[*roomkey]
// 	if !ok {
// 		log.Println("room not exist")
// 		return errors.New("ROOM_NOT_EXIST")
// 	}
// 	for k, v := range rmb.clientConn {
// 		if k != *from {
// 			(*v).Send(message)
// 		}
// 	}
// 	return nil
// }

// ---------------------------------------------------------------------------------------------
// RoomStreamBox Controlling

func (rm *RoomMgr) ClearAll() {
	log.Println("ClearAll Proc")
	for _, vc := range rm.clientConn {
		(*vc).Send(cm.MsgSystShutdown(&rm.Room.Key))
	}
	for k := range rm.clientConn {
		*(rm.clientConn[k]) = nil
		delete(rm.clientConn, k)
	}
	return
}
