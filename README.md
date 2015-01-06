#GO Dweet.io [![Build Status](https://travis-ci.org/vitorleal/go-dweetio.svg)](https://travis-ci.org/vitorleal/go-dweetio)
Go-dweetio is a simple Go client for [dweet.io](https://dweet.io/) API


##Instaling
```
$ go get github.com/vitorleal/go-dweetio
```

##Using

###GET last dweet for a *thing* [Dweet.io API](https://dweet.io/play/#!/dweets)
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


###GET last dweet for a LOCKED *thing* [Dweet.io API](https://dweet.io/play/#!/dweets)
```go
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myPrivateLockKey"}
	res, err := api.GetLastDweetFor("godweetio") //Change this value with the name of your thing

	if err != nil {
		t.Error("Error in the GetLastDweetFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
```


###GET all dweet for a *thing* - limited to 500 in the [Dweet.io API](https://dweet.io/play/#!/dweets)
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


###GET all dweet for a LOCKED *thing* - limited to 500 in the [Dweet.io API](https://dweet.io/play/#!/dweets)
```go
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myPrivateLockKey"}
	res, err := api.GetDweetsFor("godweetio") //Change this value with the name of your thing

	if err != nil {
		t.Error("Error in the GetDweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
```


###POST a new data for a *thing* [Dweet.io API](https://dweet.io/play/#!/dweets)
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


###POST a new data for a LOCKED *thing* [Dweet.io API](https://dweet.io/play/#!/dweets)
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


###Set alert for a LOCKED *thing* [Dweet.io API](https://dweet.io)
```go
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myLockKey"}
	//First parmeter thing (name of your thing)
	//Second parmeter recipients (emails that will be notifyed separated by comma)
	//Third parmeter condition (A simple javascript expression to evaluate the data in a dweet and to return whether or not an alert should be sent)
	res, err := api.SetAlertFor("godweetio", "vitorleal1@gmail.com", "if(dweet.alertValue > 10) return 'TEST: Greater than 10'; if(dweet.alertValue < 10) return 'TEST: Less than 10';") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
```


###Get alert for a LOCKED *thing* [Dweet.io API](https://dweet.io)
```go
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myLockKey"}
	res, err := api.GetAlertFor("godweetio") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the GetAlertFor")
		fmt.Println(err)
	}

	fmt.Println(res)
}
```


###Remove alert for a LOCKED *thing* [Dweet.io API](https://dweet.io)
```go
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myLockKey"}
	res, err := api.SetAlertFor("godweetio", "vitorleal1@gmail.com", "if(dweet.alertValue > 10) return 'TEST: Greater than 10'; if(dweet.alertValue < 10) return 'TEST: Less than 10';") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
```


##Supporting
* Get last dweet for a **thing**
* Get last dweet for a LOCKED **thing**
* Get all dweets for a **thing**
* Get all dweets for a LOCKED **thing**
* Post new data for a **thing**
* Post new data for a LOCKED **thing**
* Set alert for a LOCKED **thing**
* Get alert for a LOCKED **thing**
* Remove alert for a LOCKED **thing**


##Autor
|[![twitter/vitorleal](http://gravatar.com/avatar/e133221d7fbc0dee159dca127d2f6f00?s=80)](http://twitter.com/vitorleal "Follow @vitorleal on Twitter")|
|---|
|[Vitor Leal](http://vitorleal.com)|

