package runtime

import (
	"context"
	"os"
	"time"

	"github.com/u2u-labs/go-layerg-common/api"
	"github.com/u2u-labs/go-layerg-common/rtapi"
)

type LayerGModule interface {
	AuthenticateApple(ctx context.Context, token, username string, create bool) (string, string, bool, error)
	AuthenticateCustom(ctx context.Context, id, username string, create bool) (string, string, bool, error)
	AuthenticateDevice(ctx context.Context, id, username string, create bool) (string, string, bool, error)
	AuthenticateEmail(ctx context.Context, email, password, username string, create bool) (string, string, bool, error)
	AuthenticateFacebook(ctx context.Context, token string, importFriends bool, username string, create bool) (string, string, bool, error)
	AuthenticateFacebookInstantGame(ctx context.Context, signedPlayerInfo string, username string, create bool) (string, string, bool, error)
	AuthenticateGameCenter(ctx context.Context, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl, username string, create bool) (string, string, bool, error)
	AuthenticateGoogle(ctx context.Context, token, username string, create bool, code, state, errStr string) (string, string, bool, error)
	AuthenticateSteam(ctx context.Context, token, username string, create bool) (string, string, bool, error)

	AuthenticateTokenGenerate(userID, username string, exp int64, vars map[string]string) (string, int64, error)

	AccountGetId(ctx context.Context, userID string) (*api.Account, error)
	AccountsGetId(ctx context.Context, userIDs []string) ([]*api.Account, error)
	AccountUpdateId(ctx context.Context, userID, username string, metadata map[string]interface{}, displayName, timezone, location, langTag, avatarUrl string) error

	AccountDeleteId(ctx context.Context, userID string, recorded bool) error
	AccountExportId(ctx context.Context, userID string) (string, error)

	BuildContractCallRequest(ctx context.Context, in ContractCallParams) (*TransactionRequest, error)
	SendUAOnchainTX(ctx context.Context, userID string, in UATransactionRequest) (*ReceiptResponse, error)

	UsersGetId(ctx context.Context, userIDs []string, facebookIDs []string) ([]*api.User, error)
	UsersGetUsername(ctx context.Context, usernames []string) ([]*api.User, error)
	UsersGetRandom(ctx context.Context, count int) ([]*api.User, error)
	UsersBanId(ctx context.Context, userIDs []string) error
	UsersUnbanId(ctx context.Context, userIDs []string) error

	LinkApple(ctx context.Context, userID, token string) error
	LinkCustom(ctx context.Context, userID, customID string) error
	LinkDevice(ctx context.Context, userID, deviceID string) error
	LinkEmail(ctx context.Context, userID, email, password string) error
	LinkFacebook(ctx context.Context, userID, username, token string, importFriends bool) error
	LinkFacebookInstantGame(ctx context.Context, userID, signedPlayerInfo string) error
	LinkGameCenter(ctx context.Context, userID, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl string) error
	LinkGoogle(ctx context.Context, userID, token string) error
	LinkSteam(ctx context.Context, userID, username, token string, importFriends bool) error

	CronPrev(expression string, timestamp int64) (int64, error)
	CronNext(expression string, timestamp int64) (int64, error)
	ReadFile(path string) (*os.File, error)

	UnlinkApple(ctx context.Context, userID, token string) error
	UnlinkCustom(ctx context.Context, userID, customID string) error
	UnlinkDevice(ctx context.Context, userID, deviceID string) error
	UnlinkEmail(ctx context.Context, userID, email string) error
	UnlinkFacebook(ctx context.Context, userID, token string) error
	UnlinkFacebookInstantGame(ctx context.Context, userID, signedPlayerInfo string) error
	UnlinkGameCenter(ctx context.Context, userID, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl string) error
	UnlinkGoogle(ctx context.Context, userID, token string) error
	UnlinkSteam(ctx context.Context, userID, token string) error

	StreamUserList(mode uint8, subject, subcontext, label string, includeHidden, includeNotHidden bool) ([]Presence, error)
	StreamUserGet(mode uint8, subject, subcontext, label, userID, sessionID string) (PresenceMeta, error)
	StreamUserJoin(mode uint8, subject, subcontext, label, userID, sessionID string, hidden, persistence bool, status string) (bool, error)
	StreamUserUpdate(mode uint8, subject, subcontext, label, userID, sessionID string, hidden, persistence bool, status string) error
	StreamUserLeave(mode uint8, subject, subcontext, label, userID, sessionID string) error
	StreamUserKick(mode uint8, subject, subcontext, label string, presence Presence) error
	StreamCount(mode uint8, subject, subcontext, label string) (int, error)
	StreamClose(mode uint8, subject, subcontext, label string) error
	StreamSend(mode uint8, subject, subcontext, label, data string, presences []Presence, reliable bool) error
	StreamSendRaw(mode uint8, subject, subcontext, label string, msg *rtapi.Envelope, presences []Presence, reliable bool) error

	SessionDisconnect(ctx context.Context, sessionID string, reason ...PresenceReason) error
	SessionLogout(userID, token, refreshToken string) error

	MatchCreate(ctx context.Context, module string, params map[string]interface{}) (string, error)
	MatchGet(ctx context.Context, id string) (*api.Match, error)
	MatchList(ctx context.Context, limit int, authoritative bool, label string, minSize, maxSize *int, query string) ([]*api.Match, error)
	MatchSignal(ctx context.Context, id string, data string) (string, error)

	NotificationSend(ctx context.Context, userID, subject string, content map[string]interface{}, code int, sender string, persistent bool) error
	NotificationsSend(ctx context.Context, notifications []*NotificationSend) error
	NotificationSendAll(ctx context.Context, subject string, content map[string]interface{}, code int, persistent bool) error
	NotificationsDelete(ctx context.Context, notifications []*NotificationDelete) error

	WalletUpdate(ctx context.Context, userID string, changeset map[string]int64, metadata map[string]interface{}, updateLedger bool) (updated map[string]int64, previous map[string]int64, err error)
	WalletsUpdate(ctx context.Context, updates []*WalletUpdate, updateLedger bool) ([]*WalletUpdateResult, error)
	WalletLedgerUpdate(ctx context.Context, itemID string, metadata map[string]interface{}) (WalletLedgerItem, error)
	WalletLedgerList(ctx context.Context, userID string, limit int, cursor string) ([]WalletLedgerItem, string, error)

	StorageList(ctx context.Context, callerID, userID, collection string, limit int, cursor string) ([]*api.StorageObject, string, error)
	StorageRead(ctx context.Context, reads []*StorageRead) ([]*api.StorageObject, error)
	StorageWrite(ctx context.Context, writes []*StorageWrite) ([]*api.StorageObjectAck, error)
	StorageDelete(ctx context.Context, deletes []*StorageDelete) error
	StorageIndexList(ctx context.Context, callerID, indexName, query string, limit int, order []string) (*api.StorageObjects, error)

	MultiUpdate(ctx context.Context, accountUpdates []*AccountUpdate, storageWrites []*StorageWrite, storageDeletes []*StorageDelete, walletUpdates []*WalletUpdate, updateLedger bool) ([]*api.StorageObjectAck, []*WalletUpdateResult, error)

	LeaderboardCreate(ctx context.Context, id string, authoritative bool, sortOrder, operator, resetSchedule string, metadata map[string]interface{}, enableRanks bool) error
	LeaderboardDelete(ctx context.Context, id string) error
	LeaderboardList(limit int, cursor string) (*api.LeaderboardList, error)
	LeaderboardRecordsList(ctx context.Context, id string, ownerIDs []string, limit int, cursor string, expiry int64) (records []*api.LeaderboardRecord, ownerRecords []*api.LeaderboardRecord, nextCursor string, prevCursor string, err error)
	LeaderboardRecordsListCursorFromRank(id string, rank, overrideExpiry int64) (string, error)
	LeaderboardRecordWrite(ctx context.Context, id, ownerID, username string, score, subscore int64, metadata map[string]interface{}, overrideOperator *int) (*api.LeaderboardRecord, error)
	LeaderboardRecordDelete(ctx context.Context, id, ownerID string) error
	LeaderboardsGetId(ctx context.Context, ids []string) ([]*api.Leaderboard, error)
	LeaderboardRecordsHaystack(ctx context.Context, id, ownerID string, limit int, cursor string, expiry int64) (*api.LeaderboardRecordList, error)

	PurchaseValidateApple(ctx context.Context, userID, receipt string, persist bool, passwordOverride ...string) (*api.ValidatePurchaseResponse, error)
	PurchaseValidateGoogle(ctx context.Context, userID, receipt string, persist bool, overrides ...struct {
		ClientEmail string
		PrivateKey  string
	}) (*api.ValidatePurchaseResponse, error)
	PurchaseValidateHuawei(ctx context.Context, userID, signature, inAppPurchaseData string, persist bool) (*api.ValidatePurchaseResponse, error)
	PurchaseValidateFacebookInstant(ctx context.Context, userID, signedRequest string, persist bool) (*api.ValidatePurchaseResponse, error)
	PurchasesList(ctx context.Context, userID string, limit int, cursor string) (*api.PurchaseList, error)
	PurchaseGetByTransactionId(ctx context.Context, transactionID string) (*api.ValidatedPurchase, error)

	SubscriptionValidateApple(ctx context.Context, userID, receipt string, persist bool, passwordOverride ...string) (*api.ValidateSubscriptionResponse, error)
	SubscriptionValidateGoogle(ctx context.Context, userID, receipt string, persist bool, overrides ...struct {
		ClientEmail string
		PrivateKey  string
	}) (*api.ValidateSubscriptionResponse, error)
	SubscriptionsList(ctx context.Context, userID string, limit int, cursor string) (*api.SubscriptionList, error)
	SubscriptionGetByProductId(ctx context.Context, userID, productID string) (*api.ValidatedSubscription, error)

	TournamentCreate(ctx context.Context, id string, authoritative bool, sortOrder, operator, resetSchedule string, metadata map[string]interface{}, title, description string, category, startTime, endTime, duration, maxSize, maxNumScore int, joinRequired, enableRanks bool) error
	TournamentDelete(ctx context.Context, id string) error
	TournamentAddAttempt(ctx context.Context, id, ownerID string, count int) error
	TournamentJoin(ctx context.Context, id, ownerID, username string) error
	TournamentsGetId(ctx context.Context, tournamentIDs []string) ([]*api.Tournament, error)
	TournamentList(ctx context.Context, categoryStart, categoryEnd, startTime, endTime, limit int, cursor string) (*api.TournamentList, error)
	TournamentRecordsList(ctx context.Context, tournamentId string, ownerIDs []string, limit int, cursor string, overrideExpiry int64) (records []*api.LeaderboardRecord, ownerRecords []*api.LeaderboardRecord, prevCursor string, nextCursor string, err error)
	TournamentRecordWrite(ctx context.Context, id, ownerID, username string, score, subscore int64, metadata map[string]interface{}, operatorOverride *int) (*api.LeaderboardRecord, error)
	TournamentRecordDelete(ctx context.Context, id, ownerID string) error
	TournamentRecordsHaystack(ctx context.Context, id, ownerID string, limit int, cursor string, expiry int64) (*api.TournamentRecordList, error)

	GroupsGetId(ctx context.Context, groupIDs []string) ([]*api.Group, error)
	GroupCreate(ctx context.Context, userID, name, creatorID, langTag, description, avatarUrl string, open bool, metadata map[string]interface{}, maxCount int) (*api.Group, error)
	GroupUpdate(ctx context.Context, id, userID, name, creatorID, langTag, description, avatarUrl string, open bool, metadata map[string]interface{}, maxCount int) error
	GroupDelete(ctx context.Context, id string) error
	GroupUserJoin(ctx context.Context, groupID, userID, username string) error
	GroupUserLeave(ctx context.Context, groupID, userID, username string) error
	GroupUsersAdd(ctx context.Context, callerID, groupID string, userIDs []string) error
	GroupUsersBan(ctx context.Context, callerID, groupID string, userIDs []string) error
	GroupUsersKick(ctx context.Context, callerID, groupID string, userIDs []string) error
	GroupUsersPromote(ctx context.Context, callerID, groupID string, userIDs []string) error
	GroupUsersDemote(ctx context.Context, callerID, groupID string, userIDs []string) error
	GroupUsersList(ctx context.Context, id string, limit int, state *int, cursor string) ([]*api.GroupUserList_GroupUser, string, error)
	GroupsList(ctx context.Context, name, langTag string, members *int, open *bool, limit int, cursor string) ([]*api.Group, string, error)
	GroupsGetRandom(ctx context.Context, count int) ([]*api.Group, error)
	UserGroupsList(ctx context.Context, userID string, limit int, state *int, cursor string) ([]*api.UserGroupList_UserGroup, string, error)

	FriendsList(ctx context.Context, userID string, limit int, state *int, cursor string) ([]*api.Friend, string, error)
	FriendsOfFriendsList(ctx context.Context, userID string, limit int, cursor string) ([]*api.FriendsOfFriendsList_FriendOfFriend, string, error)
	FriendsAdd(ctx context.Context, userID string, username string, ids []string, usernames []string) error
	FriendsDelete(ctx context.Context, userID string, username string, ids []string, usernames []string) error
	FriendsBlock(ctx context.Context, userID string, username string, ids []string, usernames []string) error

	Event(ctx context.Context, evt *api.Event) error

	MetricsCounterAdd(name string, tags map[string]string, delta int64)
	MetricsGaugeSet(name string, tags map[string]string, value float64)
	MetricsTimerRecord(name string, tags map[string]string, value time.Duration)

	ChannelIdBuild(ctx context.Context, sender string, target string, chanType ChannelType) (string, error)
	ChannelMessageSend(ctx context.Context, channelID string, content map[string]interface{}, senderId, senderUsername string, persist bool) (*rtapi.ChannelMessageAck, error)
	ChannelMessageUpdate(ctx context.Context, channelID, messageID string, content map[string]interface{}, senderId, senderUsername string, persist bool) (*rtapi.ChannelMessageAck, error)
	ChannelMessageRemove(ctx context.Context, channelId, messageId string, senderId, senderUsername string, persist bool) (*rtapi.ChannelMessageAck, error)
	ChannelMessagesList(ctx context.Context, channelId string, limit int, forward bool, cursor string) (messages []*api.ChannelMessage, nextCursor string, prevCursor string, err error)

	GetSatori() Satori
	GetFleetManager() FleetManager

	RegisterWebhook(ctx context.Context, config WebhookConfig, handler WebhookHandler) error

	GetNFTs(ctx context.Context, params NFTQueryParams) (*NFTResponse, error)
	GetAggToken(ctx context.Context) (string, error)
	GetCollectionAsset(ctx context.Context, params CollectionAssetQueryParams, token string) (*CollectionAssetResponse, error)

	GetRequiredHeadersUA() (map[string]string, error)
}

type Media struct {
	ID      string `json:"id"`
	S3Url   string `json:"S3Url"`
	IPFSUrl string `json:"IPFSUrl"`
	AssetId string `json:"AssetId"`
}

// Struct to represent individual metadata attributes
type Attribute struct {
	Value     interface{} `json:"value"`
	TraitType string      `json:"trait_type"`
}

// Struct to represent metadata details
type MetadataDetails struct {
	Creator    string      `json:"creator"`
	Attributes []Attribute `json:"attributes"`
}

type Metadata struct {
	ID       string          `json:"id"`
	Metadata MetadataDetails `json:"metadata"`
	IPFSUrl  string          `json:"IPFSUrl"`
	AssetId  string          `json:"AssetId"`
}

type NFTData struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	TokenId           string   `json:"tokenId"`
	Description       string   `json:"description"`
	CollectionAddress string   `json:"collectionAddress"`
	Media             Media    `json:"media"`
	Metadata          Metadata `json:"metadata"`
	OffChainBalance   string   `json:"offChainBalance"`
	OnChainBalance    string   `json:"onChainBalance"`
	OwnerAddress      string   `json:"ownerAddress"`
	Type              string   `json:"type"`
}

// Struct to represent pagination information
type Paging struct {
	Page    int  `json:"page"`
	Limit   int  `json:"limit"`
	HasNext bool `json:"hasNext"`
}

// Main struct to represent the NFT response
type NFTResponse struct {
	Data   []NFTData `json:"data"`
	Paging Paging    `json:"paging"`
}

type CollectionAssetResponse struct {
	Data   []AssetData `json:"data"`
	Paging Paging      `json:"paging"`
}

type AssetData struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	TokenId           string   `json:"tokenId"`
	Description       string   `json:"description"`
	CollectionAddress string   `json:"collectionAddress"`
	Media             Media    `json:"media"`
	Metadata          Metadata `json:"metadata"`
}
type NFTQueryParams struct {
	CollectionAddress string   `json:"collectionAddress" required:"true"`
	Mode              string   `json:"mode"`
	Page              string   `json:"page"`
	Limit             string   `json:"limit"`
	OwnerAddress      string   `json:"ownerAddress"`
	TokenIds          []string `json:"tokenIds"`
}

type CollectionAssetQueryParams struct {
	CollectionId string `json:"collectionId" required:"true"`
	Page         string `json:"page"`
	Limit        string `json:"limit"`
}

type MessageList struct {
	Messages        []*Message `json:"messages,omitempty"`
	NextCursor      string     `json:"next_cursor,omitempty"`
	PrevCursor      string     `json:"prev_cursor,omitempty"`
	CacheableCursor string     `json:"cacheable_cursor,omitempty"`
}

type Message struct {
	ScheduleId  string         `json:"schedule_id,omitempty"`
	SendTime    int64          `json:"send_time,string,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreateTime  int64          `json:"create_time,string,omitempty"`
	UpdateTime  int64          `json:"update_time,string,omitempty"`
	ReadTime    int64          `json:"read_time,string,omitempty"`
	ConsumeTime int64          `json:"consume_time,string,omitempty"`
	Text        string         `json:"text,omitempty"`
	Id          string         `json:"id,omitempty"`
	Title       string         `json:"title,omitempty"`
	ImageUrl    string         `json:"image_url,omitempty"`
}

type MessageUpdate struct {
	ReadTime    int64 `json:"read_time,omitempty"`
	ConsumeTime int64 `json:"consume_time,omitempty"`
}
