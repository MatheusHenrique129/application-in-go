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
	Update(ctx context.Context, user *model.User) (int64, RepoError)
	Delete(ctx context.Context, userID int64) RepoError
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

func (u *userRepository) Update(ctx context.Context, user *model.User) (int64, RepoError) {

	qs := goquMysql.
		Update(consts.TableUser).
		Set(goqu.Record{
			consts.FieldName:        user.Name,
			consts.FieldCpf:         user.CPF,
			consts.FieldEmail:       user.Email,
			consts.FieldAddress:     user.Address,
			consts.FieldPhoneNumber: user.PhoneNumber,
			consts.FieldGender:      user.Gender,
			consts.FieldPassword:    user.Password,
			consts.FieldBirthDate:   user.BirthDate,
			consts.FieldUpdatedAt:   user.UpdatedAt,
		}).
		Where(goqu.Ex{
			consts.FieldUserID: user.ID,
		}).
		Prepared(true)

	query, args, _ := qs.ToSQL()

	rows, err := u.db.Exec(query, args...)
	if err != nil {
		u.logger.Errorf(ctx, "Could NOT update user with id: %v.", err, user.ID)
		return 0, NewFromDatabaseError(err)
	}

	count, _ := rows.RowsAffected()

	return count, nil
}

func (u *userRepository) Delete(ctx context.Context, userID int64) RepoError {

	qs := goquMysql.
		Delete(consts.TableUser).
		Where(goqu.Ex{
			consts.FieldUserID: userID,
		}).
		Prepared(true)

	query, args, _ := qs.ToSQL()

	_, err := u.db.Exec(query, args...)
	if err != nil {
		u.logger.Errorf(ctx, "Could NOT delete user '%s'.", err, userID)

		return NewFromDatabaseError(err)
	}

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

	u.logger.Debugf(ctx, "Obtaining successful user with id: %v.!", userID)
	return &user, nil
}

func NewUserRepository(conf *config.Config, db *sql.DB) UserRepository {
	logger := util.NewLogger("User Repository")

	return &userRepository{
		BaseRepository: NewBaseRepository(conf, db, logger),
	}
}
