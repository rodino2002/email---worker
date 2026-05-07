package db

import (
	"context"
	"time"
)

func CountProcessedMessages() (int, error) {
	var count int

	query := `
		SELECT COUNT(*)
		FROM receivers
		WHERE status IN ('DELIVRD', 'ESME_ROK', 'REJECTD')
		  AND updated_at >= NOW() - INTERVAL '5 minutes'
	`

	// query := `
	// 	SELECT
	// 	updated_at AS ultima_data
	// 	FROM receivers
	// 	WHERE status IN ( 'DELIVRD','ESME_ROK', 'REJECTD' )
	// 	ORDER BY updated_at DESC LIMIT 1
	// `

	err := DB.QueryRow(
		context.Background(),
		query,
	).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

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
