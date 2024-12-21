package user

import "go-template/internal/infra/database"

type Repository struct {
	database *database.Database
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}
