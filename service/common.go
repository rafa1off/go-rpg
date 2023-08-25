package service

func pagination(page, limit *int32) (newLimit, skip int) {
	p, l := 1, 5

	if page != nil {
		p = int(*page)
	}
	if limit != nil {
		l = int(*limit)
	}

	return l, l * (p - 1)
}
