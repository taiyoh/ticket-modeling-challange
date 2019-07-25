package domain_test

import (
	"testing"
	"time"

	"github.com/taiyoh/ticket-modeling-challange/domain"
)

func TestAudienceGroup(t *testing.T) {
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
		{
			"holiday: daytime, 5, 1 younger disability, 1 elementary, use parking",
			time.Date(2019, time.Month(7), 15, 13, 0, 0, 0, ja),
			true,
			5,
			[]domain.PriceSegmentType{
				domain.DisabilityYoungerSegmentType,
				domain.ElementalySchoolStudentSegmentType,
				domain.ParkingSegmentType,
			},
			(900 * 2) + 1000 + 1400 + 1400, // 5600 ?
		},
	} {
		t.Run(tt.label, func(t *testing.T) {
			tw := domain.NewTimeWindow(tt.time, tt.holiday)
			grp := domain.NewAudienceGroup(tt.num, tt.types)
			if am := grp.Amount(tw); am != tt.amount {
				t.Errorf("wrong amouth calculated: %v", am)
			}
		})
	}

}
