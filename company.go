// Copyright 2018 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package web

import (
	"github.com/erlangs/okoo/src/models"
	"github.com/erlangs/okoo/src/models/fields"
	"github.com/erlangs/pool/h"
)

var fields_Company = map[string]models.FieldDefinition{
	"DashboardBackground": fields.Binary{},
}

func init() {
	h.Company().AddFields(fields_Company)
}
