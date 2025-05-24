package postgres

import (
	"SmartHouseAPI/models"
	"context"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) UpdateUser(ctx context.Context, user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `UPDATE users SET password_hash = $1 WHERE username = $2`
	_, err = r.DB.ExecContext(ctx, query, string(hashedPassword), user.Username)
	return err
}
