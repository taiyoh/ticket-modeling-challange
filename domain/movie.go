package domain

// MovieSpecialty represents movie's attract with price change.
type MovieSpecialty int

const (
	// ThreeDeeSpecialty represents this movie is 3D.
	ThreeDeeSpecialty MovieSpecialty = iota + 1
	// PremiumLoudnessSpecialty means '極上爆音上映' in Japanese.
	PremiumLoudnessSpecialty
)

// Movie represents MovieSpetialty aggregation.
type Movie struct {
	specialties []MovieSpecialty
}
