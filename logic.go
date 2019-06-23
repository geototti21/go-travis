package travis

//App for travis
type App struct {
	Name    string
	Version int
}

//New creates a new App
func New(name string, version int)*App{
	return &App{Name: name, Version: version}
}

//NewV1 creates a new v1 App
func NewV1(name string)*App{
	return &App{Name: name, Version: 1}
}