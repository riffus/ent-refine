package entrefine

import (
	"embed"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/Masterminds/sprig/v3"
	"github.com/diazoxide/entrefine/common"
	"text/template"
)

var (
	//go:embed templates/*
	_templates embed.FS

	funcMap = template.FuncMap{}
)

type ExtensionOption = func(*Extension) error

// Extension main struct
type Extension struct {
	entc.DefaultExtension
	AppPath             string         // AppPath JS Application path (packages.json directory path)
	GraphqlURL          string         // Graphql URL
	Meta                map[string]any // Meta to share with frontend application
	Prefix              string
	GoJs                GoJSOptions
	ForceGraph2D        ForceGraph2DOptions
	DefaultEdgesDiagram string

	Auth *Auth
}

type GoJSOptions struct {
	Enabled    bool   `json:"Enabled,omitempty"`
	LicenseKey string `json:"LicenseKey,omitempty"`
}

type ForceGraph2DOptions struct {
	Enabled bool `json:"Enabled,omitempty"`
}

// WithAppPath define refine-project directory
func WithAppPath(path string) ExtensionOption {
	return func(ex *Extension) (err error) {
		ex.AppPath = path
		return nil
	}
}

// WithGraphqlURL define graphql server url
func WithGraphqlURL(url string) ExtensionOption {
	return func(ex *Extension) (err error) {
		ex.GraphqlURL = url
		return nil
	}
}

// WithTypeScriptPrefix define typescript types/vars prefix
func WithTypeScriptPrefix(prefix string) ExtensionOption {
	return func(ex *Extension) (err error) {
		ex.Prefix = prefix
		return nil
	}
}

// WithAuth configure authentication and authorization
func WithAuth(options ...AuthOption) ExtensionOption {
	return func(ex *Extension) (err error) {
		ex.Auth = NewAuth(options...)
		return nil
	}
}

// WithDefaultEdgesDiagram set default edges graph/diagram view component name
func WithDefaultEdgesDiagram(name string) ExtensionOption {
	return func(ex *Extension) (err error) {
		ex.DefaultEdgesDiagram = name
		return nil
	}
}

// WithGoJs use gojs for edges diagrams
func WithGoJs(options GoJSOptions) ExtensionOption {
	return func(ex *Extension) (err error) {
		ex.GoJs = options
		return nil
	}
}

// WithForceGraph2D use react-force-graph-2d for edges diagrams
func WithForceGraph2D(options ForceGraph2DOptions) ExtensionOption {
	return func(ex *Extension) (err error) {
		ex.ForceGraph2D = options
		return nil
	}
}

// WithMeta add metadata to `{AppPath}/entrefine.json`
func WithMeta(meta map[string]any) ExtensionOption {
	return func(ex *Extension) (err error) {
		ex.Meta = meta
		return nil
	}
}

// NewExtension initialize extension
func NewExtension(opts ...ExtensionOption) (*Extension, error) {
	ex := &Extension{
		Meta:                map[string]any{},
		Prefix:              "Ent",
		DefaultEdgesDiagram: "Diagram.GoJS",
		GoJs: GoJSOptions{
			Enabled:    true,
			LicenseKey: "test",
		},
	}

	_funcMap := template.FuncMap{
		"ER_label": common.ToLabel,
		"ER_fieldTSType": func(f gen.Field) string {
			return common.FieldTSType(f, ex.Prefix)
		},
		"ER_tsType": func(t string) string {
			return common.TSType(t, ex.Prefix)
		},
		"ER_resourceAlias":   common.ResourceAlias,
		"ER_someField":       someField,
		"ER_titleField":      titleField,
		"ER_someNode":        someNode,
		"ER_mainImageField":  mainImageField,
		"ER_getActionByName": getActionByName,
	}

	if len(funcMap) == 0 {
		for k, v := range _funcMap {
			funcMap[k] = v
		}

		for k, v := range sprig.FuncMap() {
			funcMap[k] = v
		}

		for k, v := range gen.Funcs {
			funcMap[k] = v
		}

		for k, v := range entgql.TemplateFuncs {
			funcMap[k] = v
		}
	}

	for _, opt := range opts {
		if err := opt(ex); err != nil {
			return nil, err
		}
	}
	return ex, nil
}

type Annotations struct {
	Prefix string
	Auth   *Auth
}

// Name of the annotation. Used by the codegen templates.
func (Annotations) Name() string {
	return "ENTREFINE"
}

// Annotations Define Ent annotations
func (ex *Extension) Annotations() []entc.Annotation {
	return []entc.Annotation{
		Annotations{
			Prefix: ex.Prefix,
			Auth:   ex.Auth,
		},
	}
}

// Templates Define Ent templates
func (ex *Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("search_query_apply").
			Funcs(funcMap).
			ParseFS(_templates, "templates/search_query_apply.tmpl")),
		gen.MustParse(gen.NewTemplate("auth").
			Funcs(funcMap).
			ParseFS(_templates, "templates/auth.tmpl")),
	}
}

// Hooks Define Ent hooks
func (ex *Extension) Hooks() []gen.Hook {
	return []gen.Hook{
		GenerateRefineScriptsHook(ex),
		GenerateAuthResourcesHook(ex),
	}
}
