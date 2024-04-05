package models

type Car struct {
	ID     int64  `db:"id" json:"-"`
	RegNum string `db:"reg_num" json:"regNum" validate:"required"`
	Mark   string `db:"mark" json:"mark" validate:"required"`
	Model  string `db:"model" json:"model" validate:"required"`
	Year   int    `db:"year" json:"year"`
	Owner  Human  `db:"-" json:"owner" validate:"required"`
}
