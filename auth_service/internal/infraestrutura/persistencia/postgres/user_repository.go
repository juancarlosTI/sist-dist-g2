package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgconn"
	dominio "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func isUniqueViolation(err error) bool {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		return pgErr.Code == "23505"
	}
	return false
}

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type UserRepositorySQL struct {
	db DBTX
}

func NewUserRepositorySQL(db DBTX) *UserRepositorySQL {
	return &UserRepositorySQL{db: db}
}
func (r *UserRepositorySQL) Salvar(ctx context.Context, user *dominio.User) error {
	query := `
	INSERT INTO users (id, email, name, password_hash, role_user)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID(),
		user.Email(),
		user.Nome(),
		user.Password(),
		user.Role(),
	)

	if err != nil {
		if isUniqueViolation(err) {
			return err
		}
	}

	return err
}

func (r *UserRepositorySQL) FindByID(ctx context.Context, id string) (*dominio.User, error) {

	query := `
	SELECT id, email, name, password_hash, role_user
	FROM users
	WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var (
		userID       string
		emailDB      string
		name         string
		passwordHash sql.NullString
		roleUser     string
	)

	err := row.Scan(&userID, &emailDB, &name, &passwordHash, &roleUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println("ERRO SCAN:", err)
		return nil, err
	}

	var passPtr *string
	if passwordHash.Valid {
		passPtr = &passwordHash.String
	}

	roles := shared_types.Role{
		Tipo: shared_types.RoleTipo(roleUser),
	}

	user := dominio.RehidratarUser(
		userID,
		emailDB,
		name,
		passPtr,
		roles,
	)

	return user, nil
}

func (r *UserRepositorySQL) FindByEmail(ctx context.Context, email string) (*dominio.User, error) {

	query := `
	SELECT id, email, name, password_hash, role_user
	FROM users
	WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email)

	var (
		userID       string
		emailDB      string
		name         string
		passwordHash sql.NullString
		roleUser     string
	)

	err := row.Scan(&userID, &emailDB, &name, &passwordHash, &roleUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println("ERRO SCAN:", err)
		return nil, err
	}

	var passPtr *string
	if passwordHash.Valid {
		passPtr = &passwordHash.String
	}

	roles := shared_types.Role{
		Tipo: shared_types.RoleTipo(roleUser),
	}

	user := dominio.RehidratarUser(
		userID,
		emailDB,
		name,
		passPtr,
		roles,
	)

	return user, nil
}
