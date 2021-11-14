package order

type service struct {
	mq OrderMessageQueue
}

func NewOrderService(mq OrderMessageQueue) OrderService {
	return &service{mq: mq}
}
