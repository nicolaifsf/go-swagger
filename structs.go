package swagger

// TODO:
// Operations is partially done, should continue at the responses object

type License struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
}

type ContactInfo struct {
	Name  string `json:"name" yaml:"name"`
	URL   string `json:"url" yaml:"url"`
	Email string `json:"email" yaml:"email"`
}

type Info struct {
	VendorExtensible
	Describable
	Title          string       `json:"title" yaml:"title"`
	TermsOfService string       `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *ContactInfo `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License     `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string       `json:"version" yaml:"version"`
}

type ConsumesProduces struct {
	Consumes []string `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces []string `json:"produces,omitempty" yaml:"produces,omitempty"`
	Schemes  []string `json:"schemes,omitempty" yaml:"schemes,omitempty"` // the scheme, when present must be from [http, https, ws, wss]
}

type Swagger struct {
	ConsumesProduces
	Swagger  string `json:"swagger" yaml:"swagger"`
	Info     Info   `json:"info" yaml:"info"`
	Host     string `json:"host,omitempty" yaml:"host,omitempty"`
	BasePath string `json:"basePath,omitempty" yaml:"basePath,omitempty"` // must start with a leading "/"
	Paths    Paths  `json:"paths" yaml:"paths"`                           // required
}

type Paths struct {
	VendorExtensible
	Paths map[string]PathItem // custom serializer to flatten this, each entry must start with "/"
}

type VendorExtensible struct {
	Extensions map[string]interface{} // custom extensions, omitted when empty
}

type Reference struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type PathItem struct {
	Reference
	VendorExtensible
	Get *Operation `json:"get,omitempty" yaml:"get,omitempty"`
}

type Describable struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type Operation struct {
	Describable
	VendorExtensible
	ConsumesProduces
	Tags         []string               `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	ID           string                 `json:"operationId" yaml:"operationId"`
	Deprecated   bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     map[string][]string    `json:"security,omitempty" yaml:"security,omitempty"`
}

type ExternalDocumentation struct {
	Describable
	URL string `json:"url" yaml:"url"`
}

type SecurityScheme struct {
	Describable
	VendorExtensible
	Type             string            `json:"type" yaml:"type"`
	Name             string            `json:"name,omitempty" yaml:"name,omitempty"`                         // api key
	In               string            `json:"in,omitempty" yaml:"in,omitempty"`                             // api key
	Flow             string            `json:"flow,omitempty" yaml:"flow,omitempty"`                         // oauth2
	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"` // oauth2
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`                 // oauth2
	Scopes           map[string]string `json:"scopes,omitempty" yaml:"scopes,omitempty"`                     // oauth2
}

type Parameter struct {
	Describable
	Items
	VendorExtensible
	Reference
	Type     string  `json:"type,omitempty" yaml:"type,omitempty"`
	Format   string  `json:"format,omitempty" yaml:"format,omitempty"`
	Name     string  `json:"name,omitempty" yaml:"name,omitempty"`
	In       string  `json:"in,omitempty" yaml:"in,omitempty"`
	Required bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Schema   *Schema `json:"schema,omitempty" yaml:"schema,omitempty"` // when in == "body"
}

type Items struct {
	Type             string        `json:"type,omitempty" yaml:"type,omitempty"`
	Format           string        `json:"format,omitempty" yaml:"format,omitempty"`
	Items            *Items        `json:"items,omitempty" yaml:"items,omitempty"`
	CollectionFormat string        `json:"collectionFormat,omitempty" yaml:"collectionFormat,omitempty"`
	Default          interface{}   `json:"default,omitempty" yaml:"default,omitemtpy"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          float64       `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength        int64         `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength        int64         `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern          string        `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems         int64         `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems         int64         `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems      bool          `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MultipleOf       float64       `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Enum             []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
}

type Schema struct {
	Reference
	Describable
	Format           string                 `json:"format" yaml:"format"`
	Title            string                 `json:"title" yaml:"title"`
	Description      string                 `json:"description" yaml:"description"`
	Default          interface{}            `json:"default,omitempty" yaml:"default,omitemtpy"`
	MultipleOf       float64                `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum          float64                `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum bool                   `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          float64                `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum bool                   `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength        int64                  `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength        int64                  `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern          string                 `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems         int64                  `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems         int64                  `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems      bool                   `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MaxProperties    int64                  `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	MinProperties    int64                  `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	Required         bool                   `json:"required,omitempty" yaml:"required,omitempty"`
	Enum             []interface{}          `json:"enum,omitempty" yaml:"enum,omitempty"`
	Type             *StringOrArray         `json:"type,omitempty" yaml:"type,omitempty"`
	Items            *SchemaOrArray         `json:"items,omitempty" yaml:"items,omitempty"`
	AllOf            []Schema               `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	Properties       map[string]Schema      `json:"properties,omitempty" yaml:"properties,omitempty"`
	Discriminator    string                 `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	ReadOnly         bool                   `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	XML              *XMLObject             `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExternalDocs     *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Example          interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
}

type XMLObject struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}

type BoolOrSchema struct {
	Flag   bool
	Schema *Schema
}

type StringOrArray struct {
	Single string
	Multi  []string
}

type SchemaOrArray struct {
	Single *Schema
	Multi  []Schema
}
