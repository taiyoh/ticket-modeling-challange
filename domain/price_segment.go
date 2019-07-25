package domain

// Price represents calculated price by segment.
type Price int

// PriceSegmentType provides segment type for price discount.
type PriceSegmentType int

const (
	// NormalSegmentType represents segment for no licence and identification.
	NormalSegmentType PriceSegmentType = iota
	// CinemaCitizenSegmentType represents segment for CinemaCitizen who is under 60 years old.
	CinemaCitizenSegmentType
	// CinemaCitizenElderSegmentType represents segment for CinemaCitizen who is over 60 years old.
	CinemaCitizenElderSegmentType
	// SeniorSegmentType represents segment for elder who is over 70 and have identification.
	SeniorSegmentType
	// CollageStudentSegmentType represents segment for collage student who have student ID card.
	CollageStudentSegmentType
	// HighSchoolStudentSegmentType represents segment for high school and junior high school student who have student ID card.
	HighSchoolStudentSegmentType
	// ElementalySchoolStudentSegmentType represents segment for elementary school student and child who is over 3 years old.
	ElementalySchoolStudentSegmentType
	// DisabilityElderSegmentType represents segment for disability person who is over 18 years old.
	DisabilityElderSegmentType
	// DisabilityYoungerSegmentType represents segment for disability person who is under 18 years old.
	DisabilityYoungerSegmentType
	// MICardSegmentType represents segment for MICard customer.
	MICardSegmentType
	// ParkingSegmentType represents segment for '駐車場パーク80' customer.
	ParkingSegmentType
)

func (t PriceSegmentType) targetCount() int {
	if t == DisabilityElderSegmentType || t == DisabilityYoungerSegmentType {
		return 2
	}
	return 1
}

type priceSegmentBase struct {
	wd Price
	wn Price
	hd Price
	hn Price
	md Price
}

func (s priceSegmentBase) find(typ TimeWindowType) Price {
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

func (s priceSegmentBase) Calculate(tw TimeWindow) Price {
	return s.find(tw.Type())
}

// PriceSegment provides price calculator by segment.
type PriceSegment interface {
	Calculate(TimeWindow) Price
}

// NormalSegment provides Normal segment calculator.
type NormalSegment struct {
	priceSegmentBase
}

func newNormalSegment() PriceSegment {
	seg := priceSegmentBase{1800, 1300, 1800, 1300, 1100}
	return NormalSegment{seg}
}

// CinemaCitizenSegment provides CinemaCitizen segment calculator.
type CinemaCitizenSegment struct {
	priceSegmentBase
}

func newCinemaCitizenSegment() PriceSegment {
	seg := priceSegmentBase{1000, 1000, 1300, 1000, 1000}
	return CinemaCitizenSegment{seg}
}

// CinemaCitizenElderSegment provides CinemaCitizenElder segment calculator.
type CinemaCitizenElderSegment struct {
	priceSegmentBase
}

func newCinemaCitizenElderSegment() PriceSegment {
	seg := priceSegmentBase{1000, 1000, 1000, 1000, 1000}
	return CinemaCitizenElderSegment{seg}
}

// SeniorSegment provides Senior segment calculator.
type SeniorSegment struct {
	priceSegmentBase
}

func newSeniorSegment() PriceSegment {
	seg := priceSegmentBase{1100, 1100, 1100, 1100, 1100}
	return SeniorSegment{seg}
}

// CollageStudentSegment provides CollageStudent segment calculator.
type CollageStudentSegment struct {
	priceSegmentBase
}

// Calculate provides custom calculation for CollageStudentSegment.
func (s CollageStudentSegment) Calculate(tw TimeWindow) Price {
	typ := tw.Type()
	if typ == MovieDay && tw.dayType() == weekday {
		return s.wd
	}
	return s.find(typ)
}

func newCollageStudentSegment() PriceSegment {
	seg := priceSegmentBase{1500, 1300, 1500, 1300, 1100}
	return CollageStudentSegment{seg}
}

// HighSchoolStudentSegment provides HighSchoolStudent segment calculator.
type HighSchoolStudentSegment struct {
	priceSegmentBase
}

func newHighSchoolStudentSegment() PriceSegment {
	seg := priceSegmentBase{1500, 1300, 1500, 1300, 1100}
	return HighSchoolStudentSegment{seg}
}

// ElementalySchoolStudentSegment provides ElementalySchoolStudent segment calculator.
type ElementalySchoolStudentSegment struct {
	priceSegmentBase
}

func newElementalySchoolStudentSegment() PriceSegment {
	seg := priceSegmentBase{1500, 1300, 1500, 1300, 1100}
	return ElementalySchoolStudentSegment{seg}
}

// DisabilityElderSegment provides DisabilityElder segment calculator.
type DisabilityElderSegment struct {
	priceSegmentBase
}

func newDisabilityElderSegment() PriceSegment {
	seg := priceSegmentBase{1000, 1000, 1000, 1000, 1000}
	return DisabilityElderSegment{seg}
}

// DisabilityYoungerSegment provides DisabilityYounger segment calculator.
type DisabilityYoungerSegment struct {
	priceSegmentBase
}

func newDisabilityYoungerSegment() PriceSegment {
	seg := priceSegmentBase{900, 900, 900, 900, 900}
	return DisabilityYoungerSegment{seg}
}

// MICardSegment provides MICard segment caluclator.
type MICardSegment struct {
	priceSegmentBase
}

func newMICardSegment() PriceSegment {
	seg := priceSegmentBase{1600, 1300, 1600, 1300, 1100}
	return MICardSegment{seg}
}

// ParkingSegment provides Parking segment calculator.
type ParkingSegment struct {
	priceSegmentBase
}

func newParkingSegment() PriceSegment {
	seg := priceSegmentBase{1400, 1100, 1400, 1100, 1100}
	return ParkingSegment{seg}
}

var segmentsByType = map[PriceSegmentType]func() PriceSegment{
	CinemaCitizenSegmentType:           newCinemaCitizenSegment,
	CinemaCitizenElderSegmentType:      newCinemaCitizenElderSegment,
	SeniorSegmentType:                  newSeniorSegment,
	CollageStudentSegmentType:          newCollageStudentSegment,
	HighSchoolStudentSegmentType:       newHighSchoolStudentSegment,
	ElementalySchoolStudentSegmentType: newElementalySchoolStudentSegment,
	DisabilityElderSegmentType:         newDisabilityElderSegment,
	DisabilityYoungerSegmentType:       newDisabilityYoungerSegment,
	MICardSegmentType:                  newMICardSegment,
	ParkingSegmentType:                 newParkingSegment,
}

// NewPriceSegment returns PriceSegment object.
func NewPriceSegment(typ PriceSegmentType) PriceSegment {
	if fn, exists := segmentsByType[typ]; exists {
		return fn()
	}
	return newNormalSegment()
}
