package usecase

import (
	"github.com/sergiolucena1/kafkaGolan/internal/entity"
)

type ListProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type listProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductUseCase(productRepository entity.ProductRepository) *listProductsUseCase {
	return &listProductsUseCase{ProductRepository: productRepository}
}

func (u *listProductsUseCase) Execute() ([]*ListProductOutputDto, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var productsOutPut []*ListProductOutputDto
	for _, product := range products {
		productsOutPut = append(productsOutPut, &ListProductOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsOutPut, nil
}
