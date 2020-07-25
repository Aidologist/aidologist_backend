package chatRoom

import (
	"github.com/astaxie/beego"
	models2 "wkBackEnd/models"
)

type ChatRoomController struct {
	beego.Controller
}

// @Title Post
// @Description Company posting a task
// @Success 200 {string}
// @Failure 403
// @router /init/from/task [post]
func (control *ChatRoomController) InitChatRoomFromTask() {
	var chatRoom = models2.ChatRoom{

	}
	chatRoom.Create()
}
