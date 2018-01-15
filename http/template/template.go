// Package template provides template rendering capabilities.
package template

import (
	"fmt"
	"io"
	"net/http"

	"github.com/motki/core/log"
	"github.com/motki/motki-server/http/route"
	"github.com/tyler-sommer/stick"
)

// Config describes configuration options for a template renderer.
type Config struct {
	ViewsPath string `toml:"views_path"`

	BaseURL string `toml:"base_url"`
}

// Params contains named values to pass into a template.
type Params map[string]stick.Value

// A Renderer facilitates rendering of templates.
type Renderer interface {
	// Render executes the named template and outputs to the given io.Writer.
	Render(name string, req *route.Request, w io.Writer, params Params) error

	// Error renders an error page for the given HTTP status code.
	Error(code int, req *route.Request, w http.ResponseWriter)
}

// NewRenderer creates a configured template renderer.
func NewRenderer(c Config, logger log.Logger) *renderer {
	logger.Debugf("template: init with views path: %s", c.ViewsPath)
	logger.Debugf("template: base URL: %s", c.BaseURL)

	env := stick.New(newTemplateLoader(c))

	env.Functions["dump"] = dump
	env.Functions["icon_src"] = iconFile
	env.Functions["portrait_url"] = characterPortraitURL
	env.Functions["corp_logo_url"] = corpLogoURL
	env.Functions["alliance_logo_url"] = allianceLogoURL
	env.Functions["is_currently_on"] = isCurrentlyOn
	env.Functions["is_logged_in"] = isLoggedIn
	env.Functions["user"] = getUser

	// TODO: Move this into funcs.go, make a struct that contains all these functions
	env.Functions["url"] = func(_ stick.Context, args ...stick.Value) stick.Value {
		if len(args) == 0 {
			return c.BaseURL
		}
		return fmt.Sprintf("%s/%s", c.BaseURL, stick.CoerceString(args[0]))
	}

	env.Filters["json_encode"] = jsonEncode
	env.Filters["format"] = format
	env.Filters["money"] = money
	env.Filters["url_encode"] = urlEncode

	return &renderer{env: env, logger: logger}
}

type renderer struct {
	env    *stick.Env
	logger log.Logger
}

func (r *renderer) Render(name string, req *route.Request, w io.Writer, params Params) error {
	if params == nil {
		params = make(Params)
	}
	params["request"] = req
	err := r.env.Execute(name, w, params)
	if err != nil {
		r.logger.Warnf("error rendering template %s: %s", name, err)
	}
	return err
}

func (r *renderer) Error(code int, req *route.Request, w http.ResponseWriter) {
	w.WriteHeader(code)
	r.Render(fmt.Sprintf("error/%v.html.twig", code), req, w, nil)
}
