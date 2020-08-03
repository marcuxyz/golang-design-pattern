package empregado

type ProgramadorPleno struct{}

func (pp *ProgramadorPleno) Name() string {
	return "Programador Pleno"
}

func (pp *ProgramadorPleno) Salario() float32 {
	return 3.500
}
