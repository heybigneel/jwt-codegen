## JWT Code Generator

This project generates JWT (JSON Web Token) code from a JSON config file.

Usage
Create a `config.json` file with your custom claims defined. For example:

```js
{
  "issuer": "Example Inc.",
  "audience": "Example API Users"
}
```

Run `go run main.go`. This will generate a jwt.go file with a JWT struct based on your config.

Use the generated jwt.go file in your code to create JWT tokens with your custom claims. For example:

```go
import "github.com/neel229/jwt-codegen/jwt"

func main() {
  // Create token
  token := jwt.NewJWT()
  token.Issuer = "Example Inc."
  token.Audience = "Example API Users"
  tokenString, err := token.SignedString("secret")
  
  // Parse token
  parsedToken, err := jwt.Parse(tokenString, "secret")
}
```

The generated struct will have getter and setter methods for each claim, as well as SignedString() and Parse() methods to sign and parse the tokens.
