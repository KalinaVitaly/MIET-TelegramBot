package telegrambotapi

type MessageBot struct {
	Message  string
	UserName string
}

func CreateNewMessage(message string, userName string) *MessageBot {
	return &MessageBot{
		Message:  message,
		UserName: userName,
	}
}
