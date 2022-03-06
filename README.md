# html
A library for HTML written in Go.

# Functions

```go
// IsValidHTMLTag validate whether an input string is a valid HTML Tag
func IsValidHTMLTag(tag string) bool
```

```go
// IsValidHTMLAttribute validate whether an input string is a valid HTML attribute depends on input tag
// Error integer code on false result represents one of reason:
// -2: invalid HTML tag,
// -1: invalid HTML attribute,
//  0: unsupported HTML attribute
func IsValidHTMLAttribute(tag string, attr string) (bool, error)
```

# LICENSE

Read [LICENSE](./LICENSE).