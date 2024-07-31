package adapter

import (
	"context"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

type MessageAdapter struct {
	ConversationChannels []*model.ConversationChannel
}

var adapter *MessageAdapter

func NewMessageAdapter() *MessageAdapter {
	if adapter == nil {
		adapter = &MessageAdapter{
			ConversationChannels: []*model.ConversationChannel{},
		}
	}

	return adapter
}

func (m *MessageAdapter) AddConversationChannel(channel chan []*model.Message, conversationID string) {
	convChannel := &model.ConversationChannel{
		Channel:        channel,
		ConversationID: conversationID,
	}

	m.ConversationChannels = append(m.ConversationChannels, convChannel)
}

func (m *MessageAdapter) RemoveConversationChannel(convID string) (convChannel *model.ConversationChannel) {
	for i, c := range m.ConversationChannels {
		if c.ConversationID == convID {
			m.ConversationChannels = append(m.ConversationChannels[:i], m.ConversationChannels[i+1:]...)
			return c
		}
	}

	return nil
}

func (m *MessageAdapter) CloseConversationChannel(ctx context.Context, channel chan []*model.Message, convID string) {
	go func() {
		<-ctx.Done()
		m.RemoveConversationChannel(convID)

		close(channel)
	}()
}

func (m *MessageAdapter) BroadcastMessage(message *model.Message) {
	go func() {
		for _, c := range m.ConversationChannels {
			if c.ConversationID == message.ConversationID {
				c.Channel <- []*model.Message{message}
			}
		}
	}()
}

func (m *MessageAdapter) BroadcastMessages(messages []*model.Message) {
	go func() {
		for _, c := range m.ConversationChannels {
			for _, message := range messages {
				if c.ConversationID == message.ConversationID {
					c.Channel <- messages
				}
			}
		}
	}()
}
