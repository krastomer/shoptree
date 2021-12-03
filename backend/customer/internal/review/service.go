package review

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

type service struct {
	repo ReviewRepository
}

var (
	ErrNotFoundOrders = errors.New("not found orders")
)

func NewReviewService(repo ReviewRepository) ReviewService {
	return &service{repo: repo}
}

func (s *service) GetReviews(ctx context.Context) ([]*Review, error) {
	orders, err := s.repo.GetOrdersDoneWithReview(ctx)
	if err != nil {
		return nil, ErrNotFoundOrders
	}

	var response []*Review
	for i, order := range orders {
		p := strings.LastIndex(order.Review, ", ")
		star, err := strconv.ParseInt(order.Review[p+2:], 10, 32)
		if err != nil {
			star = 0

		}
		response = append(response, &Review{
			No:     i,
			Review: order.Review[:p],
			Star:   int(star),
		})
	}

	return response, nil
}
