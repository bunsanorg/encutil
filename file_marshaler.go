package encutil

import "io/ioutil"

func MarshalToFile(m Marshaler, filename string) error {
    data, err := m.Marshal()
    if err != nil {
        return err
    }
    return ioutil.WriteFile(filename, data, 0666)
}

func UnmarshalFromFile(m Marshaler, filename string) error {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }
    return m.Unmarshal(data)
}
