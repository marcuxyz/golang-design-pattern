package empregado

type ProgramadorJunior struct{}

func (pj *ProgramadorJunior) Name() string {
	return "Programador Junior"
}

func (pj *ProgramadorJunior) Salario() float32 {
	return 2.300
}
