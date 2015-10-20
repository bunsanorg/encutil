package encutil

type Marshaler interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}
