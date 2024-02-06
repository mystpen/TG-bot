package messages

type MessageSender interface {
	SendMessage(text string, userid int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{tgClient: tgClient}
}

type Message struct {
	Text   string
	UserID int64
}

func (s *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		s.tgClient.SendMessage("hello", msg.UserID)
		return nil
	}
	s.tgClient.SendMessage("Incorrect command", msg.UserID)
	return nil
}
