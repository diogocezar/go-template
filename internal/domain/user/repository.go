package user

import (
	"fmt"
	"go-template/internal/infra/database"
	"go-template/pkg/logger"
	"go-template/pkg/utils"

	"github.com/google/uuid"
)

type UserRepository struct {
	database *database.Database
}

func NewRepository(database *database.Database) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (r *UserRepository) Create(name string, email string, password string) (*User, error) {
	user := User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Passsword: password,
	}

	if _, err := r.database.Client.Exec("INSERT INTO user (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, utils.GeneratePassword(user.Passsword)); err != nil {
		logger.Error(fmt.Sprintf("Error creating user: %v", err))
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindAll() ([]User, error) {
	var users []User

	rows, err := r.database.Client.Query("SELECT id, name, email, password FROM user")
	if err != nil {
		logger.Error(fmt.Sprintf("Error finding all users: %v", err))
		return nil, err
	}

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Passsword); err != nil {
			logger.Error(fmt.Sprintf("Error scanning user: %v", err))
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) FindOne(id string) (*User, error) {
	var user User

	if err := r.database.Client.QueryRow("SELECT id, name, email FROM user WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		logger.Error(fmt.Sprintf("Error geting one user user: %v", err))
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(id string) (*User, error) {
	var user User

	if err := r.database.Client.QueryRow("SELECT id, name, email, password FROM user WHERE email = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Passsword); err != nil {
		logger.Error(fmt.Sprintf("Error geting one user user: %v", err))
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(id string, name string, email string) (*User, error) {
	user := User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	if _, err := r.database.Client.Exec("UPDATE user SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID); err != nil {
		logger.Error(fmt.Sprintf("Error updating user: %v", err))
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Delete(id string) error {
	if _, err := r.database.Client.Exec("DELETE FROM user WHERE id = ?", id); err != nil {
		logger.Error(fmt.Sprintf("Error deleting user: %v", err))
		return err
	}
	return nil
}
