package bridge

type IMsgSender interface {
	Send(msg string) error
}

type EmailMsgSender struct {
	emails []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{
		emails: emails,
	}
}

func (s *EmailMsgSender) Send(msg string) error {
	// send message
	return nil
}

type INotification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	sender IMsgSender
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{
		sender: sender,
	}
}

func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
