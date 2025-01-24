```Go
func main() {
    var m map[string]int
    if m == nil {
        fmt.Println("Map is nil. Returning default value.")
        // Or you can initialize the map
        //m = make(map[string]int)
    } else {
        fmt.Println(m["key"])
    }
}
```