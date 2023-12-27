package postgres

import (
	"Interface_droch_3/internal/model"
	"github.com/jmoiron/sqlx"
)

type TodoPostgres struct {
	db *sqlx.DB
}

func NewTodoPostgres(db *sqlx.DB) *TodoPostgres {
	return &TodoPostgres{db: db}
}

func (r *TodoPostgres) Set(user *model.User) error {

	sqlStatement := `
INSERT INTO users(id, name,username,password_hash) VALUES ($1,$2,$3,$4)
`
	_, err := r.db.Exec(sqlStatement, user.Id, user.Name, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoPostgres) GetById(id int64) (*model.User, error) {
	sqlStatement := "SELECT id,name,username,password_hash FROM users WHERE id=$1"

	var user model.User

	err := r.db.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *TodoPostgres) CheckById(id int64) (bool, error) {
	sqlStatement := "SELECT COUNT(*) FROM users WHERE id=$1"

	var count int

	err := r.db.QueryRow(sqlStatement, id).Scan(&count)

	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *TodoPostgres) Delete(id int64) error {

	sqlStatement := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoPostgres) GetAllId() ([]int64, error) {

	sqlStatement := "SELECT id FROM users;"

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	var ids []int64

	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}
