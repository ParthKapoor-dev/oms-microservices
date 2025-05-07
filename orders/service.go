package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/parthkapoor-dev/common/api"
)

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {

	fmt.Println("Validating Orders")

	if len(p.Items) == 0 {
		return errors.New("Atleast a single items is required to place a order")
	}

	mergedItems := make([]*pb.ItemsWithQuantity, 0)

	mappedItems := make(map[string]int32)

	for _, i := range p.Items {
		if mappedItems[i.ID] == 0 {
			mappedItems[i.ID] = i.Quantity
		} else {
			mappedItems[i.ID] += i.Quantity
		}
	}

	for key := range mappedItems {
		mergedItems = append(mergedItems, &pb.ItemsWithQuantity{
			ID:       key,
			Quantity: mappedItems[key],
		})
	}

	log.Println("Merged Items are: ", mergedItems)

	return nil
}
