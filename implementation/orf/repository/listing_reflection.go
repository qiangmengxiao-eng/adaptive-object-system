package repository

type ListingReflection struct {
	Quality ListingQuality

	Recommendation string
}

func NewListingReflection(
	q ListingQuality,
) ListingReflection {

	recommendation :=
		"improve_keywords"

	if q.OverallScore >= 90 {

		recommendation =
			"keep_strategy"
	}

	return ListingReflection{

		Quality: q,

		Recommendation: recommendation,
	}
}
