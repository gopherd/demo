package entity

type Entity interface {
	Unmarshal([]byte) error
}

type Container interface {
	GetOrCreate(id int64) Entity
}
