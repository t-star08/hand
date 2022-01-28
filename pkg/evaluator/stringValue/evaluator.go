package stringValue

import (
	"regexp"

	"github.com/t-star08/hand/pkg/evaluator"
)

func NewExactEvaluator() *evaluator.Evaluator {
	e := evaluator.New()

	e.Evaluate = func(target, specimen interface{}) (bool, error) {
		return target == specimen, nil
	}

	return e
}

func NewPartialEvaluator() *evaluator.Evaluator {
	e := evaluator.New()

	e.Evaluate = func(target, specimen interface{}) (bool, error) {
		if re, err := regexp.Compile(target.(string)); err != nil {
			return false, err
		} else {
			return re.MatchString(specimen.(string)), nil
		}
	}

	return e
}
