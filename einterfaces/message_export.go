// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.
	
// Code generated by enterprise/einterfaces_gen.go; DO NOT EDIT.

package einterfaces

import (
	"context"

	"github.com/mattermost/mattermost-server/model"
)

//
type MessageExportInterface interface {
	StartSynchronizeJob(ctx context.Context, exportFromTimestamp int64) (*model.Job, *model.AppError)
	RunExport(exportType string, since int64) *model.AppError
}

	