package empregado

type ProgramadorSenior struct{}

func (ps *ProgramadorSenior) Name() string {
	return "Programador Senior"
}

func (ps *ProgramadorSenior) Salario() float32 {
	return 8.450
}
