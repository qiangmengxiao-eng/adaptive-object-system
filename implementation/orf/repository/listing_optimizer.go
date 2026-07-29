package repository

type ListingOptimizer struct {
	Engine *ListingEngine

	Quality *ListingQualityEngine

	Knowledge *KnowledgeEngine
}

func NewListingOptimizer(
	engine *ListingEngine,
	quality *ListingQualityEngine,
	knowledge *KnowledgeEngine,
) *ListingOptimizer {

	return &ListingOptimizer{

		Engine: engine,

		Quality: quality,

		Knowledge: knowledge,
	}
}

func (o *ListingOptimizer) Optimize(
	listing AmazonListing,
) (*AmazonListing, ListingQuality, error) {

	result, err :=
		o.Engine.Generate(
			listing.Object,
			listing.Product,
			listing.Features,
			listing.Keywords,
		)

	if err != nil {

		return nil,
			ListingQuality{},
			err
	}

	quality :=
		o.Quality.Analyze(
			*result,
		)

	return result, quality, nil
}
