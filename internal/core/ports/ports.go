package ports

import (
	"github.com/immersivesky/go-rest-postgresql/internal/core/domain"
)

type DB interface {
	GetChat(int) *domain.Chat
	CreateChat(int) bool
}
