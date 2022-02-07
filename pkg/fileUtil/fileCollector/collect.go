package fileCollector

import (
	"fmt"
	"os"

	"github.com/t-star08/hand/pkg/evaluator"
	"github.com/t-star08/hand/pkg/fileUtil/visitor"
)

var (
	defaultMaxDepth = 5
)

func ForwardCollect(evaluator *evaluator.Evaluator, target, startPath string) ([]string, error) {
	v := visitor.NewVisitor(evaluator, target, defaultMaxDepth)
	return forwardCollect(v, make([]string, 0), startPath, 0)
}

func ForwardCollectUntilOrderedDepth(evaluator *evaluator.Evaluator, target, startPath string, maxDepth int) ([]string, error) {
	v := visitor.NewVisitor(evaluator, target, maxDepth)
	return forwardCollect(v, make([]string, 0), startPath, 0)
}

func BackwardCollect(evaluator *evaluator.Evaluator, target, startPath string) ([]string, error) {
	v := visitor.NewVisitor(evaluator, target, defaultMaxDepth)
	return backwardCollect(v, make([]string, 0), startPath, 0)
}

func BackwardCollectUntilOrderedDepth(evaluator *evaluator.Evaluator, target, startPath string, maxDepth int) ([]string, error) {
	v := visitor.NewVisitor(evaluator, target, maxDepth)
	return backwardCollect(v, make([]string, 0), startPath, 0)
}

func forwardCollect(v *visitor.Visitor, collectibles []string, path string, depth int) ([]string, error) {
	if divable, err := v.CheckLanding(path, depth); err != nil {
		return collectibles, err
	} else if !divable {
		return collectibles, fmt.Errorf("not found")
	}

	beforeCollected := len(collectibles)
	dirEntries, _ := os.ReadDir(path)
	for _, resource := range dirEntries {
		info, _ := resource.Info()
		if result, err := v.Evaluator.Evaluate(v.Goal, info); err != nil {
			return collectibles, err
		} else if result {
			collectibles = append(collectibles, fmt.Sprintf("%s/%s", path, resource.Name()))
		}
	}

	for _, resource := range dirEntries {
		if c, err := forwardCollect(v, collectibles, fmt.Sprintf("%s/%s", path, resource.Name()), depth + 1); err != nil {
			if err.Error() == "not found" {
				continue
			}
			return collectibles, err
		} else {
			collectibles = c
		}
	}

	if afterCollected := len(collectibles); beforeCollected == afterCollected {
		return collectibles, fmt.Errorf("not found")
	}
	return collectibles, nil
}

func backwardCollect(v *visitor.Visitor, collectibles []string, path string, depth int) ([]string, error) {
	if divable, err := v.CheckLanding(path, depth); err != nil {
		return collectibles, err
	} else if !divable {
		if len(collectibles) == 0 {
			return collectibles, fmt.Errorf("not found")
		}
		return collectibles, nil
	}

	dirEntries, _ := os.ReadDir(path)
	for _, resource := range dirEntries {
		info, _ := resource.Info()
		if result, err := v.Evaluator.Evaluate(v.Goal, info); err != nil {
			return collectibles, err
		} else if result {
			collectibles = append(collectibles, fmt.Sprintf("%s/%s", path, resource.Name()))
		}
	}

	return backwardCollect(v, collectibles, "../" + path, depth + 1)
}
