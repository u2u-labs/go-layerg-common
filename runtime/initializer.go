package runtime

import (
	"context"
	"database/sql"

	"github.com/u2u-labs/go-layerg-common/api"
	"github.com/u2u-labs/go-layerg-common/rtapi"
)

/*
Initializer is used to register various callback functions with the server.
It is made available to the InitModule function as an input parameter when the function is invoked by the server when loading the module on server start.

NOTE: You must not cache the reference to this and reuse it as a later point as this could have unintended side effects.
*/
type Initializer interface {
	/*
		RegisterRpc registers a function with the given ID. This ID can be used within client code to send an RPC message to
		execute the function and return the result. Results are always returned as a JSON string (or optionally empty string).

		If there is an issue with the RPC call, return an empty string and the associated error which will be returned to the client.
	*/
	RegisterRpc(id string, fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, payload string) (string, error)) error

	/*
		RegisterBeforeRt registers a function with for a message. Any function may be registered to intercept a message received from a client and operate on it (or reject it) based on custom logic.
		This is useful to enforce specific rules on top of the standard features in the server.

		You can return `nil` instead of the `rtapi.Envelope` and this will disable that particular server functionality.

		Message names can be found here: https://heroiclabs.com/docs/runtime-code-basics/#message-names
	*/
	RegisterBeforeRt(id string, fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *rtapi.Envelope) (*rtapi.Envelope, error)) error

	/*
		RegisterAfterRt registers a function for a message. The registered function will be called after the message has been processed in the pipeline.
		The custom code will be executed asynchronously after the response message has been sent to a client

		Message names can be found here: https://heroiclabs.com/docs/runtime-code-basics/#message-names
	*/
	RegisterAfterRt(id string, fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out, in *rtapi.Envelope) error) error

	// RegisterMatchmakerMatched
	RegisterMatchmakerMatched(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, entries []MatchmakerEntry) (string, error)) error

	// RegisterMatchmakerOverride
	RegisterMatchmakerOverride(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, candidateMatches [][]MatchmakerEntry) (matches [][]MatchmakerEntry)) error

	// RegisterMatch
	RegisterMatch(name string, fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule) (Match, error)) error

	// RegisterTournamentEnd
	RegisterTournamentEnd(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, tournament *api.Tournament, end, reset int64) error) error

	// RegisterTournamentReset
	RegisterTournamentReset(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, tournament *api.Tournament, end, reset int64) error) error

	// RegisterLeaderboardReset
	RegisterLeaderboardReset(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, leaderboard *api.Leaderboard, reset int64) error) error

	// RegisterPurchaseNotificationApple
	RegisterPurchaseNotificationApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, purchase *api.ValidatedPurchase, providerPayload string) error) error

	// RegisterSubscriptionNotificationApple
	RegisterSubscriptionNotificationApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, subscription *api.ValidatedSubscription, providerPayload string) error) error

	// RegisterPurchaseNotificationGoogle
	RegisterPurchaseNotificationGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, purchase *api.ValidatedPurchase, providerPayload string) error) error

	// RegisterSubscriptionNotificationGoogle
	RegisterSubscriptionNotificationGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, subscription *api.ValidatedSubscription, providerPayload string) error) error

	// RegisterBeforeGetAccount is used to register a function invoked when the server receives the relevant request.
	RegisterBeforeGetAccount(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule) error) error

	// RegisterAfterGetAccount is used to register a function invoked after the server processes the relevant request.
	RegisterAfterGetAccount(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Account) error) error

	// RegisterBeforeUpdateAccount is used to register a function invoked when the server receives the relevant request.
	RegisterBeforeUpdateAccount(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.UpdateAccountRequest) (*api.UpdateAccountRequest, error)) error

	// RegisterAfterUpdateAccount is used to register a function invoked after the server processes the relevant request.
	RegisterAfterUpdateAccount(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.UpdateAccountRequest) error) error

	// RegisterBeforeDeleteAccount is used to register a function invoked when the server receives the relevant request.
	RegisterBeforeDeleteAccount(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule) error) error

	// RegisterAfterDeleteAccount is used to register a function invoked after the server processes the relevant request.
	RegisterAfterDeleteAccount(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule) error) error

	// RegisterBeforeSessionRefresh can be used to perform pre-refresh checks.
	RegisterBeforeSessionRefresh(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.SessionRefreshRequest) (*api.SessionRefreshRequest, error)) error

	// RegisterAfterSessionRefresh can be used to perform after successful refresh checks.
	RegisterAfterSessionRefresh(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.SessionRefreshRequest) error) error

	// RegisterBeforeSessionLogout can be used to perform pre-logout checks.
	RegisterBeforeSessionLogout(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.SessionLogoutRequest) (*api.SessionLogoutRequest, error)) error

	// RegisterAfterSessionLogout can be used to perform after successful logout checks.
	RegisterAfterSessionLogout(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.SessionLogoutRequest) error) error

	// RegisterBeforeAuthenticateApple can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateAppleRequest) (*api.AuthenticateAppleRequest, error)) error

	// RegisterAfterAuthenticateApple can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateAppleRequest) error) error

	// RegisterBeforeAuthenticateCustom can be used to perform pre-authentication checks.
	// You can use this to process the input (such as decoding custom tokens) and ensure inter-compatibility between Layerg and your own custom system.
	RegisterBeforeAuthenticateCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error)) error

	// RegisterAfterAuthenticateCustom can be used to perform after successful authentication checks.
	// For instance, you can run special logic if the account was just created like adding them to newcomers leaderboard.
	RegisterAfterAuthenticateCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateCustomRequest) error) error

	// RegisterBeforeAuthenticateDevice can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateDeviceRequest) (*api.AuthenticateDeviceRequest, error)) error

	// RegisterAfterAuthenticateDevice can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateDeviceRequest) error) error

	// RegisterBeforeAuthenticateEmail can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateEmailRequest) (*api.AuthenticateEmailRequest, error)) error

	// RegisterAfterAuthenticateEmail can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateEmailRequest) error) error

	// RegisterBeforeAuthenticateTelegram can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateTelegram(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateTelegramRequest) (*api.AuthenticateTelegramRequest, error)) error

	// RegisterAfterAuthenticateTelegram can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateTelegram(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateTelegramRequest) error) error

	// RegisterBeforeAuthenticateEvm can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateEvm(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateEvmRequest) (*api.AuthenticateEvmRequest, error)) error

	// RegisterAfterAuthenticateEvm can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateEvm(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateEvmRequest) error) error

	// RegisterBeforeAuthenticateFacebook can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateFacebookRequest) (*api.AuthenticateFacebookRequest, error)) error

	// RegisterAfterAuthenticateFacebook can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateFacebookRequest) error) error

	// RegisterBeforeAuthenticateFacebookInstantGame can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateFacebookInstantGameRequest) (*api.AuthenticateFacebookInstantGameRequest, error)) error

	// RegisterAfterAuthenticateFacebookInstantGame can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateFacebookInstantGameRequest) error) error

	// RegisterBeforeAuthenticateGameCenter can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateGameCenterRequest) (*api.AuthenticateGameCenterRequest, error)) error

	// RegisterAfterAuthenticateGameCenter can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateGameCenterRequest) error) error

	// RegisterBeforeAuthenticateGoogle can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateGoogleRequest) (*api.AuthenticateGoogleRequest, error)) error

	// RegisterAfterAuthenticateGoogle can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateGoogleRequest) error) error

	// RegisterBeforeAuthenticateSteam can be used to perform pre-authentication checks.
	RegisterBeforeAuthenticateSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AuthenticateSteamRequest) (*api.AuthenticateSteamRequest, error)) error

	// RegisterAfterAuthenticateSteam can be used to perform after successful authentication checks.
	RegisterAfterAuthenticateSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Session, in *api.AuthenticateSteamRequest) error) error

	// RegisterBeforeListChannelMessages can be used to perform additional logic before listing messages on a channel.
	RegisterBeforeListChannelMessages(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListChannelMessagesRequest) (*api.ListChannelMessagesRequest, error)) error

	// RegisterAfterListChannelMessages can be used to perform additional logic after messages for a channel is listed.
	RegisterAfterListChannelMessages(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.ChannelMessageList, in *api.ListChannelMessagesRequest) error) error

	// RegisterBeforeListChannelMessages can be used to perform additional logic before listing friends.
	RegisterBeforeListFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListFriendsRequest) (*api.ListFriendsRequest, error)) error

	// RegisterAfterListFriends can be used to perform additional logic after friends are listed.
	RegisterAfterListFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.FriendList) error) error

	// RegisterBeforeListFriendsOfFriends can be used to perform additional logic before listing friends of friends.
	RegisterBeforeListFriendsOfFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListFriendsOfFriendsRequest) (*api.ListFriendsOfFriendsRequest, error)) error

	// RegisterAfterListFriendsOfFriends can be used to perform additional logic after listing friends of friends.
	RegisterAfterListFriendsOfFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.FriendsOfFriendsList) error) error

	// RegisterBeforeAddFriends can be used to perform additional logic before friends are added.
	RegisterBeforeAddFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AddFriendsRequest) (*api.AddFriendsRequest, error)) error

	// RegisterAfterAddFriends can be used to perform additional logic after friends are added.
	RegisterAfterAddFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AddFriendsRequest) error) error

	// RegisterBeforeDeleteFriends can be used to perform additional logic before friends are deleted.
	RegisterBeforeDeleteFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteFriendsRequest) (*api.DeleteFriendsRequest, error)) error

	// RegisterAfterDeleteFriends can be used to perform additional logic after friends are deleted.
	RegisterAfterDeleteFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteFriendsRequest) error) error

	// RegisterBeforeBlockFriends can be used to perform additional logic before friends are blocked.
	RegisterBeforeBlockFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.BlockFriendsRequest) (*api.BlockFriendsRequest, error)) error

	// RegisterAfterBlockFriends can be used to perform additional logic after friends are blocked.
	RegisterAfterBlockFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.BlockFriendsRequest) error) error

	// RegisterBeforeImportFacebookFriends can be used to perform additional logic before Facebook friends are imported.
	RegisterBeforeImportFacebookFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ImportFacebookFriendsRequest) (*api.ImportFacebookFriendsRequest, error)) error

	// RegisterAfterImportFacebookFriends can be used to perform additional logic after Facebook friends are imported.
	RegisterAfterImportFacebookFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ImportFacebookFriendsRequest) error) error

	// RegisterBeforeImportSteamFriends can be used to perform additional logic before Facebook friends are imported.
	RegisterBeforeImportSteamFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ImportSteamFriendsRequest) (*api.ImportSteamFriendsRequest, error)) error

	// RegisterAfterImportSteamFriends can be used to perform additional logic after Facebook friends are imported.
	RegisterAfterImportSteamFriends(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ImportSteamFriendsRequest) error) error

	// RegisterBeforeCreateGroup can be used to perform additional logic before a group is created.
	RegisterBeforeCreateGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.CreateGroupRequest) (*api.CreateGroupRequest, error)) error

	// RegisterAfterCreateGroup can be used to perform additional logic after a group is created.
	RegisterAfterCreateGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Group, in *api.CreateGroupRequest) error) error

	// RegisterBeforeUpdateGroup can be used to perform additional logic before a group is updated.
	RegisterBeforeUpdateGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.UpdateGroupRequest) (*api.UpdateGroupRequest, error)) error

	// RegisterAfterUpdateGroup can be used to perform additional logic after a group is updated.
	RegisterAfterUpdateGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.UpdateGroupRequest) error) error

	// RegisterBeforeDeleteGroup can be used to perform additional logic before a group is deleted.
	RegisterBeforeDeleteGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteGroupRequest) (*api.DeleteGroupRequest, error)) error

	// RegisterAfterDeleteGroup can be used to perform additional logic after a group is deleted.
	RegisterAfterDeleteGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteGroupRequest) error) error

	// RegisterBeforeJoinGroup can be used to perform additional logic before user joins a group.
	RegisterBeforeJoinGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.JoinGroupRequest) (*api.JoinGroupRequest, error)) error

	// RegisterAfterJoinGroup can be used to perform additional logic after user joins a group.
	RegisterAfterJoinGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.JoinGroupRequest) error) error

	// RegisterBeforeLeaveGroup can be used to perform additional logic before user leaves a group.
	RegisterBeforeLeaveGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.LeaveGroupRequest) (*api.LeaveGroupRequest, error)) error

	// RegisterAfterLeaveGroup can be used to perform additional logic after user leaves a group.
	RegisterAfterLeaveGroup(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.LeaveGroupRequest) error) error

	// RegisterBeforeAddGroupUsers can be used to perform additional logic before user is added to a group.
	RegisterBeforeAddGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AddGroupUsersRequest) (*api.AddGroupUsersRequest, error)) error

	// RegisterAfterAddGroupUsers can be used to perform additional logic after user is added to a group.
	RegisterAfterAddGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AddGroupUsersRequest) error) error

	// RegisterBeforeBanGroupUsers can be used to perform additional logic before user is banned from a group.
	RegisterBeforeBanGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.BanGroupUsersRequest) (*api.BanGroupUsersRequest, error)) error

	// RegisterAfterBanGroupUsers can be used to perform additional logic after user is banned from a group.
	RegisterAfterBanGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.BanGroupUsersRequest) error) error

	// RegisterBeforeKickGroupUsers can be used to perform additional logic before user is kicked to a group.
	RegisterBeforeKickGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.KickGroupUsersRequest) (*api.KickGroupUsersRequest, error)) error

	// RegisterAfterKickGroupUsers can be used to perform additional logic after user is kicked from a group.
	RegisterAfterKickGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.KickGroupUsersRequest) error) error

	// RegisterBeforePromoteGroupUsers can be used to perform additional logic before user is promoted.
	RegisterBeforePromoteGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.PromoteGroupUsersRequest) (*api.PromoteGroupUsersRequest, error)) error

	// RegisterAfterPromoteGroupUsers can be used to perform additional logic after user is promoted.
	RegisterAfterPromoteGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.PromoteGroupUsersRequest) error) error

	// RegisterBeforeDemoteGroupUsers can be used to perform additional logic before user is demoted.
	RegisterBeforeDemoteGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DemoteGroupUsersRequest) (*api.DemoteGroupUsersRequest, error)) error

	// RegisterAfterDemoteGroupUsers can be used to perform additional logic after user is demoted.
	RegisterAfterDemoteGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DemoteGroupUsersRequest) error) error

	// RegisterBeforeListGroupUsers can be used to perform additional logic before users in a group is listed.
	RegisterBeforeListGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListGroupUsersRequest) (*api.ListGroupUsersRequest, error)) error

	// RegisterAfterListGroupUsers can be used to perform additional logic after users in a group is listed.
	RegisterAfterListGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.GroupUserList, in *api.ListGroupUsersRequest) error) error

	// RegisterBeforeListUserGroups can be used to perform additional logic before groups for a user is listed.
	RegisterBeforeListUserGroups(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListUserGroupsRequest) (*api.ListUserGroupsRequest, error)) error

	// RegisterAfterListUserGroups can be used to perform additional logic after groups for a user is listed.
	RegisterAfterListUserGroups(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.UserGroupList, in *api.ListUserGroupsRequest) error) error

	// RegisterBeforeListGroups can be used to perform additional logic before groups are listed.
	RegisterBeforeListGroups(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListGroupsRequest) (*api.ListGroupsRequest, error)) error

	// RegisterAfterListGroups can be used to perform additional logic after groups are listed.
	RegisterAfterListGroups(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.GroupList, in *api.ListGroupsRequest) error) error

	// RegisterBeforeDeleteLeaderboardRecord can be used to perform additional logic before deleting record from a leaderboard.
	RegisterBeforeDeleteLeaderboardRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteLeaderboardRecordRequest) (*api.DeleteLeaderboardRecordRequest, error)) error

	// RegisterAfterDeleteLeaderboardRecord can be used to perform additional logic after deleting record from a leaderboard.
	RegisterAfterDeleteLeaderboardRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteLeaderboardRecordRequest) error) error

	// RegisterBeforeDeleteTournamentRecord can be used to perform additional logic before deleting record from a leaderboard.
	RegisterBeforeDeleteTournamentRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteTournamentRecordRequest) (*api.DeleteTournamentRecordRequest, error)) error

	// RegisterAfterDeleteTournamentRecord can be used to perform additional logic after deleting record from a leaderboard.
	RegisterAfterDeleteTournamentRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteTournamentRecordRequest) error) error

	// RegisterBeforeListLeaderboardRecords can be used to perform additional logic before listing records from a leaderboard.
	RegisterBeforeListLeaderboardRecords(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListLeaderboardRecordsRequest) (*api.ListLeaderboardRecordsRequest, error)) error

	// RegisterAfterListLeaderboardRecords  can be used to perform additional logic after listing records from a leaderboard.
	RegisterAfterListLeaderboardRecords(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsRequest) error) error

	// RegisterBeforeWriteLeaderboardRecord can be used to perform additional logic before submitting new record to a leaderboard.
	RegisterBeforeWriteLeaderboardRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.WriteLeaderboardRecordRequest) (*api.WriteLeaderboardRecordRequest, error)) error

	// RegisterAfterWriteLeaderboardRecord can be used to perform additional logic after submitting new record to a leaderboard.
	RegisterAfterWriteLeaderboardRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.LeaderboardRecord, in *api.WriteLeaderboardRecordRequest) error) error

	// RegisterBeforeListLeaderboardRecordsAroundOwner can be used to perform additional logic before listing records from a leaderboard.
	RegisterBeforeListLeaderboardRecordsAroundOwner(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListLeaderboardRecordsAroundOwnerRequest) (*api.ListLeaderboardRecordsAroundOwnerRequest, error)) error

	// RegisterAfterListLeaderboardRecordsAroundOwner can be used to perform additional logic after listing records from a leaderboard.
	RegisterAfterListLeaderboardRecordsAroundOwner(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsAroundOwnerRequest) error) error

	// RegisterBeforeLinkApple can be used to perform additional logic before linking Apple ID to an account.
	RegisterBeforeLinkApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountApple) (*api.AccountApple, error)) error

	// RegisterAfterLinkApple can be used to perform additional logic after linking Apple ID to an account.
	RegisterAfterLinkApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountApple) error) error

	// RegisterBeforeLinkCustom can be used to perform additional logic before linking custom ID to an account.
	RegisterBeforeLinkCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountCustom) (*api.AccountCustom, error)) error

	// RegisterAfterLinkCustom can be used to perform additional logic after linking custom ID to an account.
	RegisterAfterLinkCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountCustom) error) error

	// RegisterBeforeLinkDevice can be used to perform additional logic before linking device ID to an account.
	RegisterBeforeLinkDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountDevice) (*api.AccountDevice, error)) error

	// RegisterAfterLinkDevice can be used to perform additional logic after linking device ID to an account.
	RegisterAfterLinkDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountDevice) error) error

	// RegisterBeforeLinkEmail can be used to perform additional logic before linking email to an account.
	RegisterBeforeLinkEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountEmail) (*api.AccountEmail, error)) error

	// RegisterAfterLinkEmail can be used to perform additional logic after linking email to an account.
	RegisterAfterLinkEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountEmail) error) error

	// RegisterBeforeLinkFacebook can be used to perform additional logic before linking Facebook to an account.
	RegisterBeforeLinkFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.LinkFacebookRequest) (*api.LinkFacebookRequest, error)) error

	// RegisterAfterLinkFacebook can be used to perform additional logic after linking Facebook to an account.
	RegisterAfterLinkFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.LinkFacebookRequest) error) error

	// RegisterBeforeLinkFacebookInstantGame can be used to perform additional logic before linking Facebook Instant Game profile to an account.
	RegisterBeforeLinkFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error)) error

	// RegisterAfterLinkFacebookInstantGame can be used to perform additional logic after linking Facebook Instant Game profile to an account.
	RegisterAfterLinkFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountFacebookInstantGame) error) error

	// RegisterBeforeLinkGameCenter can be used to perform additional logic before linking GameCenter to an account.
	RegisterBeforeLinkGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountGameCenter) (*api.AccountGameCenter, error)) error

	// RegisterAfterLinkGameCenter can be used to perform additional logic after linking GameCenter to an account.
	RegisterAfterLinkGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountGameCenter) error) error

	// RegisterBeforeLinkGoogle can be used to perform additional logic before linking Google to an account.
	RegisterBeforeLinkGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountGoogle) (*api.AccountGoogle, error)) error

	// RegisterAfterLinkGoogle can be used to perform additional logic after linking Google to an account.
	RegisterAfterLinkGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountGoogle) error) error

	// RegisterBeforeLinkSteam can be used to perform additional logic before linking Steam to an account.
	RegisterBeforeLinkSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.LinkSteamRequest) (*api.LinkSteamRequest, error)) error

	// RegisterAfterLinkSteam can be used to perform additional logic after linking Steam to an account.
	RegisterAfterLinkSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.LinkSteamRequest) error) error

	// RegisterBeforeListMatches can be used to perform additional logic before listing matches.
	RegisterBeforeListMatches(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListMatchesRequest) (*api.ListMatchesRequest, error)) error

	// RegisterAfterListMatches can be used to perform additional logic after listing matches.
	RegisterAfterListMatches(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.MatchList, in *api.ListMatchesRequest) error) error

	// RegisterBeforeMatchmakerStats is used to register a function invoked when the server receives the relevant request.
	RegisterBeforeGetMatchmakerStats(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule) error) error

	// RegisterAfterMarchmakerStats is used to register a function invoked after the server processes the relevant request.
	RegisterAfterGetMatchmakerStats(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.MatchmakerStats) error) error

	// RegisterBeforeListNotifications can be used to perform additional logic before listing notifications for a user.
	RegisterBeforeListNotifications(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListNotificationsRequest) (*api.ListNotificationsRequest, error)) error

	// RegisterAfterListNotifications can be used to perform additional logic after listing notifications for a user.
	RegisterAfterListNotifications(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.NotificationList, in *api.ListNotificationsRequest) error) error

	// RegisterBeforeDeleteNotifications can be used to perform additional logic before deleting notifications.
	RegisterBeforeDeleteNotifications(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteNotificationsRequest) (*api.DeleteNotificationsRequest, error)) error

	// RegisterAfterDeleteNotifications can be used to perform additional logic after deleting notifications.
	RegisterAfterDeleteNotifications(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteNotificationsRequest) error) error

	// RegisterBeforeListStorageObjects can be used to perform additional logic before listing storage objects.
	RegisterBeforeListStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListStorageObjectsRequest) (*api.ListStorageObjectsRequest, error)) error

	// RegisterAfterListStorageObjects can be used to perform additional logic after listing storage objects.
	RegisterAfterListStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.StorageObjectList, in *api.ListStorageObjectsRequest) error) error

	// RegisterBeforeReadStorageObjects can be used to perform additional logic before reading storage objects.
	RegisterBeforeReadStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ReadStorageObjectsRequest) (*api.ReadStorageObjectsRequest, error)) error

	// RegisterAfterReadStorageObjects can be used to perform additional logic after reading storage objects.
	RegisterAfterReadStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.StorageObjects, in *api.ReadStorageObjectsRequest) error) error

	// RegisterBeforeWriteStorageObjects can be used to perform additional logic before writing storage objects.
	RegisterBeforeWriteStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.WriteStorageObjectsRequest) (*api.WriteStorageObjectsRequest, error)) error

	// RegisterAfterWriteStorageObjects can be used to perform additional logic after writing storage objects.
	RegisterAfterWriteStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.StorageObjectAcks, in *api.WriteStorageObjectsRequest) error) error

	// RegisterBeforeDeleteStorageObjects can be used to perform additional logic before deleting storage objects.
	RegisterBeforeDeleteStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteStorageObjectsRequest) (*api.DeleteStorageObjectsRequest, error)) error

	// RegisterAfterDeleteStorageObjects can be used to perform additional logic after deleting storage objects.
	RegisterAfterDeleteStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.DeleteStorageObjectsRequest) error) error

	// RegisterBeforeJoinTournament can be used to perform additional logic before user joins a tournament.
	RegisterBeforeJoinTournament(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.JoinTournamentRequest) (*api.JoinTournamentRequest, error)) error

	// RegisterAfterJoinTournament can be used to perform additional logic after user joins a tournament.
	RegisterAfterJoinTournament(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.JoinTournamentRequest) error) error

	// RegisterBeforeListTournamentRecords can be used to perform additional logic before listing tournament records.
	RegisterBeforeListTournamentRecords(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListTournamentRecordsRequest) (*api.ListTournamentRecordsRequest, error)) error

	// RegisterAfterListTournamentRecords can be used to perform additional logic after listing tournament records.
	RegisterAfterListTournamentRecords(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsRequest) error) error

	// RegisterBeforeListTournaments can be used to perform additional logic before listing tournaments.
	RegisterBeforeListTournaments(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListTournamentsRequest) (*api.ListTournamentsRequest, error)) error

	// RegisterAfterListTournaments can be used to perform additional logic after listing tournaments.
	RegisterAfterListTournaments(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.TournamentList, in *api.ListTournamentsRequest) error) error

	// RegisterBeforeWriteTournamentRecord can be used to perform additional logic before writing tournament records.
	RegisterBeforeWriteTournamentRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.WriteTournamentRecordRequest) (*api.WriteTournamentRecordRequest, error)) error

	// RegisterAfterWriteTournamentRecord can be used to perform additional logic after writing tournament records.
	RegisterAfterWriteTournamentRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.LeaderboardRecord, in *api.WriteTournamentRecordRequest) error) error

	// RegisterBeforeListTournamentRecordsAroundOwner can be used to perform additional logic before listing tournament records.
	RegisterBeforeListTournamentRecordsAroundOwner(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListTournamentRecordsAroundOwnerRequest) (*api.ListTournamentRecordsAroundOwnerRequest, error)) error

	// RegisterAfterListTournamentRecordsAroundOwner can be used to perform additional logic after listing tournament records.
	RegisterAfterListTournamentRecordsAroundOwner(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsAroundOwnerRequest) error) error

	// RegisterBeforeValidatePurchaseApple can be used to perform additional logic before validating an Apple Store IAP receipt.
	RegisterBeforeValidatePurchaseApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ValidatePurchaseAppleRequest) (*api.ValidatePurchaseAppleRequest, error)) error

	// RegisterAfterValidatePurchaseApple can be used to perform additional logic after validating an Apple Store IAP receipt.
	RegisterAfterValidatePurchaseApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseAppleRequest) error) error

	// RegisterBeforeValidateSubscriptionApple can be used to perform additional logic before validation an Apple Store Subscription receipt.
	RegisterBeforeValidateSubscriptionApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ValidateSubscriptionAppleRequest) (*api.ValidateSubscriptionAppleRequest, error)) error

	// RegisterAfterValidateSubscriptionApple can be used to perform additional logic after validation an Apple Store Subscription receipt.
	RegisterAfterValidateSubscriptionApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.ValidateSubscriptionResponse, in *api.ValidateSubscriptionAppleRequest) error) error

	// RegisterBeforeValidatePurchaseGoogle can be used to perform additional logic before validating a Google Play Store IAP receipt.
	RegisterBeforeValidatePurchaseGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ValidatePurchaseGoogleRequest) (*api.ValidatePurchaseGoogleRequest, error)) error

	// RegisterAfterValidatePurchaseGoogle can be used to perform additional logic after validating a Google Play Store IAP receipt.
	RegisterAfterValidatePurchaseGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseGoogleRequest) error) error

	// RegisterBeforeValidateSubscriptionGoogle can be used to perform additional logic before validation an Google Store Subscription receipt.
	RegisterBeforeValidateSubscriptionGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ValidateSubscriptionGoogleRequest) (*api.ValidateSubscriptionGoogleRequest, error)) error

	// RegisterAfterValidateSubscriptionGoogle can be used to perform additional logic after validation an Google Store Subscription receipt.
	RegisterAfterValidateSubscriptionGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.ValidateSubscriptionResponse, in *api.ValidateSubscriptionGoogleRequest) error) error

	// RegisterBeforeValidatePurchaseHuawei can be used to perform additional logic before validating an Huawei App Gallery IAP receipt.
	RegisterBeforeValidatePurchaseHuawei(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ValidatePurchaseHuaweiRequest) (*api.ValidatePurchaseHuaweiRequest, error)) error

	// RegisterAfterValidatePurchaseHuawei can be used to perform additional logic after validating an Huawei App Gallery IAP receipt.
	RegisterAfterValidatePurchaseHuawei(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseHuaweiRequest) error) error

	// RegisterBeforeValidatePurchaseFacebookInstant can be used to perform additional logic before validating an Facebook Instant IAP receipt.
	RegisterBeforeValidatePurchaseFacebookInstant(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ValidatePurchaseFacebookInstantRequest) (*api.ValidatePurchaseFacebookInstantRequest, error)) error

	// RegisterAfterValidatePurchaseFacebookInstant can be used to perform additional logic after validating an Facebook Instant IAP receipt.
	RegisterAfterValidatePurchaseFacebookInstant(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseFacebookInstantRequest) error) error

	// RegisterBeforeListSubscriptions can be used to perform additional logic before listing subscriptions.
	RegisterBeforeListSubscriptions(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.ListSubscriptionsRequest) (*api.ListSubscriptionsRequest, error)) error

	// RegisterAfterListSubscriptions can be used to perform additional logic after listing subscriptions.
	RegisterAfterListSubscriptions(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.SubscriptionList, in *api.ListSubscriptionsRequest) error) error

	// RegisterBeforeGetSubscription can be used to perform additional logic before listing subscriptions.
	RegisterBeforeGetSubscription(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.GetSubscriptionRequest) (*api.GetSubscriptionRequest, error)) error

	// RegisterAfterGetSubscription can be used to perform additional logic after listing subscriptions.
	RegisterAfterGetSubscription(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.ValidatedSubscription, in *api.GetSubscriptionRequest) error) error

	// RegisterBeforeUnlinkApple can be used to perform additional logic before Apple ID is unlinked from an account.
	RegisterBeforeUnlinkApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountApple) (*api.AccountApple, error)) error

	// RegisterAfterUnlinkApple can be used to perform additional logic after Apple ID is unlinked from an account.
	RegisterAfterUnlinkApple(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountApple) error) error

	// RegisterBeforeUnlinkCustom can be used to perform additional logic before custom ID is unlinked from an account.
	RegisterBeforeUnlinkCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountCustom) (*api.AccountCustom, error)) error

	// RegisterAfterUnlinkCustom can be used to perform additional logic after custom ID is unlinked from an account.
	RegisterAfterUnlinkCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountCustom) error) error

	// RegisterBeforeUnlinkDevice can be used to perform additional logic before device ID is unlinked from an account.
	RegisterBeforeUnlinkDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountDevice) (*api.AccountDevice, error)) error

	// RegisterAfterUnlinkDevice can be used to perform additional logic after device ID is unlinked from an account.
	RegisterAfterUnlinkDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountDevice) error) error

	// RegisterBeforeUnlinkEmail can be used to perform additional logic before email is unlinked from an account.
	RegisterBeforeUnlinkEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountEmail) (*api.AccountEmail, error)) error

	// RegisterAfterUnlinkEmail can be used to perform additional logic after email is unlinked from an account.
	RegisterAfterUnlinkEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountEmail) error) error

	// RegisterBeforeUnlinkFacebook can be used to perform additional logic before Facebook is unlinked from an account.
	RegisterBeforeUnlinkFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountFacebook) (*api.AccountFacebook, error)) error

	// RegisterAfterUnlinkFacebook can be used to perform additional logic after Facebook is unlinked from an account.
	RegisterAfterUnlinkFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountFacebook) error) error

	// RegisterBeforeUnlinkFacebookInstantGame can be used to perform additional logic before Facebook Instant Game profile is unlinked from an account.
	RegisterBeforeUnlinkFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error)) error

	// RegisterAfterUnlinkFacebookInstantGame can be used to perform additional logic after Facebook Instant Game profile is unlinked from an account.
	RegisterAfterUnlinkFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountFacebookInstantGame) error) error

	// RegisterBeforeUnlinkGameCenter can be used to perform additional logic before GameCenter is unlinked from an account.
	RegisterBeforeUnlinkGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountGameCenter) (*api.AccountGameCenter, error)) error

	// RegisterAfterUnlinkGameCenter can be used to perform additional logic after GameCenter is unlinked from an account.
	RegisterAfterUnlinkGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountGameCenter) error) error

	// RegisterBeforeUnlinkGoogle can be used to perform additional logic before Google is unlinked from an account.
	RegisterBeforeUnlinkGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountGoogle) (*api.AccountGoogle, error)) error

	// RegisterAfterUnlinkGoogle can be used to perform additional logic after Google is unlinked from an account.
	RegisterAfterUnlinkGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountGoogle) error) error

	// RegisterBeforeUnlinkSteam can be used to perform additional logic before Steam is unlinked from an account.
	RegisterBeforeUnlinkSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountSteam) (*api.AccountSteam, error)) error

	// RegisterAfterUnlinkSteam can be used to perform additional logic after Steam is unlinked from an account.
	RegisterAfterUnlinkSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.AccountSteam) error) error

	// RegisterBeforeGetUsers can be used to perform additional logic before retrieving users.
	RegisterBeforeGetUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, in *api.GetUsersRequest) (*api.GetUsersRequest, error)) error

	// RegisterAfterGetUsers can be used to perform additional logic after retrieving users.
	RegisterAfterGetUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, out *api.Users, in *api.GetUsersRequest) error) error

	// RegisterEvent can be used to define a function handler that triggers when custom events are received or generated.
	RegisterEvent(fn func(ctx context.Context, logger Logger, evt *api.Event)) error

	// RegisterEventSessionStart can be used to define functions triggered when client sessions start.
	RegisterEventSessionStart(fn func(ctx context.Context, logger Logger, evt *api.Event)) error

	// RegisterEventSessionStart can be used to define functions triggered when client sessions end.
	RegisterEventSessionEnd(fn func(ctx context.Context, logger Logger, evt *api.Event)) error

	// Register a new storage index.
	RegisterStorageIndex(name, collection, key string, fields []string, sortableFields []string, maxEntries int, indexOnly bool) error

	// RegisterStorageIndexFilter can be used to define a filtering function for a given storage index.
	RegisterStorageIndexFilter(indexName string, fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, write *StorageWrite) bool) error

	// RegisterFleetManager can be used to register a FleetManager implementation that can be retrieved from the runtime using GetFleetManager().
	RegisterFleetManager(fleetManagerInit FleetManagerInitializer) error

	// RegisterShutdown can be used to register a function that is executed once the server receives a termination signal.
	// This function only fires if shutdown_grace_sec > 0 and will be terminated early if its execution takes longer than the configured grace seconds.
	RegisterShutdown(fn func(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule)) error
}
