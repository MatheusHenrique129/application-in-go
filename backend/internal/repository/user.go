package repository

import (
	"context"
	"database/sql"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/consts"
	"github.com/MatheusHenrique129/application-in-go/internal/model"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/doug-martin/goqu/v9"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) RepoError
	FindByID(ctx context.Context, userID int64) (*model.User, RepoError)
}

type userRepository struct {
	BaseRepository
}

func (u *userRepository) Create(ctx context.Context, user *model.User) RepoError {

	qs := goquMysql.
		Insert(consts.TableUser).
		Rows(user).
		Prepared(true)

	query, args, _ := qs.ToSQL()

	row, err := u.db.Exec(query, args...)

	if err != nil {
		u.logger.Error(ctx, "Error inserting user from database", err)
		return NewFromDatabaseError(err)
	}

	id, err := row.LastInsertId()

	if err != nil {
		u.logger.Error(ctx, "Could NOT get last insert ID when creating user.", err)
		return NewFromDatabaseError(err)
	}

	user.ID = id

	return nil
}

func (u *userRepository) FindByID(ctx context.Context, userID int64) (*model.User, RepoError) {

	var user model.User

	found, err := goquMysql.
		DB(u.db).
		Select(&model.User{}).
		From(consts.TableUser).
		Where(goqu.Ex{
			consts.FieldUserID: userID,
		}).
		Prepared(true).
		ScanStruct(&user)

	if err != nil {
		return nil, NewFromDatabaseError(err)
	}

	if !found {
		return nil, nil
	}

	u.logger.Debugf(ctx, "Obtaining successful '%s' user!", userID)
	return &user, nil
}

func NewUserRepository(conf *config.Config, db *sql.DB) UserRepository {
	logger := util.NewLogger("User Repository")

	return &userRepository{
		BaseRepository: NewBaseRepository(conf, db, logger),
	}
}
