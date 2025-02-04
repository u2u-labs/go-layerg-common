package masterdb

import "time"

type AssetQueryBuilder interface {
	WithChainId(chainId int32) AssetQueryBuilder
	WithCollectionId(collectionId string) AssetQueryBuilder
	WithTokenIds(tokenIds []string) AssetQueryBuilder
	WithOwner(owner string) AssetQueryBuilder
	WithCreatedAtFrom(createdAtFrom time.Time) AssetQueryBuilder
	WithCreatedAtTo(createdAtTo time.Time) AssetQueryBuilder
	WithPage(page int) AssetQueryBuilder
	WithLimit(limit int) AssetQueryBuilder
	Build() AssetQueryFunction
}

type AssetQueryFunction interface {
	GetAssetQueryBuilder() (*assetQueryBuilderParam, error)
	GetPaginatedAsset() (any, error)
}
