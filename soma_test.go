package soma

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var Resultado int

func facoASomaDe(v1, v2 int) error {
	Resultado = v1 + v2
	return nil
}

func oResultadoDeveSer(soma int) error {

	if Resultado != soma {
		return fmt.Errorf("Erro ao realizar a soma. Esperado: %v (Obtido: %v)", soma, Resultado)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^faco a soma de (\d+) \+ (\d+)$`, facoASomaDe)
	s.Step(`^o resultado deve ser (\d+)$`, oResultadoDeveSer)
}
