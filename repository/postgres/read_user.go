package postgres

import (
	"SmartHouseAPI/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// ReadUser
// возвращает объект User, где в поле password будет указано хэш-значение
func (r *Repository) ReadUser(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT password_hash FROM users WHERE username = $1 LIMIT 1`
	row := r.DB.QueryRowContext(ctx, query, username)

	var passwordHash string
	if err := row.Scan(&passwordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with username '%s' not found", username)
		}
		return nil, err
	}

	return &models.User{
		Username: username,
		Password: passwordHash,
	}, nil
}
