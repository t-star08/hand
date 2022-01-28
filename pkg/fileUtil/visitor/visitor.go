package visitor

import (
	"os"

	"github.com/t-star08/hand/pkg/evaluator"
)

type Visitor struct {
	Evaluator	*evaluator.Evaluator
	Goal		string
	MaxDepth	int
}

func NewVisitor(e *evaluator.Evaluator, goal string, maxDepth int) *Visitor {
	v := &Visitor{}

	v.Evaluator = e
	v.Goal = goal
	v.MaxDepth = maxDepth
	return v
}

func (v *Visitor) CheckLanding(path string, depth int) (bool, error) {
	if resource, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	} else {
		return resource.IsDir() && depth < v.MaxDepth, nil
	}
}
