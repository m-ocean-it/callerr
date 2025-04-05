# `werror`

Examples:
```go
// ./pkg/barer.go

type Barer struct {}
func (b *Barer) Bar() {
    return werror.New("oops")
}

// main.go

func main() {
    err := foo()
    fmt.Println(err) // main.foo: pkg.Barer.Bar: oops
}

func foo() {
    barer := pkg.Barer{}
    return werror.Wrap(barer.Bar())
}

```