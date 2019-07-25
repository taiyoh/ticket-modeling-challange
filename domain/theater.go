package domain

// Theater represents person who comes to watch the movie.
type Theater struct {
	segmentType PriceSegmentType
}

// TheaterGroup represents target persons for each billing.
type TheaterGroup []Theater

type segTypeCount map[PriceSegmentType]int

var highPriorities = []PriceSegmentType{
	CinemaCitizenSegmentType,
	CinemaCitizenElderSegmentType,
	SeniorSegmentType,
	CollageStudentSegmentType,
	HighSchoolStudentSegmentType,
	ElementalySchoolStudentSegmentType,
	DisabilityElderSegmentType,
	DisabilityYoungerSegmentType,
	MICardSegmentType,
}

func (s segTypeCount) highPriorities() segTypeCount {
	clones := segTypeCount{}
	for _, t := range highPriorities {
		if c, exists := s[t]; exists {
			clones[t] = c
		}
	}
	return clones
}

func (s segTypeCount) normalPriority() PriceSegmentType {
	if _, exists := s[ParkingSegmentType]; exists {
		return ParkingSegmentType
	}
	return NormalSegmentType
}

func newSegTypeCount(types []PriceSegmentType) segTypeCount {
	counts := segTypeCount{}
	for _, t := range types {
		counts[t]++
	}
	return counts
}

// NewTheaterGroup returns TheaterGroup object with segment type attached.
func NewTheaterGroup(num int, types []PriceSegmentType) TheaterGroup {
	theaters := TheaterGroup{}
	counts := newSegTypeCount(types)
	for t, c := range counts.highPriorities() {
		count := t.targetCount() * c
		for i := 0; i < count; i++ {
			theaters = append(theaters, Theater{t})
			num--
			if num < 1 {
				return theaters
			}
		}
	}

	typ := counts.normalPriority()
	for i := 0; i < num; i++ {
		theaters = append(theaters, Theater{typ})
	}

	return []Theater{}
}
