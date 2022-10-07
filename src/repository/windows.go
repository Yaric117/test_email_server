package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"testproject/db"
	"time"
)

const loggerWindowsTitle = "[repository][objectWindow]"

type ObjectWindow struct {
	Id        string    `db:"id" json:"id,omitempty"`
	ObjectId  string    `db:"object_id" json:"objectId,omitempty"`
	DateStart time.Time `db:"date_start" json:"dateStart,omitempty"`
	DateEnd   time.Time `db:"date_end" json:"dateEnd,omitempty"`
	Object    Object    `db:"object" json:"object"`
}

func (w *ObjectWindow) Insert(ctx context.Context) (string, error) {
	logSectionTitle := loggerWindowsTitle + "[Insert]"

	sql := `INSERT INTO technological_windows (
				object_id,
				date_start,
				date_end
			) 
			VALUES ($1,$2,$3)
			RETURNING id`

	var id string
	err := db.Conn.QueryRow(ctx, sql,
		w.ObjectId,
		w.DateStart,
		w.DateEnd,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("%s: %s\n", logSectionTitle, err)
	}
	return id, nil
}

func GetAllWindows(ctx context.Context) ([]ObjectWindow, error) {
	logSectionTitle := loggerWindowsTitle + "[GetAllWindows]"
	var windows []ObjectWindow

	sql := `
		SELECT
		w.id,
		w.object_id,
		w.date_start,
		w.date_end,
 		json_build_object(
 		    'id',
 		    o.id,
 		    'name',
 		    o.name,
 		    'timezone',
 		    o.timezone
 		) as object
    	FROM technological_windows w
    	LEFT JOIN objects o ON w.object_id = o.id
        `
	err := pgxscan.Select(ctx, db.Conn, &windows, sql)

	if err != nil {
		return windows, fmt.Errorf("%s %s\n", logSectionTitle, err)
	}
	return windows, nil
}
