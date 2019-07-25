package domain

// Price represents calculated price by segment.
type Price int

// Add returns result of object Price + argument Price.
func (p Price) Add(pr Price) Price {
	return p + pr
}

// PriceSegmentType provides segment type for price discount.
type PriceSegmentType int

const (
	// NormalSegmentType represents segment for no licence and identification.
	NormalSegmentType PriceSegmentType = iota
	// DisabilityElderSegmentType represents segment for disability person who is over 18 years old.
	DisabilityElderSegmentType
	// DisabilityYoungerSegmentType represents segment for disability person who is under 18 years old.
	DisabilityYoungerSegmentType
	// ElementalySchoolStudentSegmentType represents segment for elementary school student and child who is over 3 years old.
	ElementalySchoolStudentSegmentType
	// HighSchoolStudentSegmentType represents segment for high school and junior high school student who have student ID card.
	HighSchoolStudentSegmentType
	// CollageStudentSegmentType represents segment for collage student who have student ID card.
	CollageStudentSegmentType
	// SeniorSegmentType represents segment for elder who is over 70 and have identification.
	SeniorSegmentType
	// CinemaCitizenSegmentType represents segment for CinemaCitizen who is under 60 years old.
	CinemaCitizenSegmentType
	// CinemaCitizenElderSegmentType represents segment for CinemaCitizen who is over 60 years old.
	CinemaCitizenElderSegmentType
	// MICardSegmentType represents segment for MICard customer.
	MICardSegmentType
	// ParkingSegmentType represents segment for '駐車場パーク80' customer.
	ParkingSegmentType

	terminatedSegmentType
)

func (t PriceSegmentType) count() int {
	switch t {
	case DisabilityElderSegmentType, DisabilityYoungerSegmentType:
		return 2
	case NormalSegmentType, ParkingSegmentType:
		return 0
	default:
		return 1
	}
}

type segmentTarget int

const (
	targetSpecified segmentTarget = iota + 1
	targetWhole
)

func (t PriceSegmentType) target() segmentTarget {
	switch t {
	case NormalSegmentType, ParkingSegmentType:
		return targetWhole
	default:
		return targetSpecified
	}
}

// PriceSegment provides price calculator by segment.
type PriceSegment struct {
	typ PriceSegmentType
	wd  Price
	wn  Price
	hd  Price
	hn  Price
	md  Price
}

func (s PriceSegment) find(typ TimeWindowType) Price {
	switch typ {
	case MovieDay:
		return s.md
	case HolidayDayTime:
		return s.hd
	case HolidayNightTime:
		return s.hn
	case WeekDayNightTime:
		return s.wn
	default:
		return s.wd
	}
}

// Calculate provides custom calculation.
func (s PriceSegment) Calculate(tw TimeWindow) Price {
	typ := tw.Type()
	if s.typ == CinemaCitizenSegmentType && typ == MovieDay && tw.dayType() == weekday {
		typ = WeekDayDayTime
	}
	return s.find(typ)
}
