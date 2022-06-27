package telegrambotapi

type CommandBot struct {
	Command  string
	UserName string
}

func CreateNewCommand(command string, userName string) *CommandBot {
	return &CommandBot{
		Command:  command,
		UserName: userName,
	}
}
