package models

type Human struct {
	ID         int64  `db:"id" json:"-"`
	Name       string `db:"name" json:"name" validate:"required"`
	Surname    string `db:"surname" json:"surname" validate:"required"`
	Patronymic string `db:"patronymic" json:"patronymic"`
}
