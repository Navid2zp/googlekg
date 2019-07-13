package googlekg

type KGRequest struct {
	URL       string
}

type KG struct {
	Context         KGContext `json:"@context"`
	Type            string    `json:"@type"`
	ItemListElement []KGItem  `json:"itemListElement"`
}

type KGContext struct {
	Vocab               string `json:"@vocab"`
	Goog                string `json:"goog"`
	ResultScore         string `json:"resultScore"`
	DetailedDescription string `json:"detailedDescription"`
	EntitySearchResult  string `json:"EntitySearchResult"`
	KG                  string `json:"kg"`
}

type KGItem struct {
	Type        string       `json:"@type"`
	Result      KGItemResult `json:"result"`
	ResultScore float64      `json:"resultScore"`
}

type KGItemResult struct {
	ID                  string                `json:"@id"`
	Name                string                `json:"name"`
	Type                []string              `json:"@type"`
	Description         string                `json:"description"`
	Image               KGImage               `json:"image"`
	DetailedDescription KGDetailedDescription `json:"detailedDescription"`
	URL                 string                `json:"url"`
}

type KGImage struct {
	ContentUrl string `json:"contentUrl"`
	URL        string `json:"url"`
	License    string `json:"license"`
}

type KGDetailedDescription struct {
	ArticleBody string `json:"articleBody"`
	URL         string `json:"url"`
	License     string `json:"license"`
}
