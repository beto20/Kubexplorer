package endpoint

type Params struct {
	Name   string
	Object []Object
}

type Object struct {
	Name        string
	Description string
}

func (p *Params) getParameters() {

}
