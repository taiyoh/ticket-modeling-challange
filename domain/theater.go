package domain

// Theater represents person who comes to watch the movie.
type Theater struct {
	segmentType PriceSegmentType
}

// TheaterGroup represents target persons for each billing.
type TheaterGroup []Theater

type segTypeCount map[PriceSegmentType]int

var highPriorities = []PriceSegmentType{
	DisabilityYoungerSegmentType,
	DisabilityElderSegmentType,
	ElementalySchoolStudentSegmentType,
	HighSchoolStudentSegmentType,
	CinemaCitizenSegmentType,
	CinemaCitizenElderSegmentType,
	SeniorSegmentType,
	CollageStudentSegmentType,
	MICardSegmentType,
}

type segTypeCountPair struct {
	typ   PriceSegmentType
	count int
}

func (p segTypeCountPair) totalCount() int {
	return p.typ.targetCount() * p.count
}

func (p segTypeCountPair) newTheater() Theater {
	return Theater{p.typ}
}

func (s segTypeCount) highPriorities() []segTypeCountPair {
	pairs := []segTypeCountPair{}
	for _, t := range highPriorities {
		if c, exists := s[t]; exists {
			pairs = append(pairs, segTypeCountPair{t, c})
		}
	}
	return pairs
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
	theaters := make(TheaterGroup, 0, num)
	counts := newSegTypeCount(types)
	for _, pair := range counts.highPriorities() {
		count := pair.totalCount()
		for i := 0; i < count; i++ {
			theaters = append(theaters, pair.newTheater())
			if num--; num < 1 {
				return theaters
			}
		}
	}

	typ := counts.normalPriority()
	for i := 0; i < num; i++ {
		theaters = append(theaters, Theater{typ})
	}

	return theaters
}
