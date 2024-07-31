package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

var categoryMaxPrices = map[string]float64{
	"Watersport": 250,
	"Soccer":     150,
	"Chess":      50,
}

func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
}

type taxRate struct {
	rate, threshold float64
}

func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}
	return &taxRate{rate, threshold}
}

func (t *taxRate) calcTax(product *Product) (price float64) {
	if product.price > t.threshold {
		price = product.price + (product.price * t.rate)
	} else {
		price = product.price
	}
	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}
	return
}
