package queries

import (
	"cars/internal/app/models"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func createHuman(dbConn *pgx.Conn, h *models.Human) error {
	sql, args, err := sq.
		Insert("people").
		Columns("name", "surname", "patronymic").
		Values(h.Name, h.Surname, h.Patronymic).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = dbConn.Exec(context.Background(), sql, args...)
	if err != nil {
		return err
	}

	return setHumanID(dbConn, h)
}

func setHumanID(dbConn *pgx.Conn, h *models.Human) error {
	sql, args, err := sq.
		Select("id").
		From("people").
		Where(
			sq.Eq{"name": h.Name},
			sq.Eq{"surname": h.Surname},
			sq.Eq{"patronymic": h.Patronymic}).
		OrderBy("id DESC").
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	return dbConn.QueryRow(context.Background(), sql, args...).Scan(&h.ID)
}

func updateCarOwner(dbConn *pgx.Conn, c *models.Car) error {
	_, err := dbConn.Exec(
		context.Background(),
		"UPDATE people SET name = $1, surname = $2, patronymic = $3 where id = (SELECT owner_id FROM cars WHERE reg_num = $4)",
		c.Owner.Name, c.Owner.Surname, c.Owner.Patronymic, c.RegNum,
	)
	return err
}
