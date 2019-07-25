package domain

// Customer represents person who comes to watch the movie.
type Customer struct {
	segmentType PriceSegmentType
}

// CustomerGroup represents target persons for each billing.
type CustomerGroup []Customer

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

// NewCustomerGroup returns CustomerGroup object with segment type attached.
func NewCustomerGroup(num int, types []PriceSegmentType) CustomerGroup {
	customers := make(CustomerGroup, 0, num)
	counts := newSegTypeCount(types)
	for _, pair := range counts.detectSpecifieds() {
		count := pair.totalCount()
		for i := 0; i < count; i++ {
			customers = customers.Add(pair.typ)
			if num--; num < 1 {
				return customers
			}
		}
	}

	typ := counts.detectWhole()
	for i := 0; i < num; i++ {
		customers = customers.Add(typ)
	}

	return customers
}

// Add provides new CustomerGroup with new Customer.
func (g CustomerGroup) Add(t PriceSegmentType) CustomerGroup {
	return append(g, Customer{t})
}

// Amount provides calculating total price of this group.
func (g CustomerGroup) Amount(tw TimeWindow) Price {
	p := Price(0)
	for _, th := range g {
		seg := NewPriceSegment(th.segmentType)
		p = p.Add(seg.Calculate(tw))
	}
	return p
}
