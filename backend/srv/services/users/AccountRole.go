package users

import (
	"context"
	"database/sql"
	"fmt"

	sqrl "github.com/Masterminds/squirrel"
	"github.com/Office-Stapler/Palplan/backend/srv/db"
	"github.com/Office-Stapler/Palplan/backend/srv/model/users"
)

const TABLE_NAME_ROLES = "account_roles"

type AccountRoleService struct {
	DB *db.DB
}

func NewAccountRoleService(DB *db.DB) *AccountRoleService {
	return &AccountRoleService{
		DB,
	}
}

func (ar *AccountRoleService) GetRoleByName(ctx context.Context, roleName string) (*users.AccountRole, error) {
	sqlBuilder := sqrl.Select("*").From(TABLE_NAME_ROLES).Where(sqrl.Eq{"role_name": roleName})
	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := ar.DB.QueryRow(ctx, query, args)
	if row == nil {
		return nil, fmt.Errorf("database pool is nil")
	}

	accountRole := &users.AccountRole{}
	err = row.Scan(accountRole)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("role not found")
	}
	if err != nil {
		return nil, err
	}

	return accountRole, nil
}

func (ar *AccountRoleService) GetRoleByID(ctx context.Context, id int64) (*users.AccountRole, error) {
	sqlBuilder := sqrl.Select("*").From(TABLE_NAME_ROLES).Where(sqrl.Eq{"rold_id": id})
	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := ar.DB.QueryRow(ctx, query, args)
	if row == nil {
		return nil, fmt.Errorf("database pool is nil")
	}

	accountRole := &users.AccountRole{}
	err = row.Scan(accountRole)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("role not found")
	}
	if err != nil {
		return nil, err
	}

	return accountRole, nil
}
