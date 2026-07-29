package repository

type KeywordStrategy struct {
	Primary []string `json:"primary"`

	Secondary []string `json:"secondary"`

	LongTail []string `json:"long_tail"`
}

type KeywordStrategyEngine struct{}

func NewKeywordStrategyEngine() *KeywordStrategyEngine {

	return &KeywordStrategyEngine{}
}

func (k *KeywordStrategyEngine) Generate(
	keywords []string,
) KeywordStrategy {

	result :=
		KeywordStrategy{}

	for i, key := range keywords {

		if i == 0 {

			result.Primary =
				append(
					result.Primary,
					key,
				)

			continue
		}

		if i == 1 {

			result.Secondary =
				append(
					result.Secondary,
					key,
				)

			continue
		}

		result.LongTail =
			append(
				result.LongTail,
				key,
			)
	}

	return result
}
