To Do list for the practice.go app

The plan is to create a programm that i can run quickly in the terminal to reference programming topics in go. This will be very helpful when i code more complex topics and I would need to refer to a topic.

- find out how I can rather divide all of the projects into standalone packages, and then reference them in the practice.go package. 
- thus i would need to run them, but also i would like to display the whole package, so that the user can see the source code with the results.






CHatGPT helped:
To display source code from other Go packages in the terminal from your `main` package (`practice.go`), you need to **read the `.go` files as plain text** using Go's file reading functions.

Here’s how you can do it:

### 1. Example Project Structure

```
/myproject
  |-- main.go
  |-- utils/
        |-- helper.go
```

### 2. `helper.go` (other package)

```go
package utils

import "fmt"

func SayHello() {
    fmt.Println("Hello from utils!")
}
```

### 3. `main.go`

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    data, err := ioutil.ReadFile("utils/helper.go")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(data))
}
```

### Output

When you run `go run main.go`, it will print the source code of `utils/helper.go` in the terminal.

---

Let me know if you want to filter only functions or specific parts of the file.
