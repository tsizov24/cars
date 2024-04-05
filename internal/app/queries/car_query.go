package queries

import (
	"cars/internal/app/models"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func CreateCar(dbConn *pgx.Conn, car *models.Car) error {
	if err := createHuman(dbConn, &car.Owner); err != nil {
		return err
	}

	sql, args, err := sq.
		Insert("cars").
		Columns("reg_num", "mark", "model", "year", "owner_id").
		Values(car.RegNum, car.Mark, car.Model, car.Year, car.Owner.ID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = dbConn.Exec(context.Background(), sql, args...)
	return err
}

func DeleteCar(dbConn *pgx.Conn, regNum string) error {
	sql, args, err := sq.
		Delete("cars").
		Where(sq.Eq{"reg_num": regNum}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = dbConn.Exec(context.Background(), sql, args...)

	return err
}

func GetCar(dbConn *pgx.Conn, regNum string) (*models.Car, error) {
	c := &models.Car{RegNum: regNum}

	sql, args, err := sq.
		Select("c.mark", "c.model", "c.year", "p.name", "p.surname", "p.patronymic").
		From("cars c").
		LeftJoin("people p ON c.owner_id = p.id").
		Where(sq.Eq{"reg_num": regNum}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = dbConn.QueryRow(context.Background(), sql, args...).
		Scan(&c.Mark, &c.Model, &c.Year, &c.Owner.Name, &c.Owner.Surname, &c.Owner.Patronymic)

	if err == pgx.ErrNoRows {
		logrus.Info(err)
		c.RegNum = ""
		return c, nil
	}

	return c, err
}

func GetCars(dbConn *pgx.Conn, limit, offset int) ([]models.Car, error) {
	cars := make([]models.Car, 0, limit)

	sql, args, err := sq.
		Select("c.reg_num", "c.mark", "c.model", "c.year", "p.name", "p.surname", "p.patronymic").
		From("cars c").
		LeftJoin("people p ON c.owner_id = p.id").
		OrderBy("c.created_at").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := dbConn.Query(context.Background(), sql, args...)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := models.Car{}
		err = rows.Scan(&c.RegNum, &c.Mark, &c.Model, &c.Year, &c.Owner.Name, &c.Owner.Surname, &c.Owner.Patronymic)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}

	return cars, nil
}

func IsCarExists(dbConn *pgx.Conn, regNum string) (bool, error) {
	sql, args, err := sq.
		Select("reg_num").
		From("cars").
		Where(sq.Eq{"reg_num": regNum}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return false, err
	}

	row := dbConn.QueryRow(context.Background(), sql, args...)
	err = row.Scan(&regNum)
	if err == pgx.ErrNoRows {
		logrus.Info(err)
		return false, nil
	}
	return err == nil, err
}

func UpdateCar(dbConn *pgx.Conn, car *models.Car) error {
	if err := updateCarOwner(dbConn, car); err != nil {
		return err
	}

	sql, args, err := sq.
		Update("cars").
		Set("mark", car.Mark).
		Set("model", car.Model).
		Set("year", car.Year).
		Where(sq.Eq{"reg_num": car.RegNum}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = dbConn.Exec(context.Background(), sql, args...)
	return err
}
