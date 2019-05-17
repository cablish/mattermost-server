// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.
	
// Code generated by enterprise/einterfaces_gen.go; DO NOT EDIT.

package einterfaces

import (
	"github.com/hashicorp/memberlist"
	"github.com/mattermost/mattermost-server/einterfaces"
	"github.com/mattermost/mattermost-server/model"
)

//
type ClusterInterface interface {
	GetClusterId() string
	IsLeader() bool
	NotifyJoin(node *memberlist.Node)
	NotifyLeave(node *memberlist.Node)
	NotifyUpdate(node *memberlist.Node)
	NodeMeta(limit int) []byte
	NotifyMsg(buf []byte)
	RegisterClusterMessageHandler(event string, crm einterfaces.ClusterMessageHandler)
	GetBroadcasts(overhead, limit int) [][]byte
	LocalState(join bool) []byte
	MergeRemoteState(buf []byte, join bool)
	StartInterNodeCommunication()
	StopInterNodeCommunication()
	GetMyClusterInfo() *model.ClusterInfo
	GetClusterInfos() []*model.ClusterInfo
	GetClusterStats() ([]*model.ClusterStats, *model.AppError)
	GetLogs(page, perPage int) ([]string, *model.AppError)
	GetPluginStatuses() (model.PluginStatuses, *model.AppError)
	ConfigChanged(previousConfig *model.Config, newConfig *model.Config, sendToOtherServer bool) *model.AppError
	SendClusterMessage(msg *model.ClusterMessage)
}

	