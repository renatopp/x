# X

This is an attempt to extend the golang builtin package with more utilities while improving the consistence of its interface.

TODO:

- strx.Wrap (breaking by word, breaking by letter)
- strx.TableWriter (with styles, with colspan?)


| **Package**                | **Description**             | **Status**     |
|----------------------------|-----------------------------|----------------|
| [fsx](./fsx/README.md)     | File System                 | stable         |
| [jsonx](./jsonx/README.md) | Json                        | stable         |
| [mathx](./mathx/README.md) | Math for ints and floats    | stable         |
| [randx](./randx/README.md) | Random                      | stable         |
| [runex](./runex/README.md) | Runes                       | stable         |
| [strx](./strx/README.md)   | Strings                     | stable         |
| [dsx](./dsx/README.md)     | Data Structures             | _experimental_ |
| [httpx](./httpx/README.md) | Networking                  | _experimental_ |
| [logx](./logx/README.md)   | Logs, prints and formatting | _experimental_ |
| [iterx](./iterx/README.md) | Iterators                   | _wip_          |
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