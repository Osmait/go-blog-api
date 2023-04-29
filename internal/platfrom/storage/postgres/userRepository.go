package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/osmait/blog-go/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(id, email, password) VALUES ($1, $2, $3)", user.ID().String(), user.GetEmail(), user.GetPassword())
	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM users WHERE email =$1", email)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = domain.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {

			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil

}
