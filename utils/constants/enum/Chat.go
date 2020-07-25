package enum

type MessageType int

const (
	ChatText = 1
	ChatLink = 2
	ChatPhoto = 101
	ChatAlbum = 102
	ChatVideo = 103
	ChatFile = 104
	ChatSound = 105
	ChatResource = 151
	ChatGitHub = 201
	ChatPinterest = 202
	ChatKaggle = 203
	ChatStackexchange = 205
)

type ChatRoomType int

const (
	ProjectChat = 1
	TaskChat = 2
)