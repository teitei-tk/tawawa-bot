package util

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func GetReplySource(event *linebot.Event) string {
	var replySource string
	switch event.Source.Type {
	default:
	case linebot.EventSourceTypeUser:
		replySource = event.Source.UserID
	case linebot.EventSourceTypeRoom:
		replySource = event.Source.RoomID
	case linebot.EventSourceTypeGroup:
		replySource = event.Source.GroupID
	}

	return replySource
}
