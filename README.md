Hume API Go Client
==================

This is a Go client for the [Hume API](https://dev.hume.ai/reference).

### DISCLAIMER
This project is in no way affiliated with Hume AI Inc. and is not endorsed by them.
This project is provided as-is, without any warranty or support.

### Installation

```bash
go get github.com/sleepyyui/humetogo
```

### Usage

```go
package main

import (
	"fmt"
	"github.com/sleepyyui/humetogo"
)

func main() {
	api := hume.New("YOUR_API_KEY", "YOUR_API_SECRET")
	tools, err := api.EviGetTools(0, 10, false, "web")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tools)
}
```