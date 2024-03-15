package chat

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/realPointer/chat-server/pkg/chat_v1"
)

func (i *Implementation) CreateChat(ctx context.Context, _ *emptypb.Empty) (*desc.CreateChatResponse, error) {
	chatID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	i.channels[chatID.String()] = make(chan *desc.Message, 100)

	return &desc.CreateChatResponse{
		ChatId: chatID.String(),
	}, nil
}
