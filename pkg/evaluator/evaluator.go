package evaluator

type Evaluator struct {
	Evaluate func(target, specimen interface{}) (bool, error)
}

func New() *Evaluator {
	return &Evaluator{}
}
