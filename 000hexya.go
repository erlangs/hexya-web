// Copyright 2016 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"github.com/erlangs/okoo/src/i18n/translations"
	// Loading dependencies as blank imports
	_ "github.com/erlangs/hexya-base"
	"github.com/erlangs/hexya-web/controllers"

	"github.com/erlangs/okoo/src/server"
	"github.com/erlangs/okoo/src/tools/logging"
	// Loading controllers package
	_ "github.com/erlangs/hexya-web/controllers"
	"github.com/erlangs/hexya-web/scripts"
)

const (
	// MODULE_NAME is the name of this module.
	// It is used by the framework to recognize a module from a simple go package.
	MODULE_NAME = "web"
)

var log logging.Logger

func init() {
	log = logging.GetLogger("web")
	server.RegisterModule(&server.Module{
		Name: MODULE_NAME,
		PreInit: func() {
			controllers.RegisterRoutes()
		},
	})

	translations.RegisterRuleSet(MODULE_NAME, &translations.RuleSet{
		Rules: [][]string{
			{`/static/src/.*\.js`},
			{`/static/src/.*\.xml`},
		},
	})

	translations.RegisterFunc(MODULE_NAME, scripts.UpdateFunc)
}
