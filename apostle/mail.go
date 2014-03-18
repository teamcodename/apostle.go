package apostle

type Mail struct {
	Data       map[string]interface{} `json:"data"`
	Email      string                 `json:"-"`
	From       string                 `json:"from"`
	Headers    map[string][]string    `json:"headers"`
	LayoutId   string                 `json:"layout_id"`
	Name       string                 `json:"name"`
	ReplyTo    string                 `json:"reply_to"`
	TemplateId string                 `json:"template_id"`
}

func NewMail(t string, e string) (m Mail, err error) {
	if len(t) == 0 {
		err = NoTemplateError{}
		return
	}
	if len(e) == 0 {
		err = NoEmailError{}
		return
	}
	return Mail{
		TemplateId: t,
		Email:      e,
		Headers:    make(map[string][]string),
		Data:       make(map[string]interface{}),
	}, nil
}

func (m *Mail) Deliver() error {
	q := NewQueue()
	q.Add(*m)
	return q.Deliver()
}
