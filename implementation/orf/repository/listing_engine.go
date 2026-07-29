package repository

import "time"

// ListingEngine generates Amazon listings.
type ListingEngine struct {
	Knowledge *KnowledgeEngine

	Store *ListingStore
}

// NewListingEngine creates listing engine.
func NewListingEngine(
	k *KnowledgeEngine,
	s *ListingStore,
) *ListingEngine {

	return &ListingEngine{

		Knowledge: k,

		Store: s,
	}
}

// Generate creates Amazon listing.
func (l *ListingEngine) Generate(
	object string,
	product string,
	features []string,
	keywords []string,
) (*AmazonListing, error) {

	title :=
		product

	for _, keyword := range keywords {

		title +=
			" " +
				keyword
	}

	listing :=
		AmazonListing{

			Object: object,

			Product: product,

			Features: features,

			Title: title,

			BulletPoints: []string{

				features[0],

				features[1],

				"Premium quality design",

				"Easy to use",

				"Reliable daily performance",
			},

			Description: "Designed for customers seeking a reliable and practical solution.",

			Keywords: keywords,

			Strategy: "keyword_front_loading",

			Version: 1,

			CreatedAt: time.Now(),
		}

	if l.Store != nil {

		err :=
			l.Store.Append(
				listing,
			)

		if err != nil {

			return nil, err
		}
	}

	return &listing, nil
}
