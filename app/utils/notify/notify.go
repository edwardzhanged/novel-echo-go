package notify

type Notifier interface {
	Compose() error
	Notify() error
}

type Email struct {
	To      string
	Subject string
}

func (email *Email) Compose() error {
	return nil
}

func (email *Email) Notify() error {
	return nil
}
