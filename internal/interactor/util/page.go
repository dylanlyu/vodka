package util

func Pagination(totalRecord, limit int64) int64 {
	if totalRecord <= 0 {
		return 0
	}

	return (totalRecord + limit - 1) / limit
}
