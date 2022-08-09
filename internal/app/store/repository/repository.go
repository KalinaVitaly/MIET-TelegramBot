package repository

import "database/sql"

type Repository struct {
	db             *sql.DB
	userRepository *UserRepository
}

func NewRepository(_db *sql.DB) *Repository {
	return &Repository{
		db: _db,
	}
}

func (r *Repository) User() *UserRepository {
	if r.userRepository != nil {
		return r.userRepository
	}

	r.userRepository = &UserRepository{
		repository: r,
	}

	return r.userRepository
}
