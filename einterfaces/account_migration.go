// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.
	
// Code generated by enterprise/einterfaces_gen.go; DO NOT EDIT.

package einterfaces

import (
	"github.com/mattermost/mattermost-server/model"
)

//
type AccountMigrationInterface interface {
	MigrateToLdap(fromAuthService string, foreignUserFieldNameToMatch string, force bool, dryRun bool) *model.AppError
	MigrateToSaml(fromAuthService string, usersMap map[string]string, auto bool, dryRun bool) *model.AppError
}
