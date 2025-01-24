```Go
func main() {
    var m map[string]int
    fmt.Println(m[invalidKey]) // This will panic if m is nil
}
```