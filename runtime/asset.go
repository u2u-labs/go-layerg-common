package runtime

type CreateNFTArgs struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	TokenId      string `json:"tokenId"`
	CollectionId string `json:"collectionId"`
	Quantity     string `json:"quantity"`
	Media        struct {
		S3Url string `json:"S3Url"`
	} `json:"media"`
	Metadata struct {
		Metadata NFTMetadata `json:"metadata"`
	} `json:"metadata"`
}

type NFTMetadata struct {
	Attributes []MetadataAttribute `json:"attributes"`
}

type MetadataAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}
