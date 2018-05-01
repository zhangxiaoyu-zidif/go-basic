package model

import "errors"

var (
	ErrStockNotEnough = errors.New("Stock is not enough.")
)

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

func CreateBook(name string, total int, author string, createTime string) (b *Book) {
	b = &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createTime,
	}
	return b
}

func (b *Book) CanBorrow(c int) bool {
	return (b.Total >= c)
}

func (b *Book) Borrow(c int) (err error) {
	if b.CanBorrow(c) {
		b.Total -= c
		err = nil
		return
	} else {
		err = ErrStockNotEnough
		return
	}
}

func (b *Book) Back(c int) error {
	b.Total += c
	return nil
}
