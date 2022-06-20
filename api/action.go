package api

import (
	"GameServerV2/core"
	"GameServerV2/pb"
	"GameServerV2/ziface"
	"GameServerV2/znet"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type ActionApi struct {
	znet.BaseRouter
}

func (*ActionApi) Handle(request ziface.IRequest) {
	//1. 将客户端传来的proto协议解码
	actionMsg := &pb.PlayerAction{}
	err := proto.Unmarshal(request.GetData(), actionMsg)
	if err != nil {
		fmt.Println("Action: PlayerAction Unmarshal error ", err)
		return
	}

	//if (msg.X < float32(core.AOI_MIN_X)) || (msg.X > float32(core.AOI_MAX_X)) || (msg.Y > float32(core.AOI_MAX_Y)) || (msg.Y < float32(core.AOI_MIN_Y)) {
	//	//fmt.Println("out if AOI!")
	//	//TODO 广播下线消息
	//	return
	//}

	//2. 得知当前的消息是从哪个玩家传递来的,从连接属性pID中获取
	pID, err := request.GetConnection().GetProperty("pID")
	if err != nil {
		fmt.Println("GetProperty pID error", err)
		request.GetConnection().Stop()
		return
	}

	//fmt.Printf("user pID = %d , move(%f,%f,%f,%f)\n", pID, msg.X, msg.Y, msg.Z, msg.V)

	//3. 根据pID得到player对象
	player := core.WorldMgrObj.GetPlayerByPID(pID.(int32))

	//4. 让player对象发起动作信息广播
	player.UpdateAction(actionMsg)

	//if actionMsg.Input != nil {
	//	fmt.Println("ActionApi is called：" + actionMsg.String())
	//}
}
