package factory

import (
	"myapp/empregado"
)

func Factory(funcionario string) empregado.Empregado {
	switch funcionario {
	case "junior":
		return &empregado.ProgramadorJunior{}
	case "pleno":
		return &empregado.ProgramadorPleno{}
	case "senior":
		return &empregado.ProgramadorSenior{}
	default:
		return nil
	}
}
