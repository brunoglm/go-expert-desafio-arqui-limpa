package usecase

import "project/internal/entity"

type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type OrderListOutputDTO struct {
	Orders []*OrderOutput
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute() (*OrderListOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	orderOutputList := []*OrderOutput{}
	for _, order := range orders {
		orderOutput := OrderOutput{
			ID: order.ID,
			Price: order.Price,
			Tax: order.Tax,
			FinalPrice: order.FinalPrice,
		}
		orderOutputList = append(orderOutputList, &orderOutput)
	}
	
	return &OrderListOutputDTO{
		Orders: orderOutputList,
	}, nil
}
