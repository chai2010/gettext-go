- *赞助 BTC: 1Cbd6oGAUUyBi7X7MaR4np4nTmQZXVgkCW*
- *赞助 ETH: 0x623A3C3a72186A6336C79b18Ac1eD36e1c71A8a6*
- *Go语言付费QQ群: 1055927514*

----

# gettext-go

- PkgDoc: [http://godoc.org/github.com/chai2010/gettext-go](http://godoc.org/github.com/chai2010/gettext-go)
- PkgDoc: [http://pkg.go.dev/github.com/chai2010/gettext-go](http://pkg.go.dev/github.com/chai2010/gettext-go)

## Install

1. `go get github.com/chai2010/gettext-go`
2. `go run hello.go`

The godoc.org or go.dev has more information.

## Example

```Go
package main

import (
	"fmt"

	"github.com/chai2010/gettext-go"
)

func main() {
	gettext.SetLocale("zh_CN")
	gettext.Textdomain("hello")

	gettext.BindTextdomain("hello", "local", nil)

	// gettext.BindTextdomain("hello", "local", nil)         // from local dir
	// gettext.BindTextdomain("hello", "local.zip", nil)     // from local zip file
	// gettext.BindTextdomain("hello", "local.zip", zipData) // from embedded zip data

	// translate source text
	fmt.Println(gettext.Gettext("Hello, world!"))
	// Output: 你好, 世界!

	// if no msgctxt in PO file (only msgid and msgstr),
	// specify context as "" by
	fmt.Println(gettext.PGettext("", "Hello, world!"))
	// Output: 你好, 世界!

	// translate resource
	fmt.Println(string(gettext.Getdata("poems.txt"))))
	// Output: ...
}
```

Go file: [hello.go](https://github.com/chai2010/gettext-go/blob/master/examples/hello.go); PO file: [hello.po](https://github.com/chai2010/gettext-go/blob/master/examples/local/default/LC_MESSAGES/hello.po);

----

## API Changes (v0.1.0 vs v1.0.0)

### Renamed package path

| v0.1.0                                          | v1.0.0                                  |
| ----------------------------------------------- | --------------------------------------- |
| `github.com/chai2010/gettext-go/gettext`        | `github.com/chai2010/gettext-go`        |
| `github.com/chai2010/gettext-go/gettext/po`     | `github.com/chai2010/gettext-go/po`     |
| `github.com/chai2010/gettext-go/gettext/mo`     | `github.com/chai2010/gettext-go/mo`     |
| `github.com/chai2010/gettext-go/gettext/plural` | `github.com/chai2010/gettext-go/plural` |

### Renamed function

| v0.1.0                           | v1.0.0                    |
| -------------------------------- | ------------------------- |
| `gettext-go/gettext.*`           | `gettext-go.*`            |
| `gettext-go/gettext/po.Load`     | `gettext-go/po.LoadFile`  |
| `gettext-go/gettext/po.LoadData` | `gettext-go/po.Load`      |
| `gettext-go/gettext/mo.Load`     | `gettext-go/mo.LoadFile`  |
| `gettext-go/gettext/mo.LoadData` | `gettext-go/mo.Load`      |

### Use empty string as the default context for `gettext.Gettext`

```go
package main

// v0.1.0
// if the **context** missing, use `callerName(2)` as the context:

func main() {
	gettext.Gettext("hello")          
	// => gettext.PGettext("main.main", "hello")

	gettext.DGettext("domain", "hello") 
	// => gettext.DPGettext("domain", "main.main", "hello")

	gettext.NGettext("domain", "hello", "hello2", n)
	// => gettext.PNGettext("domain", "main.main", "hello", "hello2", n)

	gettext.DNGettext("domain", "hello", "hello2", n)
	// => gettext.DPNGettext("domain", "main.main", "hello", "hello2", n)
}
```

### `BindTextdomain` support `FileSystem` interface

```go
// Use FileSystem:
//	BindTextdomain("poedit", "name", OS("path/to/dir")) // bind "poedit" domain
//	BindTextdomain("poedit", "name", OS("path/to.zip")) // bind "poedit" domain
```

### New API in v1.0.0

`FileSystem` interface:

```go
type FileSystem interface {
	LocaleList() []string
	LoadMessagesFile(domain, local, ext string) ([]byte, error)
	LoadResourceFile(domain, local, name string) ([]byte, error)
	String() string
}

func NewFS(name string, x interface{}) FileSystem
func NilFS(name string) FileSystem
func OS(root string) FileSystem
func ZipFS(r *zip.Reader, name string) FileSystem
```

`DomainManager`:

```go
type DomainManager struct {
	// Has unexported fields.
}

func NewDomainManager() *DomainManager
func (p *DomainManager) Bind(domain, path string, data interface{}) (domains, paths []string)
func (p *DomainManager) DGetdata(domain, name string) []byte
func (p *DomainManager) DPNGettext(domain, msgctxt, msgid, msgidPlural string, n int) string
func (p *DomainManager) Getdata(name string) []byte
func (p *DomainManager) PNGettext(msgctxt, msgid, msgidPlural string, n int) string
func (p *DomainManager) SetDomain(domain string) string
func (p *DomainManager) SetLocale(locale string) string
```

And `DefaultManager`:

```go
var DefaultManager = NewDomainManager()
```

----

## BUGS

Please report bugs to <chaishushan@gmail.com>.

Thanks!
