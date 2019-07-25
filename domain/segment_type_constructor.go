package domain

func newNormalSegment() PriceSegment {
	return PriceSegment{NormalSegmentType, 1800, 1300, 1800, 1300, 1100}
}

func newCinemaCitizenSegment() PriceSegment {
	return PriceSegment{CinemaCitizenSegmentType, 1000, 1000, 1300, 1000, 1000}
}

func newCinemaCitizenElderSegment() PriceSegment {
	return PriceSegment{CinemaCitizenElderSegmentType, 1000, 1000, 1000, 1000, 1000}
}

func newSeniorSegment() PriceSegment {
	return PriceSegment{SeniorSegmentType, 1100, 1100, 1100, 1100, 1100}
}

func newCollageStudentSegment() PriceSegment {
	return PriceSegment{CollageStudentSegmentType, 1500, 1300, 1500, 1300, 1100}
}

func newHighSchoolStudentSegment() PriceSegment {
	return PriceSegment{HighSchoolStudentSegmentType, 1500, 1300, 1500, 1300, 1100}
}

func newElementalySchoolStudentSegment() PriceSegment {
	return PriceSegment{ElementalySchoolStudentSegmentType, 1500, 1300, 1500, 1300, 1100}
}

func newDisabilityElderSegment() PriceSegment {
	return PriceSegment{DisabilityElderSegmentType, 1000, 1000, 1000, 1000, 1000}
}

func newDisabilityYoungerSegment() PriceSegment {
	return PriceSegment{DisabilityYoungerSegmentType, 900, 900, 900, 900, 900}
}

func newMICardSegment() PriceSegment {
	return PriceSegment{MICardSegmentType, 1600, 1300, 1600, 1300, 1100}
}

func newParkingSegment() PriceSegment {
	return PriceSegment{ParkingSegmentType, 1400, 1100, 1400, 1100, 1100}
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
