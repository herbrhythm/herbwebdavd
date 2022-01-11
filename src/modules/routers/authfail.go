package routers

import (
	"fmt"
	"net/http"

	"github.com/herb-go/herb/middleware/action"
)

const BasicauthRealm = "herbwebdavd"

var ActionAuthfail = action.New(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\", charset=\"UTF-8\"", BasicauthRealm))
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
})
