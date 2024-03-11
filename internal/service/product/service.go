package product

import (
	"fmt"
)

type Service struct {
	p []*Product
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get(id int) (*Product, error) {
	if id >= len(s.p) {
		return nil, fmt.Errorf("неверный идентификатор")
	}

	return s.p[id], nil
}

func (s *Service) Create(product *Product) uint64 {
	s.p = append(s.p, product)
	return uint64(len(s.p) - 1)
}

func (s *Service) List(cursor int, limit int) ([]*Product, error) {
	if limit >= len(s.p) {
		limit = len(s.p)
	}
	if cursor > limit || cursor < 0 {
		return nil, fmt.Errorf("неверные промежуток списка товаров")
	}
	return s.p[cursor:limit], nil
}

func (s *Service) Remove(id int) {
	newProducts := make([]*Product, len(s.p)-1)
	for i := range s.p {
		if i == id {
			continue
		}
		newProducts[i] = s.p[i]
	}
	s.p = newProducts
}

func (s *Service) Update(id int, product *Product) error {
	if id >= len(s.p) {
		return fmt.Errorf("неверный идентификатор")
	}
	s.p[id] = product
	return nil
}
