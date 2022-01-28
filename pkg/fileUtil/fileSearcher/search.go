package fileSearcher

import (
	"fmt"
	"os"

	"github.com/t-star08/hand/pkg/evaluator"
	"github.com/t-star08/hand/pkg/fileUtil/visitor"
)

var (
	defaultMaxDepth = 5
)

func ForwardSearch(evaluator *evaluator.Evaluator, target, startPath string) (string, error) {
	v := visitor.NewVisitor(evaluator, target, defaultMaxDepth)
	return forwardSearch(v, startPath, 0)
}

func ForwardSearchUntilOrderedDepth(evaluator *evaluator.Evaluator, target, startPath string, maxDepth int) (string, error) {
	v := visitor.NewVisitor(evaluator, target, maxDepth)
	return forwardSearch(v, startPath, 0)
}

func BackwardSearch(evaluator *evaluator.Evaluator, target, startPath string) (string, error) {
	v := visitor.NewVisitor(evaluator, target, defaultMaxDepth)
	return backwardSearch(v, startPath, 0)
}

func BackwardSearchUntilOrderedDepth(evaluator *evaluator.Evaluator, target, startPath string, maxDepth int) (string, error) {
	v := visitor.NewVisitor(evaluator, target, maxDepth)
	return backwardSearch(v, startPath, 0)
}

func forwardSearch(v *visitor.Visitor, path string, depth int) (string, error) {
	if divable, err := v.CheckLanding(path, depth); err != nil {
		return path, err
	} else if !divable {
		return path, fmt.Errorf("not found")
	}

	dirEntries, _ := os.ReadDir(path)
	for _, resource := range dirEntries {
		info, _ := resource.Info()
		if result, err := v.Evaluator.Evaluate(v.Goal, info); err != nil {
			return path, err
		} else if result {
			return path, nil
		}
	}

	for _, dir := range dirEntries {
		if p, err := forwardSearch(v, fmt.Sprintf("%s/%s", path, dir.Name()), depth + 1); err != nil {
			if err.Error() == "not found" {
				continue
			}
			return path, err
		} else {
			return p, nil
		}
	}

	return path, fmt.Errorf("not found")
}

func backwardSearch(v *visitor.Visitor, path string, depth int) (string, error) {
	if divable, err := v.CheckLanding(path, depth); err != nil {
		return path, err
	} else if !divable {
		return path, fmt.Errorf("not found")
	}

	dirEntries, _ := os.ReadDir(path)
	for _, resource := range dirEntries {
		info, _ := resource.Info()
		if result, err := v.Evaluator.Evaluate(v.Goal, info); err != nil {
			return path, err
		} else if result {
			return path, nil
		}
	}

	return backwardSearch(v, "../" + path, depth + 1)
}
