package repository

import (
	"context"

	"github.com/immersivesky/go-rest-postgresql/internal/core/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	client *pgxpool.Pool
}

func NewDB(dsn string) *DB {
	client, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background()); err != nil {
		panic(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetChat(chatID int) *domain.Chat {
	chat := new(domain.Chat)

	if err := db.client.QueryRow(context.Background(), "SELECT chat_id FROM chats WHERE chat_id = $1;", chatID).Scan(&chat.ID); err != nil {
		return nil
	}

	return chat
}

func (db *DB) CreateChat(chatID int) bool {
	if err := db.client.QueryRow(context.Background(), "INSERT INTO chats(chat_id) VALUES($1);", chatID); err != nil {
		return false
	}

	return true
}
