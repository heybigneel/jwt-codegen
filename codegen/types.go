package codegen

// CustomClaims holds the custom claims
// for the jwt token
type CustomClaims struct {
	// The fields that you want as payload
	Fields []Field `json:"fields"`

	// Minutes after which you want the token
	// to expire
	ExpiresAfter int `json:"expiresAfter"`
}

// Field represents a payload field
// with the name and the type of the field
type Field struct {
	Name     string `json:"name"`
	TypeName string `json:"type"`
}
