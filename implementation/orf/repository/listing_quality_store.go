package repository

import "gopkg.in/yaml.v3"

type ListingQualityStore struct {
	fs MutableGraphFS
}

type listingQualityDocument struct {
	Quality []ListingQuality `yaml:"quality"`
}

func NewListingQualityStore(
	fs MutableGraphFS,
) *ListingQualityStore {

	return &ListingQualityStore{

		fs: fs,
	}
}

func (s *ListingQualityStore) Load() ([]ListingQuality, error) {

	data, err :=
		s.fs.ReadFile(
			ListingQualityFilePath(),
		)

	if err != nil {

		return []ListingQuality{}, nil
	}

	var doc listingQualityDocument

	err =
		yaml.Unmarshal(
			data,
			&doc,
		)

	if err != nil {

		return nil, err
	}

	return doc.Quality, nil
}

func (s *ListingQualityStore) Append(
	item ListingQuality,
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
			listingQualityDocument{

				Quality: list,
			},
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		ListingQualityFilePath(),
		data,
	)
}

func ListingQualityFilePath() string {

	return "/listing_quality.yaml"
}
