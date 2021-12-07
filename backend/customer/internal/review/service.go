package review

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type service struct {
	repo ReviewRepository
}

var (
	ErrNotFoundOrders     = errors.New("not found orders")
	ErrUpdateReviewFailed = errors.New("update review failed")
	ErrDataReviewBad      = errors.New("data review bad")
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
		var star int64
		p := strings.LastIndex(order.Review, ", ")
		if p != -1 {
			star, err = strconv.ParseInt(order.Review[p+2:], 10, 32)
			if err != nil {
				star = 0

			}
			response = append(response, &Review{
				No:     i,
				Review: order.Review[:p],
				Star:   int(star),
			})
		}
	}

	return response, nil
}

func (s *service) GetOrdersDoneCustomer(ctx context.Context, custID int) ([]*Order, error) {
	orders, err := s.repo.GetOrdersDoneCustomer(ctx, custID)
	if err != nil {
		return nil, ErrNotFoundOrders
	}

	return orders, nil
}

func (s *service) UpdateOrderReview(ctx context.Context, orderID, custID int, review string, star int) error {
	if review == "" || star < 0 || star > 5 {
		return ErrDataReviewBad
	}

	review = fmt.Sprintf("%s, %d", review, star)

	orders, err := s.repo.GetOrdersDoneCustomer(ctx, custID)
	if err != nil {
		return ErrNotFoundOrders
	}

	for _, order := range orders {
		if order.ID == orderID {
			err := s.repo.UpdateOrderReview(ctx, orderID, review)
			if err != nil {
				return ErrUpdateReviewFailed
			}
			return nil
		}
	}

	return ErrNotFoundOrders
}
