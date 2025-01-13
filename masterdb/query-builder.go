type CommonAssetData struct {
	ID           uuid.UUID `json:"id"`
	ChainID      int32     `json:"chainId"`
	CollectionID string    `json:"collectionId"`
	Owner        string    `json:"owner"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    uuid.UUID `json:"updatedBy"`
}

type assetQueryBuilderParam struct {
	chainId       int32
	collectionId  *string
	tokenId       *string
	owner         *string
	createdAtFrom *time.Time
	createdAtTo   *time.Time
	page          *int
	limit         *int
	offset        *int
}

type AssetQueryFunction interface {
	GetAssetQueryBuilder()
	GetPaginatedAsset() Pagination[CommonAssetData]
}

type AssetQueryBuilder interface {
	WithChainId(chainId int32) AssetQueryBuilder
	WithCollectionId(collectionId string) AssetQueryBuilder
	WithTokenId(tokenId string) AssetQueryBuilder
	WithOwner(owner string) AssetQueryBuilder
	WithCreatedAtFrom(createdAtFrom time.Time) AssetQueryBuilder
	WithCreatedAtTo(createdAtTo time.Time) AssetQueryBuilder
	WithPage(page int) AssetQueryBuilder
	WithLimit(limit int) AssetQueryBuilder
	WithOffset(offset int) AssetQueryBuilder
	Build() AssetQueryFunction
}

func (b *assetQueryBuilderParam) Build() AssetQueryFunction {
	return b
}

// GetAssetQueryBuilder implements AssetQueryFunction.
func (b *assetQueryBuilderParam) GetAssetQueryBuilder() {
	panic("unimplemented")
}

// GetPaginatedAsset implements AssetQueryFunction.
func (b *assetQueryBuilderParam) GetPaginatedAsset() Pagination[CommonAssetData] {
	panic("unimplemented")
}

func NewAssetQueryBuilder() AssetQueryBuilder {
	return &assetQueryBuilderParam{}
}

// WithChainId implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithChainId(chainId int32) AssetQueryBuilder {
	b.chainId = chainId
	return b
}

// WithCollectionId implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithCollectionId(collectionId string) AssetQueryBuilder {
	b.collectionId = &collectionId
	return b
}

// WithCreatedAtFrom implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithCreatedAtFrom(createdAtFrom time.Time) AssetQueryBuilder {
	b.createdAtFrom = &createdAtFrom
	return b
}

// WithCreatedAtTo implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithCreatedAtTo(createdAtTo time.Time) AssetQueryBuilder {
	b.createdAtTo = &createdAtTo
	return b
}

// WithLimit implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithLimit(limit int) AssetQueryBuilder {
	// set limit to 10 if limit is <= 0
	defaultLimit := 10
	if limit <= 0 {
		b.limit = &defaultLimit
	} else {
		b.limit = &limit
	}
	return b
}

// WithOffset implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithOffset(offset int) AssetQueryBuilder {
	// set offset to 0 if offset is < 0
	defaultOffset := 0
	if offset < 0 {
		b.offset = &defaultOffset
	} else {
		b.offset = &offset
	}
	return b
}

// WithOwner implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithOwner(owner string) AssetQueryBuilder {
	b.owner = &owner
	return b
}

// WithPage implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithPage(page int) AssetQueryBuilder {
	// set page to 1 if page is <= 0
	defaultPage := 1
	if page <= 0 {
		b.page = &defaultPage
	} else {
		b.page = &page
	}
	return b
}

// WithTokenId implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithTokenId(tokenId string) AssetQueryBuilder {
	b.tokenId = &tokenId
	return b
}

// IMPLEMENTATION
// asset := NewAssetQueryBuilder().
// 		WithChainId(1).
// 		WithCollectionId("1").
// 		WithTokenId("1").
// 		WithOwner("1").
// 		WithCreatedAtFrom(time.Now().Add(-1 * time.Hour)).
// 		WithCreatedAtTo(time.Now()).
// 		WithPage(1).
// 		WithLimit(10).
// 		WithOffset(0).
// 		Build()

// 	asset.GetPaginatedAsset()
