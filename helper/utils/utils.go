package utils

type Utils struct {
	JaegerTracer *JaegerTracer
	ApiKey       *ApiKey
}

func NewUtils() *Utils {
	return &Utils{
		JaegerTracer: NewJaegerTracer(),
		ApiKey:       NewApiKey(),
	}
}
