package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/KoLLlaka/http-rest-api/internal/app/model"
	"github.com/KoLLlaka/http-rest-api/internal/app/store"
)

const (
	usersTable = "users"
)

type UserRepository struct {
	store *Store
}

// Create user
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	// INSERT INTO users (email, encrypted_password) VALUES ("Den@mail.ru", "12345")
	stmt := fmt.Sprintf("INSERT INTO %s (email, encrypted_password) VALUES (?, ?)", usersTable)

	result, err := r.store.db.Exec(stmt, u.Email, u.EncryptedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = int(id)

	return nil
}

// Find User by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	// SELECT id, email, encrypted_password FROM users WHERE email = "Den@mail.ru"
	stmt := fmt.Sprintf("SELECT id, email, encrypted_password FROM %s WHERE email = ?", usersTable)

	if err := r.store.db.QueryRow(stmt, email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// Find User by ID
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	// SELECT id, email, encrypted_password FROM users WHERE id = 1
	stmt := fmt.Sprintf("SELECT id, email, encrypted_password FROM %s WHERE id = ?", usersTable)

	if err := r.store.db.QueryRow(stmt, id).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
