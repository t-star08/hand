package fileName

import (
	"os"

	"github.com/t-star08/hand/pkg/evaluator"
	"github.com/t-star08/hand/pkg/evaluator/stringValue"
)

func NewExactEvaluator() *evaluator.Evaluator {
	e := evaluator.New()
	s := stringValue.NewExactEvaluator()

	e.Evaluate = func(target, specimen interface{}) (bool, error) {
		return s.Evaluate(target, specimen.(os.FileInfo).Name())
	}

	return e
}

func NewPartialEvaluator() *evaluator.Evaluator {
	e := evaluator.New()
	s := stringValue.NewPartialEvaluator()

	e.Evaluate = func(target, specimen interface{}) (bool, error) {
		return s.Evaluate(target, specimen.(os.FileInfo).Name())
	}

	return e
}
