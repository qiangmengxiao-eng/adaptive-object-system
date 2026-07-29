package repository

import (
	"gopkg.in/yaml.v3"
)

const ListingFileName = "listings.yaml"

type ListingStore struct {
	fs MutableGraphFS
}

type listingDocument struct {
	Listings []AmazonListing `yaml:"listings"`
}

// NewListingStore creates listing storage.
func NewListingStore(
	fs MutableGraphFS,
) *ListingStore {

	return &ListingStore{

		fs: fs,
	}
}

// Load listings.
func (s *ListingStore) Load() ([]AmazonListing, error) {

	data, err :=
		s.fs.ReadFile(
			ListingFilePath(),
		)

	if err != nil {

		return []AmazonListing{}, nil
	}

	var doc listingDocument

	err =
		yaml.Unmarshal(
			data,
			&doc,
		)

	if err != nil {

		return nil, err
	}

	return doc.Listings, nil
}

// Append listing.
func (s *ListingStore) Append(
	item AmazonListing,
) error {

	list, err :=
		s.Load()

	if err != nil {

		return err
	}

	list =
		append(
			list,
			item,
		)

	data, err :=
		yaml.Marshal(
			listingDocument{

				Listings: list,
			},
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		ListingFilePath(),
		data,
	)
}

func ListingFilePath() string {

	return "/listings.yaml"
}
