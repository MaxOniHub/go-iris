package data_mappers

type Limit struct {
	offset int64
	limit  int64
}

func NewLimit(offset, limit int64) *Limit {
	l := new(Limit)
	l.offset = offset
	l.limit = limit
	return l
}

func (l *Limit) GetOffset() int64 {
	return l.offset
}

func (l *Limit) GetLimit() int64 {
	return l.limit
}

func (l *Limit) SetOffset(offset int64) {
	l.offset = offset
}

func (l *Limit) SetLimit(limit int64) {
	l.limit = limit
}
