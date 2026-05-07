package db

import (
	"context"
	"time"
)

func CountPendingMessages() (int, error) {
	var count int
	now := time.Now()

	query := `
		SELECT COUNT(*)
		FROM receivers
		WHERE status = 'PENDING' AND created_at between $1 and $2
	`

	// início do dia (00:00:00)
	startOfDay := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	// fim do dia (23:59:59.999...)
	endOfDay := startOfDay.Add(24*time.Hour - time.Nanosecond)
	err := DB.QueryRow(
		context.Background(),
		query,
		startOfDay,
		endOfDay,
	).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
