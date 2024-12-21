package user

import (
	"go-template/internal/infra/database"

	"github.com/google/uuid"
)

type Repository struct {
	database *database.Database
}

func (r *Repository) Create(name string, email string) (*User, error) {
	user := User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}

	if _, err := r.database.Client.Exec("INSERT INTO user (id, name, email) VALUES (?, ?, ?)", user.ID, user.Name, user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) FindAll(name string, email string) ([]User, error) {
	var users []User

	rows, err := r.database.Client.Query("SELECT id, name, email FROM user")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *Repository) FindOne(id string) (*User, error) {
	var user User

	if err := r.database.Client.QueryRow("SELECT id, name, email FROM user WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Update(id string, name string, email string) (*User, error) {
	user := User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	if _, err := r.database.Client.Exec("UPDATE user SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Delete(id string) error {
	if _, err := r.database.Client.Exec("DELETE FROM user WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}
