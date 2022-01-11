package folders

import (
	"fmt"
	"herbwebdavd/modules/app"
	"net/http"

	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "900systems.folders"

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
		Folders = map[string]http.Handler{}
		for k, v := range app.System.Folders {
			if Folders[k] != nil {
				panic(fmt.Errorf("folder name `%s` registered", k))
			}
			if !FolderNameRegexp.MatchString(k) {
				panic(fmt.Errorf("folder name `%s` not available", k))
			}
			Folders[k] = MustCreateFolder(k, v)
		}
	})
}
