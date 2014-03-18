package apostle

import "os"

type config struct {
	DomainKey    string
	DeliveryHost string
}

func loadConfig() (c *config) {
	c = &config{}
	if deliveryHost := os.Getenv("APOSTLE_DELIVERY_HOST"); len(deliveryHost) != 0 {
		c.DeliveryHost = deliveryHost
	} else {
		c.DeliveryHost = "https://deliver.apostle.io"
	}

	if domainKey := os.Getenv("APOSTLE_DOMAIN_KEY"); len(domainKey) != 0 {
		c.DomainKey = domainKey
	}
	return
}

func SetDomainKey(dk string) {
	conf.DomainKey = dk
}
func SetDeliveryHost(dh string) {
	conf.DeliveryHost = dh
}
