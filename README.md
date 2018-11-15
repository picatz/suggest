# suggest
🃏 Simple google search suggestion.

# Install
```bash
$ go get -u github.com/picatz/suggest
```

# Command-Line Usage
Once installed, you will have the `suggest` command-line utility.

> **Example**: You want suggestions for the phrase `install internet explorer`:
```bash
$ suggest install linux
```
Might output the following:
```
install internet explorer 11
install internet explorer on mac
install internet explorer 9
install internet explorer windows 10
install internet explorer on chromebook
install internet explorer for windows 7
install internet explorer 8
install internet explorer 7
install internet explorer 11 on windows 10
install internet explorer 11 for mac
install internet explorer on ubuntu
install internet explorer 9 on windows 10
install internet explorer 10 on windows 7
install internet explorer on linux
install internet explorer 8 on windows 10
install internet explorer on android
install internet explorer wine
install internet explorer 9 on windows 7
install internet explorer for windows 8.1
install internet explorer 6 on windows 10
```

# Library Usage
You can also use `suggest` as a golang package.
```go
package main

import (
    "fmt"

	suggest "github.com/picatz/suggest/core"   
)

func main() {
    query := "apples and oranges"

	suggestions, err := Suggest(query)

	if err != nil {
		panic(err)
	}

	for _, suggestion := range suggestions {
        fmt.Println(suggestion)
    }
}
```
