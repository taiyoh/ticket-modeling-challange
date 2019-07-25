package domain

var (
	specifiedTargets = []PriceSegmentType{}
	wholeTargets     = []PriceSegmentType{}
)

func init() {
	for i := 0; i < int(terminatedSegmentType); i++ {
		typ := PriceSegmentType(i)
		if typ.target() == targetSpecified {
			specifiedTargets = append(specifiedTargets, typ)
		} else {
			wholeTargets = append(wholeTargets, typ)
		}
	}
}
