package humango

type Rule struct {
	Name string `json:name`
	Context map[string]string `json:context`
}