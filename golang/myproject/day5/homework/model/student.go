package model

type Student struct {
	Name  string
	Grade string
	Id    string
	Sex   string
	Books []*BorrowItem
}

type BorrowItem struct {
	book *Book
	num  int
}

func CreateStudent(name string, grade, id, sex string) *Student {
	stu := &Student{
		Name:  name,
		Grade: grade,
		Id:    id,
		Sex:   sex,
	}
	return stu
}

func (s *Student) AddBook(b *BorrowItem) {
	s.Books = append(s.Books, b)
}

func (s *Student) DelBook(b *BorrowItem) (err error) {
	for i := 0; i < len(s.Books); i++ {
		if s.Books[i].book.Name == b.book.Name {
			if b.num == s.Books[i].num {
				front := s.Books[0:i]
				left := s.Books[i+1:]
				front = append(front, left...)
				s.Books = front
				return
			}
			s.Books[i].num -= b.num
			return
		}
	}
	err = ErrStockNotEnough
	return
}

func (s *Student) GetBookList() []*BorrowItem {
	return s.Books
}
