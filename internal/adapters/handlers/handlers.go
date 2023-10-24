package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/immersivesky/go-rest-postgresql/internal/core/ports"
)

type HTTPHandlers struct {
	db ports.DB
}

func NewHTTPHandlers(db ports.DB) *HTTPHandlers {
	return &HTTPHandlers{
		db: db,
	}
}

func (h *HTTPHandlers) GetChat(ctx *fiber.Ctx) error {
	chat := h.db.GetChat(1)

	return ctx.JSON(chat)
}

func (h *HTTPHandlers) CreateChat(ctx *fiber.Ctx) error {
	ok := h.db.CreateChat(1)

	return ctx.JSON(ok)
}
