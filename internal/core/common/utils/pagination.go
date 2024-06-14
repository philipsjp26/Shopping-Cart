package utils

func Pagination(limit, page int) (int, int) {
	offset := (page - 1) * limit
	return limit, offset
}
