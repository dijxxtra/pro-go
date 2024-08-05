package main

import "fmt"

type ChannelMessage struct {
	Category      string
	Total         float64
	CategoryError error
}

// type CategoryError struct {
// 	requestedCategory string
// }

// func (c *CategoryError) Error() string {
// 	return "Category " + c.requestedCategory + "does not exits"
// }

func (slice ProductSlice) TotalPrice(category string) (total float64, err error) {
	productCount := 0
	for _, p := range slice {
		if category == p.Category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		err = fmt.Errorf("cannot find category: %v", category)
	}
	return
}

func (p *ProductSlice) TotalPriceAsync(categories []string, channel chan<- ChannelMessage) {
	for _, cat := range categories {
		total, err := p.TotalPrice(cat)
		channel <- ChannelMessage{
			Category:      cat,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
