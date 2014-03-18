# apostle.go

Go bindings for [Apostle.io](http://apostle.io).


## Installation

```sh
import "github.com/apostle/apostle.go/apostle"
```

## Usage

### Domain Key
You will need to provide your apostle domain key to send emails. You can either place this value into your environment as `APOSTLE_DOMAIN_KEY`, or specify it in your code.

```go
apostle.SetDomainKey("Your Domain Key")
```

### Sending Email

Sending an email is easy, a minimal example may look like this.

```go
mail := apostle.NewMail("welcome_email", "mal@apostle.io")
mail.Deliver()
```

You can pass any information that your Apostle.io template might need to the `Data` attribute. Since these structs are converted using the stdlib `encoding/json` package, only exportable attributes will be encoded. You can change key names by adding `` `json:"key_name"` `` in the type definition.

```go
type Order struct {
	Id 		int		`json:"id"`
	Total 	float64	`json:"total"`
	Items 	int		`json:"items"`
	email	string
}

func MailOrder() {
	o := Order{1234, 12.56, 3, "mal@apostle.io"}
	
	m := apostle.Mail("order_email", o.email)
	m.Data["order"] = o
	m.Deliver()
}
```

### Sending multiple emails

You can send multiple emails at once by using a queue.
```go
q := apostle.NewQueue()

q.Add(apostle.NewMail("welcome_email", "mal@apostle.io"))
q.Add(apostle.NewMail("user_signed_up", "admin@apostle.io"))

q.Deliver()
```

### Errors

#### Validation errors

Both `NewMail` and `Queue.Add` return an error to be checked. It will be one of the following:

* `NoEmailError`: You supplied an invalid email (an empty string)
* `NoTemplateError`: You suppled no template id

You can check the type of error like so:

```go
err := NewMail(someTemplateVar, someEmailVar)
if err != nil{
	if apostle.IsNoEmailError(err) {
		// Email error
	}
	if apostle.IsNoTemplateError(err) {
		// Template error
	}
}
```

#### Delivery Errors


`Mail.Deliver` and `Queue.Deliver` return an error to be checked. It will be one of the following:

* `NoDomainKeyError`: You haven't set a domain key. Either pop your domain key in the `APOSTLE_DOMAIN_KEY` environment variable, or call `apostle.SetDomainKey` with it.
* `InvalidDomainKeyError`: The server rejected the supplied domain key (HTTP 403)
* `UnprocessableEntityError`: The supplied payload was invalid. An invalid payload was supplied, usually a missing email or template id, or no recipients key. `apostle.go` should validate before sending, so it is unlikely you will see this response.
* `ServerError`: (HTTP >= 500) – Server error. Something went wrong at the Apostle API, you should try again with exponential backoff.
* `DeliveryError` – Any response code that is not covered by the above exceptions.

All errors have a relevant checker, e.g. `IsInvalidDomainKeyError`.

All delivery errors, with the exception of `NoDomainKeyError`, have `Request` and `Response` attributes.



## Contributors

* [John Barton](http://whoisjohnbarton.com)
* [Mal Curtis](https://github.com/snikch)
