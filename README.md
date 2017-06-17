# Clipboard replacer
## Modify the clipboard using regular expresions

```go
package main

import clipboard "github.com/alepacheco/Clipboard-Replacer"

func main() {
	clipboard.AddRegEx("John", "Davis")
	clipboard.AddRegEx("p([a-z]+)ch", "something")
	clipboard.Watch(1) // Check clipboard every 1 second. (Blocking call)
}
```
