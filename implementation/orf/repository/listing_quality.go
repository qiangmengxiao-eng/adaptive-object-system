package repository

// ListingQuality represents Amazon listing evaluation.
type ListingQuality struct {
	Object string `yaml:"object" json:"object"`

	TitleScore float64 `yaml:"title_score" json:"title_score"`

	KeywordScore float64 `yaml:"keyword_score" json:"keyword_score"`

	BulletScore float64 `yaml:"bullet_score" json:"bullet_score"`

	DescriptionScore float64 `yaml:"description_score" json:"description_score"`

	OverallScore float64 `yaml:"overall_score" json:"overall_score"`

	Assessment string `yaml:"assessment" json:"assessment"`

	Strategy string `yaml:"strategy" json:"strategy"`
}

// ListingQualityEngine evaluates listings.
type ListingQualityEngine struct{}

// NewListingQualityEngine creates quality engine.
func NewListingQualityEngine() *ListingQualityEngine {

	return &ListingQualityEngine{}
}

// Analyze evaluates listing.
func (q *ListingQualityEngine) Analyze(
	listing AmazonListing,
) ListingQuality {

	titleScore :=
		60.0

	if len(listing.Title) >= 50 &&
		len(listing.Title) <= 200 {

		titleScore = 90
	}

	keywordScore :=
		50.0

	if len(listing.Keywords) >= 3 {

		keywordScore = 90
	}

	bulletScore :=
		float64(
			len(listing.BulletPoints),
		) * 20

	if bulletScore > 100 {

		bulletScore = 100
	}

	descriptionScore :=
		60.0

	if len(listing.Description) > 80 {

		descriptionScore = 90
	}

	overall :=
		(titleScore +
			keywordScore +
			bulletScore +
			descriptionScore) / 4

	assessment :=
		"needs_improvement"

	if overall >= 80 {

		assessment =
			"good"
	}

	if overall >= 90 {

		assessment =
			"excellent"
	}

	return ListingQuality{

		Object: listing.Object,

		TitleScore: titleScore,

		KeywordScore: keywordScore,

		BulletScore: bulletScore,

		DescriptionScore: descriptionScore,

		OverallScore: overall,

		Assessment: assessment,

		Strategy: listing.Strategy,
	}
}
