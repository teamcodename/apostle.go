package apostle

import "net/http"

type NoEmailError struct{}

func (e NoEmailError) Error() string {
	return "apostle.Mail contained no email address"
}

func IsNoEmailError(err error) (ok bool) {
	_, ok = err.(NoEmailError)
	return
}

type NoTemplateError struct{}

func (e NoTemplateError) Error() string {
	return "apostle.Mail contained no template id"
}

func IsNoTemplateError(err error) (ok bool) {
	_, ok = err.(NoTemplateError)
	return
}

type NoDomainKeyError struct{}

func (e NoDomainKeyError) Error() string {
	return "No DomainKey is set. Provide one via ENV['APOSTLE_DOMAIN_KEY'], or call apostle.SetDomainKey()"
}

func IsNoDomainKeyError(err error) (ok bool) {
	_, ok = err.(NoDomainKeyError)
	return
}

type DeliveryError struct {
	Request  *http.Request
	Response *http.Response
}

func (e DeliveryError) Error() string {
	return "A delivery error occurred"
}

func IsDeliveryError(err error) (ok bool) {
	_, ok = err.(InvalidDomainKeyError)
	return
}

type InvalidDomainKeyError struct {
	DeliveryError
}

func (e InvalidDomainKeyError) Error() string {
	return "The domain key was rejected"
}

func IsInvalidDomainKeyError(err error) (ok bool) {
	_, ok = err.(InvalidDomainKeyError)
	return
}

type UnprocessableEntityError struct {
	DeliveryError
}

func (e UnprocessableEntityError) Error() string {
	return "Unprocessable json was supplied"
}

func IsUnprocessableEntityError(err error) (ok bool) {
	_, ok = err.(UnprocessableEntityError)
	return
}

type ServerError struct {
	DeliveryError
}

func (e ServerError) Error() string {
	return "A server error occurred"
}

func IsServerError(err error) (ok bool) {
	_, ok = err.(ServerError)
	return
}
