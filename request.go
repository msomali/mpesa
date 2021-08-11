package mpesa

type Request struct {
	Method   string
	Type     RequestType
	Endpoint string
	Payload  interface{}
	Headers  map[string]string
}

func (t RequestType) name() string {
	values := map[int]string{
		0: "session key",
		1: "c2b single stage",
	}

	return values[int(t)]
}
