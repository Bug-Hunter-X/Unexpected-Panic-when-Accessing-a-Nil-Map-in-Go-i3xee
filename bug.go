```go
func main() {
    var m map[string]int
    m["a"] = 1 // This will cause a panic if m is nil
    fmt.Println(m["a"])
}
```