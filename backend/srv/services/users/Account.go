package users

import (
	"context"
	"database/sql"
	"fmt"

	sqrl "github.com/Masterminds/squirrel"
	"github.com/Office-Stapler/Palplan/backend/srv/db"
	"github.com/Office-Stapler/Palplan/backend/srv/dbmodel"
	"github.com/Office-Stapler/Palplan/backend/srv/model/users"
)

const TABLE_NAME_ACCOUNTS = "accounts"

type AccountService struct {
	DB  *db.DB
	Ars *AccountRoleService
}

func NewAccountService(DB *db.DB, Ars *AccountRoleService) *AccountService {
	return &AccountService{
		DB,
		Ars,
	}
}

func (a *AccountService) GetAccountByID(ctx context.Context, id int64) (*users.Account, error) {
	sqlBuilder := sqrl.Select("*").From(TABLE_NAME_ACCOUNTS).Where(sqrl.Eq{"account_id": id})
	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := a.DB.QueryRow(ctx, query, args)
	if row == nil {
		return nil, fmt.Errorf("database pool is nil")
	}

	dbAccount := &dbmodel.Account{}
	err = row.Scan(dbAccount)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("account not found")
	}

	if err != nil {
		return nil, err
	}
	role, err := a.Ars.GetRoleByID(ctx, dbAccount.ID)
	if err != nil {
		return nil, err
	}

	account := &users.Account{
		CreatedAt: dbAccount.CreatedAt,
		Email:     dbAccount.Email,
		ID:        dbAccount.ID,
		UpdatedAt: dbAccount.UpdatedAt,
		Role:      role,
	}
	return account, nil
}
