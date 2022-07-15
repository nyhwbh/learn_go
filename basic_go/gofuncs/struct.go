package gofuncs

type Person struct {
	name string
	sex  bool
	age  int
}

//JSON
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}
