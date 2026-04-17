# X

This is an attempt to extend the golang builtin package with more utilities while improving the consistence of its interface.


| **Package**                | **Description**             | **Status**     |
|----------------------------|-----------------------------|----------------|
| [dsx](./dsx/README.md)     | Data Structures             | _experimental_ |
| [fsx](./fsx/README.md)     | File System                 | stable         |
| [httpx](./httpx/README.md) | Networking                  | _experimental_ |
| [iterx](./iterx/README.md) | Iterators                   | _wip_          |
| [jsonx](./jsonx/README.md) | Json                        | stable         |
| [logx](./logx/README.md)   | Logs, prints and formatting | _experimental_ |
| [mathx](./mathx/README.md) | Math for ints and floats    | _experimental_ |
| [randx](./randx/README.md) | Random                      | stable         |
| [runex](./runex/README.md) | Runes                       | stable         |
| [strx](./strx/README.md)   | Strings                     | stable         |
| [syncx](./syncx/README.md) | Sync/Async                  | _wip           |

## Installation

```bash
go get github.com/renatopp/x
```

## Usage

```go
import (
  "github.com/renatopp/x/httpx"
  "github.com/renatopp/x/logx"
)

func main() {
  response := httpx.Fetch("GET", "https://google.com")
  if !response.Is2xx() {
    logx.Println("Bad request!")
    return
  }

  logx.Println("Body: %v", response.Text())
}
```