package routers

import (
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/middleware/router/httprouter"
)

//RouterFactory base router factory.
var RouterFactory = router.NewFactory(func() router.Router {
	var Router = httprouter.New()

	return Router
})
