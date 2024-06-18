package api

import (
	"reflect"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 128)

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_goTypes = []interface{}{
	(StoreProvider)(0),                               // 0: layerg.api.StoreProvider
	(StoreEnvironment)(0),                            // 1: layerg.api.StoreEnvironment
	(Operator)(0),                                    // 2: layerg.api.Operator
	(Friend_State)(0),                                // 3: layerg.api.Friend.State
	(GroupUserList_GroupUser_State)(0),               // 4: layerg.api.GroupUserList.GroupUser.State
	(UserGroupList_UserGroup_State)(0),               // 5: layerg.api.UserGroupList.UserGroup.State
	(*Account)(nil),                                  // 6: layerg.api.Account
	(*AccountRefresh)(nil),                           // 7: layerg.api.AccountRefresh
	(*AccountApple)(nil),                             // 8: layerg.api.AccountApple
	(*AccountCustom)(nil),                            // 9: layerg.api.AccountCustom
	(*AccountDevice)(nil),                            // 10: layerg.api.AccountDevice
	(*AccountEmail)(nil),                             // 11: layerg.api.AccountEmail
	(*AccountFacebook)(nil),                          // 12: layerg.api.AccountFacebook
	(*AccountFacebookInstantGame)(nil),               // 13: layerg.api.AccountFacebookInstantGame
	(*AccountGameCenter)(nil),                        // 14: layerg.api.AccountGameCenter
	(*AccountGoogle)(nil),                            // 15: layerg.api.AccountGoogle
	(*AccountSteam)(nil),                             // 16: layerg.api.AccountSteam
	(*AddFriendsRequest)(nil),                        // 17: layerg.api.AddFriendsRequest
	(*AddGroupUsersRequest)(nil),                     // 18: layerg.api.AddGroupUsersRequest
	(*SessionRefreshRequest)(nil),                    // 19: layerg.api.SessionRefreshRequest
	(*SessionLogoutRequest)(nil),                     // 20: layerg.api.SessionLogoutRequest
	(*AuthenticateAppleRequest)(nil),                 // 21: layerg.api.AuthenticateAppleRequest
	(*AuthenticateCustomRequest)(nil),                // 22: layerg.api.AuthenticateCustomRequest
	(*AuthenticateDeviceRequest)(nil),                // 23: layerg.api.AuthenticateDeviceRequest
	(*AuthenticateEmailRequest)(nil),                 // 24: layerg.api.AuthenticateEmailRequest
	(*AuthenticateFacebookRequest)(nil),              // 25: layerg.api.AuthenticateFacebookRequest
	(*AuthenticateFacebookInstantGameRequest)(nil),   // 26: layerg.api.AuthenticateFacebookInstantGameRequest
	(*AuthenticateGameCenterRequest)(nil),            // 27: layerg.api.AuthenticateGameCenterRequest
	(*AuthenticateGoogleRequest)(nil),                // 28: layerg.api.AuthenticateGoogleRequest
	(*AuthenticateSteamRequest)(nil),                 // 29: layerg.api.AuthenticateSteamRequest
	(*BanGroupUsersRequest)(nil),                     // 30: layerg.api.BanGroupUsersRequest
	(*BlockFriendsRequest)(nil),                      // 31: layerg.api.BlockFriendsRequest
	(*ChannelMessage)(nil),                           // 32: layerg.api.ChannelMessage
	(*ChannelMessageList)(nil),                       // 33: layerg.api.ChannelMessageList
	(*CreateGroupRequest)(nil),                       // 34: layerg.api.CreateGroupRequest
	(*DeleteFriendsRequest)(nil),                     // 35: layerg.api.DeleteFriendsRequest
	(*DeleteGroupRequest)(nil),                       // 36: layerg.api.DeleteGroupRequest
	(*DeleteLeaderboardRecordRequest)(nil),           // 37: layerg.api.DeleteLeaderboardRecordRequest
	(*DeleteNotificationsRequest)(nil),               // 38: layerg.api.DeleteNotificationsRequest
	(*DeleteTournamentRecordRequest)(nil),            // 39: layerg.api.DeleteTournamentRecordRequest
	(*DeleteStorageObjectId)(nil),                    // 40: layerg.api.DeleteStorageObjectId
	(*DeleteStorageObjectsRequest)(nil),              // 41: layerg.api.DeleteStorageObjectsRequest
	(*Event)(nil),                                    // 42: layerg.api.Event
	(*Friend)(nil),                                   // 43: layerg.api.Friend
	(*FriendList)(nil),                               // 44: layerg.api.FriendList
	(*FriendsOfFriendsList)(nil),                     // 45: layerg.api.FriendsOfFriendsList
	(*GetUsersRequest)(nil),                          // 46: layerg.api.GetUsersRequest
	(*GetSubscriptionRequest)(nil),                   // 47: layerg.api.GetSubscriptionRequest
	(*Group)(nil),                                    // 48: layerg.api.Group
	(*GroupList)(nil),                                // 49: layerg.api.GroupList
	(*GroupUserList)(nil),                            // 50: layerg.api.GroupUserList
	(*ImportFacebookFriendsRequest)(nil),             // 51: layerg.api.ImportFacebookFriendsRequest
	(*ImportSteamFriendsRequest)(nil),                // 52: layerg.api.ImportSteamFriendsRequest
	(*JoinGroupRequest)(nil),                         // 53: layerg.api.JoinGroupRequest
	(*JoinTournamentRequest)(nil),                    // 54: layerg.api.JoinTournamentRequest
	(*KickGroupUsersRequest)(nil),                    // 55: layerg.api.KickGroupUsersRequest
	(*Leaderboard)(nil),                              // 56: layerg.api.Leaderboard
	(*LeaderboardList)(nil),                          // 57: layerg.api.LeaderboardList
	(*LeaderboardRecord)(nil),                        // 58: layerg.api.LeaderboardRecord
	(*LeaderboardRecordList)(nil),                    // 59: layerg.api.LeaderboardRecordList
	(*LeaveGroupRequest)(nil),                        // 60: layerg.api.LeaveGroupRequest
	(*LinkFacebookRequest)(nil),                      // 61: layerg.api.LinkFacebookRequest
	(*LinkSteamRequest)(nil),                         // 62: layerg.api.LinkSteamRequest
	(*ListChannelMessagesRequest)(nil),               // 63: layerg.api.ListChannelMessagesRequest
	(*ListFriendsRequest)(nil),                       // 64: layerg.api.ListFriendsRequest
	(*ListFriendsOfFriendsRequest)(nil),              // 65: layerg.api.ListFriendsOfFriendsRequest
	(*ListGroupsRequest)(nil),                        // 66: layerg.api.ListGroupsRequest
	(*ListGroupUsersRequest)(nil),                    // 67: layerg.api.ListGroupUsersRequest
	(*ListLeaderboardRecordsAroundOwnerRequest)(nil), // 68: layerg.api.ListLeaderboardRecordsAroundOwnerRequest
	(*ListLeaderboardRecordsRequest)(nil),            // 69: layerg.api.ListLeaderboardRecordsRequest
	(*ListMatchesRequest)(nil),                       // 70: layerg.api.ListMatchesRequest
	(*ListNotificationsRequest)(nil),                 // 71: layerg.api.ListNotificationsRequest
	(*ListStorageObjectsRequest)(nil),                // 72: layerg.api.ListStorageObjectsRequest
	(*ListSubscriptionsRequest)(nil),                 // 73: layerg.api.ListSubscriptionsRequest
	(*ListTournamentRecordsAroundOwnerRequest)(nil),  // 74: layerg.api.ListTournamentRecordsAroundOwnerRequest
	(*ListTournamentRecordsRequest)(nil),             // 75: layerg.api.ListTournamentRecordsRequest
	(*ListTournamentsRequest)(nil),                   // 76: layerg.api.ListTournamentsRequest
	(*ListUserGroupsRequest)(nil),                    // 77: layerg.api.ListUserGroupsRequest
	(*Match)(nil),                                    // 78: layerg.api.Match
	(*MatchList)(nil),                                // 79: layerg.api.MatchList
	(*Notification)(nil),                             // 80: layerg.api.Notification
	(*NotificationList)(nil),                         // 81: layerg.api.NotificationList
	(*PromoteGroupUsersRequest)(nil),                 // 82: layerg.api.PromoteGroupUsersRequest
	(*DemoteGroupUsersRequest)(nil),                  // 83: layerg.api.DemoteGroupUsersRequest
	(*ReadStorageObjectId)(nil),                      // 84: layerg.api.ReadStorageObjectId
	(*ReadStorageObjectsRequest)(nil),                // 85: layerg.api.ReadStorageObjectsRequest
	(*Rpc)(nil),                                      // 86: layerg.api.Rpc
	(*Session)(nil),                                  // 87: layerg.api.Session
	(*StorageObject)(nil),                            // 88: layerg.api.StorageObject
	(*StorageObjectAck)(nil),                         // 89: layerg.api.StorageObjectAck
	(*StorageObjectAcks)(nil),                        // 90: layerg.api.StorageObjectAcks
	(*StorageObjects)(nil),                           // 91: layerg.api.StorageObjects
	(*StorageObjectList)(nil),                        // 92: layerg.api.StorageObjectList
	(*Tournament)(nil),                               // 93: layerg.api.Tournament
	(*TournamentList)(nil),                           // 94: layerg.api.TournamentList
	(*TournamentRecordList)(nil),                     // 95: layerg.api.TournamentRecordList
	(*UpdateAccountRequest)(nil),                     // 96: layerg.api.UpdateAccountRequest
	(*UpdateGroupRequest)(nil),                       // 97: layerg.api.UpdateGroupRequest
	(*User)(nil),                                     // 98: layerg.api.User
	(*UserGroupList)(nil),                            // 99: layerg.api.UserGroupList
	(*Users)(nil),                                    // 100: layerg.api.Users
	(*ValidatePurchaseAppleRequest)(nil),             // 101: layerg.api.ValidatePurchaseAppleRequest
	(*ValidateSubscriptionAppleRequest)(nil),         // 102: layerg.api.ValidateSubscriptionAppleRequest
	(*ValidatePurchaseGoogleRequest)(nil),            // 103: layerg.api.ValidatePurchaseGoogleRequest
	(*ValidateSubscriptionGoogleRequest)(nil),        // 104: layerg.api.ValidateSubscriptionGoogleRequest
	(*ValidatePurchaseHuaweiRequest)(nil),            // 105: layerg.api.ValidatePurchaseHuaweiRequest
	(*ValidatePurchaseFacebookInstantRequest)(nil),   // 106: layerg.api.ValidatePurchaseFacebookInstantRequest
	(*ValidatedPurchase)(nil),                        // 107: layerg.api.ValidatedPurchase
	(*ValidatePurchaseResponse)(nil),                 // 108: layerg.api.ValidatePurchaseResponse
	(*ValidateSubscriptionResponse)(nil),             // 109: layerg.api.ValidateSubscriptionResponse
	(*ValidatedSubscription)(nil),                    // 110: layerg.api.ValidatedSubscription
	(*PurchaseList)(nil),                             // 111: layerg.api.PurchaseList
	(*SubscriptionList)(nil),                         // 112: layerg.api.SubscriptionList
	(*WriteLeaderboardRecordRequest)(nil),            // 113: layerg.api.WriteLeaderboardRecordRequest
	(*WriteStorageObject)(nil),                       // 114: layerg.api.WriteStorageObject
	(*WriteStorageObjectsRequest)(nil),               // 115: layerg.api.WriteStorageObjectsRequest
	(*WriteTournamentRecordRequest)(nil),             // 116: layerg.api.WriteTournamentRecordRequest
	nil,                                              // 117: layerg.api.AccountRefresh.VarsEntry
	nil,                                              // 118: layerg.api.AccountApple.VarsEntry
	nil,                                              // 119: layerg.api.AccountCustom.VarsEntry
	nil,                                              // 120: layerg.api.AccountDevice.VarsEntry
	nil,                                              // 121: layerg.api.AccountEmail.VarsEntry
	nil,                                              // 122: layerg.api.AccountFacebook.VarsEntry
	nil,                                              // 123: layerg.api.AccountFacebookInstantGame.VarsEntry
	nil,                                              // 124: layerg.api.AccountGameCenter.VarsEntry
	nil,                                              // 125: layerg.api.AccountGoogle.VarsEntry
	nil,                                              // 126: layerg.api.AccountSteam.VarsEntry
	nil,                                              // 127: layerg.api.SessionRefreshRequest.VarsEntry
	nil,                                              // 128: layerg.api.Event.PropertiesEntry
	(*FriendsOfFriendsList_FriendOfFriend)(nil),      // 129: layerg.api.FriendsOfFriendsList.FriendOfFriend
	(*GroupUserList_GroupUser)(nil),                  // 130: layerg.api.GroupUserList.GroupUser
	(*UserGroupList_UserGroup)(nil),                  // 131: layerg.api.UserGroupList.UserGroup
	(*WriteLeaderboardRecordRequest_LeaderboardRecordWrite)(nil), // 132: layerg.api.WriteLeaderboardRecordRequest.LeaderboardRecordWrite
	(*WriteTournamentRecordRequest_TournamentRecordWrite)(nil),   // 133: layerg.api.WriteTournamentRecordRequest.TournamentRecordWrite
	(*timestamppb.Timestamp)(nil),                                // 134: google.protobuf.Timestamp
	(*wrapperspb.BoolValue)(nil),                                 // 135: google.protobuf.BoolValue
	(*wrapperspb.Int32Value)(nil),                                // 136: google.protobuf.Int32Value
	(*wrapperspb.StringValue)(nil),                               // 137: google.protobuf.StringValue
	(*wrapperspb.UInt32Value)(nil),                               // 138: google.protobuf.UInt32Value
	(*wrapperspb.Int64Value)(nil),                                // 139: google.protobuf.Int64Value
}
var file_api_proto_depIdxs = []int32{
	98,  // 0: layerg.api.Account.user:type_name -> layerg.api.User
	10,  // 1: layerg.api.Account.devices:type_name -> layerg.api.AccountDevice
	134, // 2: layerg.api.Account.verify_time:type_name -> google.protobuf.Timestamp
	134, // 3: layerg.api.Account.disable_time:type_name -> google.protobuf.Timestamp
	117, // 4: layerg.api.AccountRefresh.vars:type_name -> layerg.api.AccountRefresh.VarsEntry
	118, // 5: layerg.api.AccountApple.vars:type_name -> layerg.api.AccountApple.VarsEntry
	119, // 6: layerg.api.AccountCustom.vars:type_name -> layerg.api.AccountCustom.VarsEntry
	120, // 7: layerg.api.AccountDevice.vars:type_name -> layerg.api.AccountDevice.VarsEntry
	121, // 8: layerg.api.AccountEmail.vars:type_name -> layerg.api.AccountEmail.VarsEntry
	122, // 9: layerg.api.AccountFacebook.vars:type_name -> layerg.api.AccountFacebook.VarsEntry
	123, // 10: layerg.api.AccountFacebookInstantGame.vars:type_name -> layerg.api.AccountFacebookInstantGame.VarsEntry
	124, // 11: layerg.api.AccountGameCenter.vars:type_name -> layerg.api.AccountGameCenter.VarsEntry
	125, // 12: layerg.api.AccountGoogle.vars:type_name -> layerg.api.AccountGoogle.VarsEntry
	126, // 13: layerg.api.AccountSteam.vars:type_name -> layerg.api.AccountSteam.VarsEntry
	127, // 14: layerg.api.SessionRefreshRequest.vars:type_name -> layerg.api.SessionRefreshRequest.VarsEntry
	8,   // 15: layerg.api.AuthenticateAppleRequest.account:type_name -> layerg.api.AccountApple
	135, // 16: layerg.api.AuthenticateAppleRequest.create:type_name -> google.protobuf.BoolValue
	9,   // 17: layerg.api.AuthenticateCustomRequest.account:type_name -> layerg.api.AccountCustom
	135, // 18: layerg.api.AuthenticateCustomRequest.create:type_name -> google.protobuf.BoolValue
	10,  // 19: layerg.api.AuthenticateDeviceRequest.account:type_name -> layerg.api.AccountDevice
	135, // 20: layerg.api.AuthenticateDeviceRequest.create:type_name -> google.protobuf.BoolValue
	11,  // 21: layerg.api.AuthenticateEmailRequest.account:type_name -> layerg.api.AccountEmail
	135, // 22: layerg.api.AuthenticateEmailRequest.create:type_name -> google.protobuf.BoolValue
	12,  // 23: layerg.api.AuthenticateFacebookRequest.account:type_name -> layerg.api.AccountFacebook
	135, // 24: layerg.api.AuthenticateFacebookRequest.create:type_name -> google.protobuf.BoolValue
	135, // 25: layerg.api.AuthenticateFacebookRequest.sync:type_name -> google.protobuf.BoolValue
	13,  // 26: layerg.api.AuthenticateFacebookInstantGameRequest.account:type_name -> layerg.api.AccountFacebookInstantGame
	135, // 27: layerg.api.AuthenticateFacebookInstantGameRequest.create:type_name -> google.protobuf.BoolValue
	14,  // 28: layerg.api.AuthenticateGameCenterRequest.account:type_name -> layerg.api.AccountGameCenter
	135, // 29: layerg.api.AuthenticateGameCenterRequest.create:type_name -> google.protobuf.BoolValue
	15,  // 30: layerg.api.AuthenticateGoogleRequest.account:type_name -> layerg.api.AccountGoogle
	135, // 31: layerg.api.AuthenticateGoogleRequest.create:type_name -> google.protobuf.BoolValue
	16,  // 32: layerg.api.AuthenticateSteamRequest.account:type_name -> layerg.api.AccountSteam
	135, // 33: layerg.api.AuthenticateSteamRequest.create:type_name -> google.protobuf.BoolValue
	135, // 34: layerg.api.AuthenticateSteamRequest.sync:type_name -> google.protobuf.BoolValue
	136, // 35: layerg.api.ChannelMessage.code:type_name -> google.protobuf.Int32Value
	134, // 36: layerg.api.ChannelMessage.create_time:type_name -> google.protobuf.Timestamp
	134, // 37: layerg.api.ChannelMessage.update_time:type_name -> google.protobuf.Timestamp
	135, // 38: layerg.api.ChannelMessage.persistent:type_name -> google.protobuf.BoolValue
	32,  // 39: layerg.api.ChannelMessageList.messages:type_name -> layerg.api.ChannelMessage
	40,  // 40: layerg.api.DeleteStorageObjectsRequest.object_ids:type_name -> layerg.api.DeleteStorageObjectId
	128, // 41: layerg.api.Event.properties:type_name -> layerg.api.Event.PropertiesEntry
	134, // 42: layerg.api.Event.timestamp:type_name -> google.protobuf.Timestamp
	98,  // 43: layerg.api.Friend.user:type_name -> layerg.api.User
	136, // 44: layerg.api.Friend.state:type_name -> google.protobuf.Int32Value
	134, // 45: layerg.api.Friend.update_time:type_name -> google.protobuf.Timestamp
	43,  // 46: layerg.api.FriendList.friends:type_name -> layerg.api.Friend
	129, // 47: layerg.api.FriendsOfFriendsList.friends_of_friends:type_name -> layerg.api.FriendsOfFriendsList.FriendOfFriend
	135, // 48: layerg.api.Group.open:type_name -> google.protobuf.BoolValue
	134, // 49: layerg.api.Group.create_time:type_name -> google.protobuf.Timestamp
	134, // 50: layerg.api.Group.update_time:type_name -> google.protobuf.Timestamp
	48,  // 51: layerg.api.GroupList.groups:type_name -> layerg.api.Group
	130, // 52: layerg.api.GroupUserList.group_users:type_name -> layerg.api.GroupUserList.GroupUser
	12,  // 53: layerg.api.ImportFacebookFriendsRequest.account:type_name -> layerg.api.AccountFacebook
	135, // 54: layerg.api.ImportFacebookFriendsRequest.reset:type_name -> google.protobuf.BoolValue
	16,  // 55: layerg.api.ImportSteamFriendsRequest.account:type_name -> layerg.api.AccountSteam
	135, // 56: layerg.api.ImportSteamFriendsRequest.reset:type_name -> google.protobuf.BoolValue
	2,   // 57: layerg.api.Leaderboard.operator:type_name -> layerg.api.Operator
	134, // 58: layerg.api.Leaderboard.create_time:type_name -> google.protobuf.Timestamp
	56,  // 59: layerg.api.LeaderboardList.leaderboards:type_name -> layerg.api.Leaderboard
	137, // 60: layerg.api.LeaderboardRecord.username:type_name -> google.protobuf.StringValue
	134, // 61: layerg.api.LeaderboardRecord.create_time:type_name -> google.protobuf.Timestamp
	134, // 62: layerg.api.LeaderboardRecord.update_time:type_name -> google.protobuf.Timestamp
	134, // 63: layerg.api.LeaderboardRecord.expiry_time:type_name -> google.protobuf.Timestamp
	58,  // 64: layerg.api.LeaderboardRecordList.records:type_name -> layerg.api.LeaderboardRecord
	58,  // 65: layerg.api.LeaderboardRecordList.owner_records:type_name -> layerg.api.LeaderboardRecord
	12,  // 66: layerg.api.LinkFacebookRequest.account:type_name -> layerg.api.AccountFacebook
	135, // 67: layerg.api.LinkFacebookRequest.sync:type_name -> google.protobuf.BoolValue
	16,  // 68: layerg.api.LinkSteamRequest.account:type_name -> layerg.api.AccountSteam
	135, // 69: layerg.api.LinkSteamRequest.sync:type_name -> google.protobuf.BoolValue
	136, // 70: layerg.api.ListChannelMessagesRequest.limit:type_name -> google.protobuf.Int32Value
	135, // 71: layerg.api.ListChannelMessagesRequest.forward:type_name -> google.protobuf.BoolValue
	136, // 72: layerg.api.ListFriendsRequest.limit:type_name -> google.protobuf.Int32Value
	136, // 73: layerg.api.ListFriendsRequest.state:type_name -> google.protobuf.Int32Value
	136, // 74: layerg.api.ListFriendsOfFriendsRequest.limit:type_name -> google.protobuf.Int32Value
	136, // 75: layerg.api.ListGroupsRequest.limit:type_name -> google.protobuf.Int32Value
	136, // 76: layerg.api.ListGroupsRequest.members:type_name -> google.protobuf.Int32Value
	135, // 77: layerg.api.ListGroupsRequest.open:type_name -> google.protobuf.BoolValue
	136, // 78: layerg.api.ListGroupUsersRequest.limit:type_name -> google.protobuf.Int32Value
	136, // 79: layerg.api.ListGroupUsersRequest.state:type_name -> google.protobuf.Int32Value
	138, // 80: layerg.api.ListLeaderboardRecordsAroundOwnerRequest.limit:type_name -> google.protobuf.UInt32Value
	139, // 81: layerg.api.ListLeaderboardRecordsAroundOwnerRequest.expiry:type_name -> google.protobuf.Int64Value
	136, // 82: layerg.api.ListLeaderboardRecordsRequest.limit:type_name -> google.protobuf.Int32Value
	139, // 83: layerg.api.ListLeaderboardRecordsRequest.expiry:type_name -> google.protobuf.Int64Value
	136, // 84: layerg.api.ListMatchesRequest.limit:type_name -> google.protobuf.Int32Value
	135, // 85: layerg.api.ListMatchesRequest.authoritative:type_name -> google.protobuf.BoolValue
	137, // 86: layerg.api.ListMatchesRequest.label:type_name -> google.protobuf.StringValue
	136, // 87: layerg.api.ListMatchesRequest.min_size:type_name -> google.protobuf.Int32Value
	136, // 88: layerg.api.ListMatchesRequest.max_size:type_name -> google.protobuf.Int32Value
	137, // 89: layerg.api.ListMatchesRequest.query:type_name -> google.protobuf.StringValue
	136, // 90: layerg.api.ListNotificationsRequest.limit:type_name -> google.protobuf.Int32Value
	136, // 91: layerg.api.ListStorageObjectsRequest.limit:type_name -> google.protobuf.Int32Value
	136, // 92: layerg.api.ListSubscriptionsRequest.limit:type_name -> google.protobuf.Int32Value
	138, // 93: layerg.api.ListTournamentRecordsAroundOwnerRequest.limit:type_name -> google.protobuf.UInt32Value
	139, // 94: layerg.api.ListTournamentRecordsAroundOwnerRequest.expiry:type_name -> google.protobuf.Int64Value
	136, // 95: layerg.api.ListTournamentRecordsRequest.limit:type_name -> google.protobuf.Int32Value
	139, // 96: layerg.api.ListTournamentRecordsRequest.expiry:type_name -> google.protobuf.Int64Value
	138, // 97: layerg.api.ListTournamentsRequest.category_start:type_name -> google.protobuf.UInt32Value
	138, // 98: layerg.api.ListTournamentsRequest.category_end:type_name -> google.protobuf.UInt32Value
	138, // 99: layerg.api.ListTournamentsRequest.start_time:type_name -> google.protobuf.UInt32Value
	138, // 100: layerg.api.ListTournamentsRequest.end_time:type_name -> google.protobuf.UInt32Value
	136, // 101: layerg.api.ListTournamentsRequest.limit:type_name -> google.protobuf.Int32Value
	136, // 102: layerg.api.ListUserGroupsRequest.limit:type_name -> google.protobuf.Int32Value
	136, // 103: layerg.api.ListUserGroupsRequest.state:type_name -> google.protobuf.Int32Value
	137, // 104: layerg.api.Match.label:type_name -> google.protobuf.StringValue
	78,  // 105: layerg.api.MatchList.matches:type_name -> layerg.api.Match
	134, // 106: layerg.api.Notification.create_time:type_name -> google.protobuf.Timestamp
	80,  // 107: layerg.api.NotificationList.notifications:type_name -> layerg.api.Notification
	84,  // 108: layerg.api.ReadStorageObjectsRequest.object_ids:type_name -> layerg.api.ReadStorageObjectId
	134, // 109: layerg.api.StorageObject.create_time:type_name -> google.protobuf.Timestamp
	134, // 110: layerg.api.StorageObject.update_time:type_name -> google.protobuf.Timestamp
	134, // 111: layerg.api.StorageObjectAck.create_time:type_name -> google.protobuf.Timestamp
	134, // 112: layerg.api.StorageObjectAck.update_time:type_name -> google.protobuf.Timestamp
	89,  // 113: layerg.api.StorageObjectAcks.acks:type_name -> layerg.api.StorageObjectAck
	88,  // 114: layerg.api.StorageObjects.objects:type_name -> layerg.api.StorageObject
	88,  // 115: layerg.api.StorageObjectList.objects:type_name -> layerg.api.StorageObject
	134, // 116: layerg.api.Tournament.create_time:type_name -> google.protobuf.Timestamp
	134, // 117: layerg.api.Tournament.start_time:type_name -> google.protobuf.Timestamp
	134, // 118: layerg.api.Tournament.end_time:type_name -> google.protobuf.Timestamp
	2,   // 119: layerg.api.Tournament.operator:type_name -> layerg.api.Operator
	93,  // 120: layerg.api.TournamentList.tournaments:type_name -> layerg.api.Tournament
	58,  // 121: layerg.api.TournamentRecordList.records:type_name -> layerg.api.LeaderboardRecord
	58,  // 122: layerg.api.TournamentRecordList.owner_records:type_name -> layerg.api.LeaderboardRecord
	137, // 123: layerg.api.UpdateAccountRequest.username:type_name -> google.protobuf.StringValue
	137, // 124: layerg.api.UpdateAccountRequest.display_name:type_name -> google.protobuf.StringValue
	137, // 125: layerg.api.UpdateAccountRequest.avatar_url:type_name -> google.protobuf.StringValue
	137, // 126: layerg.api.UpdateAccountRequest.lang_tag:type_name -> google.protobuf.StringValue
	137, // 127: layerg.api.UpdateAccountRequest.location:type_name -> google.protobuf.StringValue
	137, // 128: layerg.api.UpdateAccountRequest.timezone:type_name -> google.protobuf.StringValue
	137, // 129: layerg.api.UpdateGroupRequest.name:type_name -> google.protobuf.StringValue
	137, // 130: layerg.api.UpdateGroupRequest.description:type_name -> google.protobuf.StringValue
	137, // 131: layerg.api.UpdateGroupRequest.lang_tag:type_name -> google.protobuf.StringValue
	137, // 132: layerg.api.UpdateGroupRequest.avatar_url:type_name -> google.protobuf.StringValue
	135, // 133: layerg.api.UpdateGroupRequest.open:type_name -> google.protobuf.BoolValue
	134, // 134: layerg.api.User.create_time:type_name -> google.protobuf.Timestamp
	134, // 135: layerg.api.User.update_time:type_name -> google.protobuf.Timestamp
	131, // 136: layerg.api.UserGroupList.user_groups:type_name -> layerg.api.UserGroupList.UserGroup
	98,  // 137: layerg.api.Users.users:type_name -> layerg.api.User
	135, // 138: layerg.api.ValidatePurchaseAppleRequest.persist:type_name -> google.protobuf.BoolValue
	135, // 139: layerg.api.ValidateSubscriptionAppleRequest.persist:type_name -> google.protobuf.BoolValue
	135, // 140: layerg.api.ValidatePurchaseGoogleRequest.persist:type_name -> google.protobuf.BoolValue
	135, // 141: layerg.api.ValidateSubscriptionGoogleRequest.persist:type_name -> google.protobuf.BoolValue
	135, // 142: layerg.api.ValidatePurchaseHuaweiRequest.persist:type_name -> google.protobuf.BoolValue
	135, // 143: layerg.api.ValidatePurchaseFacebookInstantRequest.persist:type_name -> google.protobuf.BoolValue
	0,   // 144: layerg.api.ValidatedPurchase.store:type_name -> layerg.api.StoreProvider
	134, // 145: layerg.api.ValidatedPurchase.purchase_time:type_name -> google.protobuf.Timestamp
	134, // 146: layerg.api.ValidatedPurchase.create_time:type_name -> google.protobuf.Timestamp
	134, // 147: layerg.api.ValidatedPurchase.update_time:type_name -> google.protobuf.Timestamp
	134, // 148: layerg.api.ValidatedPurchase.refund_time:type_name -> google.protobuf.Timestamp
	1,   // 149: layerg.api.ValidatedPurchase.environment:type_name -> layerg.api.StoreEnvironment
	107, // 150: layerg.api.ValidatePurchaseResponse.validated_purchases:type_name -> layerg.api.ValidatedPurchase
	110, // 151: layerg.api.ValidateSubscriptionResponse.validated_subscription:type_name -> layerg.api.ValidatedSubscription
	0,   // 152: layerg.api.ValidatedSubscription.store:type_name -> layerg.api.StoreProvider
	134, // 153: layerg.api.ValidatedSubscription.purchase_time:type_name -> google.protobuf.Timestamp
	134, // 154: layerg.api.ValidatedSubscription.create_time:type_name -> google.protobuf.Timestamp
	134, // 155: layerg.api.ValidatedSubscription.update_time:type_name -> google.protobuf.Timestamp
	1,   // 156: layerg.api.ValidatedSubscription.environment:type_name -> layerg.api.StoreEnvironment
	134, // 157: layerg.api.ValidatedSubscription.expiry_time:type_name -> google.protobuf.Timestamp
	134, // 158: layerg.api.ValidatedSubscription.refund_time:type_name -> google.protobuf.Timestamp
	107, // 159: layerg.api.PurchaseList.validated_purchases:type_name -> layerg.api.ValidatedPurchase
	110, // 160: layerg.api.SubscriptionList.validated_subscriptions:type_name -> layerg.api.ValidatedSubscription
	132, // 161: layerg.api.WriteLeaderboardRecordRequest.record:type_name -> layerg.api.WriteLeaderboardRecordRequest.LeaderboardRecordWrite
	136, // 162: layerg.api.WriteStorageObject.permission_read:type_name -> google.protobuf.Int32Value
	136, // 163: layerg.api.WriteStorageObject.permission_write:type_name -> google.protobuf.Int32Value
	114, // 164: layerg.api.WriteStorageObjectsRequest.objects:type_name -> layerg.api.WriteStorageObject
	133, // 165: layerg.api.WriteTournamentRecordRequest.record:type_name -> layerg.api.WriteTournamentRecordRequest.TournamentRecordWrite
	98,  // 166: layerg.api.FriendsOfFriendsList.FriendOfFriend.user:type_name -> layerg.api.User
	98,  // 167: layerg.api.GroupUserList.GroupUser.user:type_name -> layerg.api.User
	136, // 168: layerg.api.GroupUserList.GroupUser.state:type_name -> google.protobuf.Int32Value
	48,  // 169: layerg.api.UserGroupList.UserGroup.group:type_name -> layerg.api.Group
	136, // 170: layerg.api.UserGroupList.UserGroup.state:type_name -> google.protobuf.Int32Value
	2,   // 171: layerg.api.WriteLeaderboardRecordRequest.LeaderboardRecordWrite.operator:type_name -> layerg.api.Operator
	2,   // 172: layerg.api.WriteTournamentRecordRequest.TournamentRecordWrite.operator:type_name -> layerg.api.Operator
	173, // [173:173] is the sub-list for method output_type
	173, // [173:173] is the sub-list for method input_type
	173, // [173:173] is the sub-list for extension type_name
	173, // [173:173] is the sub-list for extension extendee
	0,   // [0:173] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountRefresh); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountApple); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountCustom); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountEmail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountFacebook); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountFacebookInstantGame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountGameCenter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountGoogle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountSteam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGroupUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionRefreshRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionLogoutRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateAppleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateCustomRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateDeviceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateEmailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateFacebookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateFacebookInstantGameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateGameCenterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateGoogleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateSteamRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BanGroupUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockFriendsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[27].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelMessageList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[28].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[29].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFriendsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[30].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[31].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLeaderboardRecordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[32].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNotificationsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[33].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTournamentRecordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[34].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStorageObjectId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[35].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStorageObjectsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[36].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[37].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Friend); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[38].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[39].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendsOfFriendsList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[40].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[41].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[42].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[43].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[44].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupUserList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[45].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportFacebookFriendsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[46].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportSteamFriendsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[47].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[48].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinTournamentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[49].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickGroupUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[50].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Leaderboard); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[51].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[52].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[53].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardRecordList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[54].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[55].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkFacebookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[56].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkSteamRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[57].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChannelMessagesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[58].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[59].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendsOfFriendsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[60].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[61].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[62].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLeaderboardRecordsAroundOwnerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[63].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLeaderboardRecordsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[64].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMatchesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[65].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNotificationsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[66].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStorageObjectsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[67].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[68].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTournamentRecordsAroundOwnerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[69].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTournamentRecordsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[70].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTournamentsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[71].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserGroupsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[72].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[73].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[74].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[75].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[76].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromoteGroupUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[77].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoteGroupUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[78].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadStorageObjectId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[79].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadStorageObjectsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[80].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rpc); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[81].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[82].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageObject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[83].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageObjectAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[84].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageObjectAcks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[85].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageObjects); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[86].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageObjectList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[87].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tournament); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[88].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TournamentList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[89].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TournamentRecordList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[90].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[91].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[92].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[93].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[94].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[95].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePurchaseAppleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[96].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateSubscriptionAppleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[97].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePurchaseGoogleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[98].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateSubscriptionGoogleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[99].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePurchaseHuaweiRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[100].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePurchaseFacebookInstantRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[101].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatedPurchase); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[102].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePurchaseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[103].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateSubscriptionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[104].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatedSubscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[105].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[106].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[107].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteLeaderboardRecordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[108].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteStorageObject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[109].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteStorageObjectsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[110].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteTournamentRecordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[123].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendsOfFriendsList_FriendOfFriend); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[124].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupUserList_GroupUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[125].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupList_UserGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[126].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteLeaderboardRecordRequest_LeaderboardRecordWrite); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[127].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteTournamentRecordRequest_TournamentRecordWrite); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   128,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
