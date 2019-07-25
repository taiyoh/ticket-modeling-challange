package domain_test

import (
	"testing"
	"time"

	"github.com/taiyoh/movie-ticket/domain"
)

func TestCustomerGroup(t *testing.T) {
	ja, _ := time.LoadLocation("Asia/Tokyo")

	for _, tt := range []struct {
		label   string
		time    time.Time
		holiday bool
		num     int
		types   []domain.PriceSegmentType
		amount  domain.Price
	}{
		{
			"weekday: daytime, 3, all normal",
			time.Date(2019, time.Month(7), 25, 15, 0, 0, 0, ja),
			false,
			3,
			[]domain.PriceSegmentType{},
			1800 * 3,
		},
	} {
		t.Run(tt.label, func(t *testing.T) {
			tw := domain.NewTimeWindow(tt.time, tt.holiday)
			grp := domain.NewCustomerGroup(tt.num, tt.types)
			if am := grp.Amount(tw); am != tt.amount {
				t.Errorf("wrong amouth calculated: %v", am)
			}
		})
	}

}
