package event

type Data map[string]interface{}

type Event struct {
        Name string
        Data Data
}
