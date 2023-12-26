// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: alerts.sql

package database

import (
	"context"
)

const updateAlertStatus = `-- name: UpdateAlertStatus :exec
UPDATE "Alerts" SET
  status = $2
WHERE "id" = $1
`

type UpdateAlertStatusParams struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateAlertStatus(ctx context.Context, arg UpdateAlertStatusParams) error {
	_, err := q.db.Exec(ctx, updateAlertStatus, arg.ID, arg.Status)
	return err
}