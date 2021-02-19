package server

import (
	"context"

	protos "../protos/currency"
	"github.com/hashicorp/go-hclog"
)

// Currency ...
type Currency struct {
	rates *data.ExchangeRates
	log   hclog.Logger
}

// NewCurrency ...
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{r, l}
}

// GetRate ...
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}
