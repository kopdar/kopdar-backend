package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kopdar/kopdar-backend/internal/user"
	"github.com/kopdar/kopdar-backend/pkg/pglib"
)

// UserRepository implements operations for working with the user data storage.
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*user.Model, error) {
	emptyParams := map[string]interface{}{}
	var users []*user.Model
	err := pglib.Query(ctx, r.db, queryFindAll, emptyParams, &users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, user.ErrOrderNotFound
	}
	return users, nil
}

const userQueryBase = `
select
	id,
	name,
	email,
	phone_number,
	pin
`

const queryFindAll = userQueryBase + `
FROM "user"
`
