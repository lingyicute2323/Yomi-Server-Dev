// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dialog

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/teamgram/proto/mtproto"
)

const (
	dialogKeyPrefix = "dialog.1"
)

const (
	dialogFiltersKeyPrefix = "dialog_filters"
)

const (
	conversationsKeyPrefix = "user_conversations.1"
	chatsKeyPrefix         = "user_chats.1"
	channelsKeyPrefix      = "user_channels.1"
)

var (
	cachePinnedDialogIdListPrefix       = "pinned_dialog_id_list.1"
	cacheFolderPinnedDialogIdListPrefix = "folder_pinned_dialog_id_list.1"
	cacheDialogIdListPrefix             = "dialog_id_list.1"
	cacheFolderDialogIdListPrefix       = "folder_dialog_id_list.1"
)

func GetCacheOneKey(prefix string, k int64) string {
	return prefix + "#" + strconv.FormatInt(k, 10)
}

func GetCacheTwoKey(prefix string, k1, k2 int64) string {
	return prefix + "#" + strconv.FormatInt(k1, 10) + "_" + strconv.FormatInt(k2, 10)
}

func GetCacheThreeKey(prefix string, k1, k2, k3 int64) string {
	return prefix + "#" + strconv.FormatInt(k1, 10) + "_" + strconv.FormatInt(k2, 10) + "_" + strconv.FormatInt(k3, 10)
}

func ParseCacheOneKey(k, prefix string) int64 {
	if strings.HasPrefix(k, prefix+"#") {
		v, _ := strconv.ParseInt(k[len(prefix)+1:], 10, 64)
		return v
	}

	return 0
}

func ParseCacheTwoKey(k, prefix string) (int64, int64) {
	if strings.HasPrefix(k, prefix+"#") {
		v := strings.Split(k[len(prefix)+1:], "_")
		if len(v) != 2 {
			return 0, 0
		}
		v1, _ := strconv.ParseInt(v[0], 10, 64)
		v2, _ := strconv.ParseInt(v[1], 10, 64)

		return v1, v2
	}

	return 0, 0
}

func ParseCacheThreeKey(k, prefix string) (int64, int64, int64) {
	if strings.HasPrefix(k, prefix+"#") {
		v := strings.Split(k[len(prefix)+1:], "_")
		if len(v) != 3 {
			return 0, 0, 0
		}
		v1, _ := strconv.ParseInt(v[0], 10, 64)
		v2, _ := strconv.ParseInt(v[1], 10, 64)
		v3, _ := strconv.ParseInt(v[2], 10, 64)

		return v1, v2, v3
	}

	return 0, 0, 0
}

func GetPinnedDialogIdListCacheKey(userId int64) string {
	return GetCacheOneKey(cachePinnedDialogIdListPrefix, userId)
}
func GetNotPinnedDialogIdListCacheKey(userId int64) string {
	return GetCacheOneKey(cacheDialogIdListPrefix, userId)
}

func GetFolderPinnedDialogIdListCacheKey(userId int64) string {
	return GetCacheOneKey(cacheFolderPinnedDialogIdListPrefix, userId)
}

func GetFolderNotPinnedDialogIdListCacheKey(userId int64) string {
	return GetCacheOneKey(cacheFolderDialogIdListPrefix, userId)
}

func GetDialogCacheKey(userId, peerDialogId int64) string {
	return GetCacheTwoKey(dialogKeyPrefix, userId, peerDialogId)
}

func ParseDialogCacheKey(k string) (int64, int64) {
	return ParseCacheTwoKey(k, dialogKeyPrefix)
}

func GetDialogCacheKeyByPeer(userId int64, peerType int32, peerId int64) string {
	return GetDialogCacheKey(userId, mtproto.MakePeerDialogId(peerType, peerId))
}

func GetDialogFilterCacheKey(userId int64) string {
	return fmt.Sprintf("%s#%d", dialogFiltersKeyPrefix, userId)
}

func GetConversationsCacheKey(userId int64) string {
	return fmt.Sprintf("%s#%d", conversationsKeyPrefix, userId)
}

func GetChatsCacheKey(userId int64) string {
	return fmt.Sprintf("%s#%d", chatsKeyPrefix, userId)
}

func GetChannelsCacheKey(userId int64) string {
	return fmt.Sprintf("%s#%d", channelsKeyPrefix, userId)
}

func GetCacheKeyByPeerType(userId int64, peerType int32) string {
	switch peerType {
	case mtproto.PEER_USER:
		return GetConversationsCacheKey(userId)
	case mtproto.PEER_CHAT:
		return GetChatsCacheKey(userId)
	case mtproto.PEER_CHANNEL:
		return GetChannelsCacheKey(userId)
	}

	return ""
}
