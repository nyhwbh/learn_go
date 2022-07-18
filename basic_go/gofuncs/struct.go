package gofuncs

type Person struct {
	Name string
	Sex  bool
	Age  int
}

//JSON
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}
