package payloads

type Payload interface {
	String() string
}

type PayloadWithOffset interface {
	WithOffset(offset int) Payload
}
