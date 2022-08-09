package store

import "MIET-TelegramBot/internal/app/store/repository"

type Store interface {
	User() repository.UserRepository
}
