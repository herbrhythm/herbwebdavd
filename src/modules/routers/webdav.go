package routers

import (
	"herbwebdavd/modules/systems/folders"
	"net/http"
	"strings"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsecurity/authorize/role"
	"github.com/herb-go/herbsecurity/authorize/role/roleparser"

	"github.com/herb-go/herbmodules/protecter"

	"github.com/herb-go/herb/middleware/action"
)

var Webdav = action.New(func(w http.ResponseWriter, r *http.Request) {
	pathlist := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
	folder := pathlist[0]
	if folder == "" {
		http.NotFound(w, r)
		return
	}

	auth := protecter.LoadAuth(r)
	userroles, err := roleparser.Parse(auth.Payloads().LoadString(authority.PayloadRoles))
	if err != nil {
		panic(err)
	}
	folderauth := role.NewRoles(role.NewRole("folder").WithNewAttribute("name", []byte(folder)))
	required := role.Any(role.Superuser, folderauth)

	ok, err := required.Authorize(userroles)
	if err != nil {
		panic(err)
	}
	if !ok {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	h := folders.Folders[folder]
	if h == nil {
		http.NotFound(w, r)
		return
	}
	h.ServeHTTP(w, r)
})
