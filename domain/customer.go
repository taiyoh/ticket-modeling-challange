package domain

// Audience represents person who comes to watch the movie.
type Audience struct {
	segmentType PriceSegmentType
}

// AudienceGroup represents target persons for each billing.
type AudienceGroup []Audience

type segTypeCount map[PriceSegmentType]int

type segTypeCountPair struct {
	typ   PriceSegmentType
	count int
}

func (p segTypeCountPair) totalCount() int {
	return p.typ.count() * p.count
}

func (s segTypeCount) detectSpecifieds() []segTypeCountPair {
	pairs := []segTypeCountPair{}
	for _, t := range specifiedTargets {
		if c, exists := s[t]; exists {
			pairs = append(pairs, segTypeCountPair{t, c})
		}
	}
	return pairs
}

func (s segTypeCount) detectWhole() PriceSegmentType {
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

// NewAudienceGroup returns AudienceGroup object with segment type attached.
func NewAudienceGroup(num int, types []PriceSegmentType) AudienceGroup {
	audiences := make(AudienceGroup, 0, num)
	counts := newSegTypeCount(types)
	for _, pair := range counts.detectSpecifieds() {
		count := pair.totalCount()
		for i := 0; i < count; i++ {
			audiences = audiences.Add(pair.typ)
			if num--; num < 1 {
				return audiences
			}
		}
	}

	typ := counts.detectWhole()
	for i := 0; i < num; i++ {
		audiences = audiences.Add(typ)
	}

	return audiences
}

// Add provides new AudienceGroup with new Audience.
func (g AudienceGroup) Add(t PriceSegmentType) AudienceGroup {
	return append(g, Audience{t})
}

// Amount provides calculating total price of this group.
func (g AudienceGroup) Amount(tw TimeWindow) Price {
	p := Price(0)
	for _, th := range g {
		seg := NewPriceSegment(th.segmentType)
		p = p.Add(seg.Calculate(tw))
	}
	return p
}
