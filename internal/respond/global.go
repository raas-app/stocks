package respond

var (
	globalResponder = &Responder{}
)

func NewResponse() *Responder { return globalResponder }
