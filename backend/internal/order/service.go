package order

type service struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &service{repo: repo}
}

func (s *service) CreateOrder() {

}
