package mq

import (
	"fmt"

	pkgmq "sunflower-blog-svc/pkg/mq"
)

type SenderFactory struct {
	senders map[string]pkgmq.Sender
}

func NewSenderFactory() *SenderFactory {
	return &SenderFactory{
		senders: make(map[string]pkgmq.Sender),
	}
}

func (f *SenderFactory) AddSender(name string, sender pkgmq.Sender) {
	f.senders[name] = sender
}

func (f *SenderFactory) AddSenders(senders map[string]pkgmq.Sender) {
	for name, sender := range senders {
		f.senders[name] = sender
	}
}

func (f *SenderFactory) GetSender(name string) (pkgmq.Sender, error) {
	s, ok := f.senders[name]
	if !ok {
		return nil, fmt.Errorf("sender not found: %s", name)
	}
	return s, nil
}

func (f *SenderFactory) Close() error {
	for _, sender := range f.senders {
		_ = sender.Close()
	}
	return nil
}
