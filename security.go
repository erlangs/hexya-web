// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package web

import (
	"github.com/erlangs/okoo/src/models/security"
	"github.com/erlangs/pool/h"
)

func init() {
	h.Filter().Methods().AllowAllToGroup(security.GroupEveryone)
}
