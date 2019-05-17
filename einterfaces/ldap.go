// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.
	
// Code generated by enterprise/einterfaces_gen.go; DO NOT EDIT.

package einterfaces

import (
	"github.com/mattermost/mattermost-server/model"
)

//
type LdapInterface interface {
	GetUser(loginIdAttribute string) (*model.User, *model.AppError)
	GetUserAttributes(ldapId string, attributes []string) (map[string]string, *model.AppError)
	CheckPassword(loginIdAttribute string, password string) *model.AppError
	CheckPasswordAuthData(idAttribute string, password string) *model.AppError
	DoLogin(idAttribute string, password string) (*model.User, *model.AppError)
	SwitchToLdap(userId, ldapLoginId, ldapPassword string) *model.AppError
	StartSynchronizeJob(waitForJobToFinish bool) (*model.Job, *model.AppError)
	RunTest() *model.AppError
	GetAllLdapUsers() ([]*model.User, *model.AppError)
	MigrateIDAttribute(toAttribute string) error
	// GetAllGroupsPage retrieves all visible LDAP groups with the associated Mattermost group ID, if available.
	GetAllGroupsPage(page int, perPage int, opts model.LdapGroupSearchOpts) ([]*model.Group, int, *model.AppError)
	// GetGroup retrieves a single group by ldap group ID (RemoteId).
	GetGroup(ldapGroupID string) (*model.Group, *model.AppError)
	// LinkedReachableGroups retrieves all of the LDAP-linked groups that are reachable in the LDAP groups graph paths
	// starting from all direct group memberships of the given user.
	//
	// For example given these pseudo LDAP entries:
	//
	// 	dn: Developers
	// 	objectClass: group
	// 	member: Backend
	// 	member: Frontend
	// 	member: Joe.Intern
	//
	// 	dn: Backend
	// 	objectClass: group
	// 	member: Jane.Backend
	//
	// 	dn: Frontend
	// 	objectClass: group
	// 	member: Gary.Frontend
	//
	//
	// 	LinkedReachableGroups("Joe.Intern")
	//	// [Developers]
	//
	// 	LinkedReachableGroups("Jane.Backend")
	//	// [Developers, Backend]
	//
	// 	LinkedReachableGroups("Gary.Frontend")
	//	// [Developers, Frontend]
	LinkedReachableGroups(userAuthData string) ([]*model.Group, *model.AppError)
	FirstLoginSync(userId, userAuthService, userAuthData string) *model.AppError
}
