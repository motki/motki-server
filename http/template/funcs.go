package template

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/motki/motki-server/http/auth"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki/eveapi"
	"github.com/motki/motki/evedb"
	"github.com/tyler-sommer/stick"
)

// dump returns a debug view of the given arguments.
func dump(_ stick.Context, args ...stick.Value) stick.Value {
	var vals []interface{}
	for _, arg := range args {
		vals = append(vals, arg)
	}
	return spew.Sdump(vals...)
}

func iconFile(_ stick.Context, args ...stick.Value) stick.Value {
	if len(args) > 0 {
		if ico, ok := args[0].(evedb.Icon); ok {
			// TODO: This is coupled to the web server, move it somewhere else
			if ico.IconFile != "" {
				return strings.Replace(ico.IconFile, "res:/ui/texture/icons", "/images/Icons/items", 1)
			}
		}
	}
	return "/images/Icons/items/7_64_15.png"
}

func characterPortraitURL(_ stick.Context, args ...stick.Value) stick.Value {
	if len(args) < 1 {
		return eveapi.ImageURL(eveapi.ImageCharacterPortrait, 0, 256)
	}
	width := 256
	if len(args) == 2 {
		width = int(stick.CoerceNumber(args[1]))
	}
	id := int(stick.CoerceNumber(args[0]))
	return eveapi.ImageURL(eveapi.ImageCharacterPortrait, id, width)
}

func corpLogoURL(_ stick.Context, args ...stick.Value) stick.Value {
	if len(args) < 1 {
		return eveapi.ImageURL(eveapi.ImageCorpLogo, 0, 128)
	}
	width := 128
	if len(args) == 2 {
		width = int(stick.CoerceNumber(args[1]))
	}
	id := int(stick.CoerceNumber(args[0]))
	return eveapi.ImageURL(eveapi.ImageCorpLogo, id, width)
}

func allianceLogoURL(_ stick.Context, args ...stick.Value) stick.Value {
	if len(args) < 1 {
		return eveapi.ImageURL(eveapi.ImageAllianceLogo, 0, 128)
	}
	width := 128
	if len(args) == 2 {
		width = int(stick.CoerceNumber(args[1]))
	}
	id := int(stick.CoerceNumber(args[0]))
	return eveapi.ImageURL(eveapi.ImageAllianceLogo, id, width)
}

func getRequest(sctx stick.Context) *route.Request {
	c, ok := sctx.Scope().Get("request")
	if !ok {
		return nil
	}
	if r, ok := c.(*route.Request); ok {
		return r
	}
	return nil
}

func isCurrentlyOn(sctx stick.Context, args ...stick.Value) stick.Value {
	if len(args) != 1 {
		return false
	}
	ch, ok := args[0].(string)
	if !ok {
		return false
	}
	req := getRequest(sctx)
	if !ok {
		return false
	}
	if ch == "/" {
		return req.URL.Path == "/"
	}
	return strings.HasPrefix(req.URL.Path, ch)
}

func getUser(sctx stick.Context, args ...stick.Value) stick.Value {
	req := getRequest(sctx)
	au, ok := req.Auth()
	if !ok {
		return nil
	}
	return au.User()
}

func isLoggedIn(sctx stick.Context, args ...stick.Value) stick.Value {
	req := getRequest(sctx)
	_, ok := req.Auth()
	if ok {
		return true
	}
	sess, ok := req.Session()
	if !ok {
		return false
	}
	if _, ok = sess.Get(auth.AuthenticatedUserSessionKey); ok {
		// TODO: this isnt always true
		return true
	}
	return false
}
