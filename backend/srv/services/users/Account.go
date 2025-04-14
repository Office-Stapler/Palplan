package users

import (
	"context"

	sqrl "github.com/Masterminds/squirrel"
	"github.com/Office-Stapler/Palplan/backend/srv/db"
	"github.com/Office-Stapler/Palplan/backend/srv/model/users"
)

const TABLE_NAME = "accounts"

type AccountService struct {
	DB *db.DB
}

func NewAccountService(DB *db.DB) *AccountService {
	return &AccountService{
		DB,
	}
}

func (a *AccountService) GetAccountByID(ctx context.Context, id int64) (*users.Account, error) {
	sqlBuilder := sqrl.Select("*").From(TABLE_NAME).Where(sqrl.Eq{"account_id": id})
	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := a.DB.QueryRow(ctx, query, args)
	account := &users.Account{}
	err = row.Scan(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
