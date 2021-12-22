package protobuf

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"learnGo/testDemo/protobuf/pb"
	"testing"
)

func TestProto(t *testing.T) {
	u := &pb.UserInfo{
		UserType: 1,
		UserName: "Jok",
		UserInfo: "I am a woker!",
	}

	//序列化
	userdata, err := proto.Marshal(u)
	if err != nil {
		fmt.Println("Marshaling error: ", err)
	}

	//反序列化
	encodingmsg := &pb.UserInfo{}
	err = proto.Unmarshal(userdata, encodingmsg)

	if err != nil {
		fmt.Println("Unmarshaling error: ", err)
	}

	fmt.Printf("GetUserType: %d\n", encodingmsg.GetUserType())
	fmt.Printf("GetUserName: %s\n", encodingmsg.GetUserName())
	fmt.Printf("GetUserInfo: %s\n", encodingmsg.GetUserInfo())
}
