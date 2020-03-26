package serverCtlNoRedis

import (
	cm "ULZRoomService/pkg/common"
	cf "ULZRoomService/pkg/config"
	ws "ULZRoomService/pkg/websocket"
	pb "ULZRoomService/proto"
	"context"
	"fmt"
	"log"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (this *ULZRoomServiceBackend) ServerBroadcast(rReq *pb.RoomReq, stream pb.RoomService_ServerBroadcastServer) error {
	// log.Println("\nServer Broadcast Connect\n methods: ServerBroadcast")
	cm.PrintReqLog(nil, "server-broadcast", rReq)
	return status.Error(codes.Internal, "SkipImpl")
	// _, err := this.AddStream(&rReq.Key, &rReq.User.Id, &stream)
	// if err != nil {
	// 	return status.Error(codes.NotFound, err.Error())
	// }

	// go func() {
	// 	<-stream.Context().Done()
	// 	log.Println("close done")
	// 	_, err := this.DelStream(&rReq.Key, &rReq.User.Id)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	this.BroadCast(&rReq.Key, &rReq.User.Id,
	// 		cm.MsgUserQuitRoom(&rReq.Key, &rReq.User.Id, &rReq.User.Name))
	// }()
	// for {
	// }
}

func (this *ULZRoomServiceBackend) SendMessage(ctx context.Context, msg *pb.RoomMsg) (*pb.Empty, error) {
	cm.PrintReqLog(nil, "server-broadcast:msg", msg)
	// this.BroadCast(&msg.Key, &msg.FromId, msg)
	return &pb.Empty{}, nil
}

func (rsb *ULZRoomServiceBackend) BroadCast(cp *pb.RoomMsg) error {
	rsb.castServer.Broadcast(cp)
	return nil
}
func (rsb *ULZRoomServiceBackend) RunWebSocketServer(config cf.CfAPIServer) error {
	hub := ws.NewHub()
	go hub.Run()
	router := gin.New()
	router.GET("/:roomId", Wrapfunc(rsb, hub))
	rsb.castServer = hub
	return router.Run(config.IP + ":" + strconv.Itoa(config.PollingPort))
}

// wraper to gin handler
func Wrapfunc(rsb *ULZRoomServiceBackend, hub *ws.SocketHub) gin.HandlerFunc {
	return func(c *gin.Context) {
		serveWs(rsb, hub, c)
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(rsb *ULZRoomServiceBackend, hub *ws.SocketHub, c *gin.Context) {
	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(c.Param("roomId"))
	found := false
	for _, rm := range rsb.Roomlist {
		if rm.Key == c.Param("roomId") {
			found = true
		}
	}
	if found {
		client := ws.NewClient(c.Param("roomId"), hub, conn)
		go client.WritePump()
		go client.ReadPump()
	} else {
		c.AbortWithStatus(412)
	}
}
