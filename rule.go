package humango

type Rule struct {
	Name    string                 `json:"name"`
	Context map[string]interface{} `json:"context"`
}
