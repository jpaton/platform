// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package einterfaces

import (
	"github.com/mattermost/platform/model"
)

type SamlInterface interface {
	ConfigureSP() *model.AppError
	BuildRequest() (*model.SamlAuthRequest, *model.AppError)
	DoLogin(encodedXML string) (*model.User, *model.AppError)
}

var theSamlInterface SamlInterface

func RegisterSamlInterface(newInterface SamlInterface) {
	theSamlInterface = newInterface
}

func GetSamlInterface() SamlInterface {
	return theSamlInterface
}
