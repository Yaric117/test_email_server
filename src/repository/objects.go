package repository

import (
	"context"
	"fmt"
	"testproject/db"
)

const loggerObjectsTitle = "[repository][object]"

type Object struct {
	Id       string `db:"id" json:"id,omitempty"`
	Name     string `db:"name" json:"name,omitempty"`
	Timezone int    `db:"timezone" json:"timezone,omitempty"`
}

func (o *Object) Insert(ctx context.Context) (string, error) {
	logSectionTitle := loggerObjectsTitle + "[Insert]"

	sql := `INSERT INTO objects (
				name,
				timezone
			) 
			VALUES ($1,$2)
			RETURNING id`

	var id string
	err := db.Conn.QueryRow(ctx, sql,
		o.Name,
		o.Timezone,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("%s: %s\n", logSectionTitle, err)
	}
	return id, nil

}

func DeleteAllObjects(ctx context.Context) error {
	logSectionTitle := loggerObjectsTitle + "[DeleteAllObjects]"
	_, err := db.Conn.Exec(ctx, `DELETE FROM objects`)

	if err != nil {
		return fmt.Errorf("%s %s\n", logSectionTitle, err)
	}
	return nil
}
