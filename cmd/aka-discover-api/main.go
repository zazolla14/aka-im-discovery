package main

import (
	_ "github.com/1nterdigital/aka-im-discover/docs/swag/discover"
	"github.com/1nterdigital/aka-im-discover/pkg/common/cmd"
	"github.com/1nterdigital/aka-im-tools/system/program"
)

// @title						Discover API Documentation
// @version						1.0
// @description					This is the Discover API
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url					http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url					http://www.apache.org/licenses/LICENSE-2.0.html
// @host						stag-v2.akachat.me/im-discover
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						token

func main() {
	if err := cmd.NewDiscoverApiCmd().Exec(); err != nil {
		program.ExitWithError(err)
	}
}
