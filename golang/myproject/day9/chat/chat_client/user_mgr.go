package main

import "fmt"
import "myproject/day9/chat/common"
import "myproject/day9/chat/proto"

var onlineUserMap map[int]*common.User = make(map[int]*common.User, 16)

func outputUserOnline() {
	fmt.Println("Online User List:")
	for id, _ := range onlineUserMap {
		if id == userId {
			continue
		}
		fmt.Println("user:", id)
	}
}

func updateUserStatus(userStatus proto.UserStatusNotify) {
	user, ok := onlineUserMap[userStatus.UserId]
	if !ok {
		user = &common.User{}
		user.UserId = userStatus.UserId
	}

	user.Status = userStatus.Status
	onlineUserMap[user.UserId] = user

	outputUserOnline()
}
