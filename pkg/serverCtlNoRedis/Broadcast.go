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
}

func (this *ULZRoomServiceBackend) BroadcastInfo(ctx context.Context, rReq *pb.RoomReq) (*pb.RoomBroadcastInfo, error) {
	cm.PrintReqLog(nil, "server-broadcast", rReq)
	return nil, status.Error(codes.Internal, "SkipImpl")
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
