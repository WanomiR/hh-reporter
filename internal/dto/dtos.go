package dto

import "time"

type User struct {
	TelegramId     int `db:"telegram_id,int"`
	TelegramChatId int `db:"telegram_chat_id,int"`
}

type Vacancy struct {
	VacancyId    int       `db:"vacancy_id,int"`
	VacancyName  int       `db:"vacancy_name"`
	RoleId       int       `db:"role_id,int"`
	Experience   string    `db:"experience"`
	Url          string    `db:"url"`
	PublisthedAt time.Time `db:"published_at"`
}

type Role struct {
	RoleId   int    `db:"role_id,int"`
	RoleName string `db:"role_name"`
}
