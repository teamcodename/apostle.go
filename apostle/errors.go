package apostle

import "net/http"

type NoEmailError struct{}

func (e NoEmailError) Error() string {
	return "apostle.Mail contained no email address"
}

type NoTemplateError struct{}

func (e NoTemplateError) Error() string {
	return "apostle.Mail contained no template id"
}

type NoDomainKeyError struct{}

func (e NoDomainKeyError) Error() string {
	return "No DomainKey is set. Provide one via ENV['APOSTLE_DOMAIN_KEY'], or call apostle.SetDomainKey()"
}

type DeliveryError struct {
	Request  *http.Request
	Response *http.Response
}

func (e DeliveryError) Error() string {
	return "A delivery error occurred"
}

type InvalidDomainKeyError struct {
	Request  *http.Request
	Response *http.Response
}

func (e InvalidDomainKeyError) Error() string {
	return "The domain key was rejected"
}

type UnprocessableEntityError struct {
	Request  *http.Request
	Response *http.Response
}

func (e UnprocessableEntityError) Error() string {
	return "Unprocessable json was supplied"
}

type ServerError struct {
	Request  *http.Request
	Response *http.Response
}

func (e ServerError) Error() string {
	return "A server error occurred"
}
