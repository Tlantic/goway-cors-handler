package goway_cors_handler

import (
	"net/http"
	"github.com/andrepinto/goway/router"
	. "github.com/andrepinto/goway/handlers"
	"bytes"
	"fmt"
)


//noinspection GoUnusedConst
const NAME = "CORS"

func Handle(route *router.Route, req *http.Request) (*HandlerError){
	fmt.Println(req.Method)
	if req.Method == http.MethodOptions {
		req.Write(bytes.NewBufferString("OK"))
		return NewHttpError(http.StatusOK, "OK")
	}
	return nil
}
