package serverCtl

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
	// "time"
)

var (
	RunningConfig *cf.CfAPIServer
)

func (this *ULZRoomServiceBackend) ServerBroadcast(rReq *pb.RoomReq, stream pb.RoomService_ServerBroadcastServer) error {
	// log.Println("\nServer Broadcast Connect\n methods: ServerBroadcast")
	cm.PrintReqLog(nil, "server-broadcast", rReq)
	return status.Error(codes.Internal, "SkipImpl")
}

func (this *ULZRoomServiceBackend) BroadcastInfo(ctx context.Context, rReq *pb.RoomReq) (*pb.RoomBroadcastInfo, error) {
	cm.PrintReqLog(nil, "server-broadcast", rReq)
	var tmp pb.Room
	wkbox := this.searchAliveClient()
	if _, err := wkbox.GetPara(&rReq.Key, &tmp); err != nil {
		return nil, status.Error(codes.NotFound, "not exist")
	}
	return &pb.RoomBroadcastInfo{
		Key:      rReq.Key,
		Ip:       RunningConfig.IP,
		Port:     int32(RunningConfig.PollingPort),
		Protocal: "ws",
		Securl:   nil,
	}, nil
}

func (this *ULZRoomServiceBackend) SendMessage(ctx context.Context, msg *pb.RoomMsg) (*pb.Empty, error) {
	cm.PrintReqLog(ctx, "send-message", msg)
	this.BroadCast(msg)
	return &pb.Empty{}, nil
}

func (this *ULZRoomServiceBackend) BroadCast(msg *pb.RoomMsg) {
	this.castServer.Broadcast(msg)
	return
}

func (rsb *ULZRoomServiceBackend) RunWebSocketServer(config cf.CfAPIServer) error {
	hub := ws.NewHub()
	go hub.Run()
	router := gin.New()
	router.GET("/:roomId", Wrapfunc(rsb, hub))
	rsb.castServer = hub
	RunningConfig = &config
	return router.Run(":" + strconv.Itoa(config.PollingPort))
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
	reqKey := c.Param("roomId")

	var tmp pb.Room
	wkbox := rsb.searchAliveClient()
	if _, err := wkbox.GetPara(&reqKey, &tmp); err != nil {
		c.AbortWithStatus(412)
	}
	client := ws.NewClient(reqKey, hub, conn)
	go client.WritePump()
	go client.ReadPump()
}
