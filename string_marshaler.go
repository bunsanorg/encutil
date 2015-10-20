package encutil

func MarshalToString(m Marshaler) (string, error) {
	data, err := m.Marshal()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func UnmarshalFromString(m Marshaler, data string) error {
	return m.Unmarshal([]byte(data))
}
