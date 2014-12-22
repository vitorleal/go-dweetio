#GO Dweetio

Go-dweetio is a simple Go client for [dweetio](https://dweet.io/)

###GET last dweet for a *thing* [Dweet.IO API](https://dweet.io/play/#!/dweets)
```go
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{}
	res, err := api.GetLastDweetFor("godweetio") //Change this value with the name of your thing

	if err != nil {
		t.Error("Error in the GetLastDweetFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
```


###GET all dweet for a *thing* - limited to 500 in the [Dweet.IO API](https://dweet.io/play/#!/dweets)
```go
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{}
	res, err := api.GetDweetsFor("godweetio") //Change this value with the name of your thing

	if err != nil {
		t.Error("Error in the GetDweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
```


###POST a new dweet for a *thing* [Dweet.IO API](https://dweet.io/play/#!/dweets)
```go
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{}
	res, err := api.DweetFor("godweetio", "{ \"lorem\": \"ipsum\" }")

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
```


###POST a new dweet for a LOCKED *thing* [Dweet.IO API](https://dweet.io/play/#!/dweets)
```go
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myPrivateLockKey"}
	res, err := api.DweetFor("godweetio", "{ \"lorem\": \"ipsum\" }")

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
```
