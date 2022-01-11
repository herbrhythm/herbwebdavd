package folders

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/herb-go/util"
	"golang.org/x/net/webdav"
)

var Folders map[string]http.Handler

var FolderNameRegexp = regexp.MustCompile(`^[a-zA-Z0-9\-_\.]{1,64}$`)

func MustCreateFolder(name string, folder string) http.Handler {
	if folder == "" {
		panic(fmt.Errorf("empty folder for `%s`", name))
	}
	stat, err := os.Stat(folder)
	if err != nil {
		panic(fmt.Errorf("read folder `%s` for `%s` fail [%s]", folder, name, err.Error()))
	}
	if !stat.IsDir() {
		panic(fmt.Errorf("folder `%s` for `%s` is not directory", folder, name))
	}
	return &webdav.Handler{
		Prefix:     "/" + name,
		FileSystem: webdav.Dir(folder),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				if os.IsNotExist(err) {
					return
				}
				util.LogError(err)
			}
		},
	}
}
