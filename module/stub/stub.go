package stub

import "go-crawler/module"

type myModule struct {
	mid             module.MID
	addr            string
	score           uint64
	scoreCalculator module.CalculateScore
	calledCount     uint64
	acceptedCount   uint64
	completedCount  uint64
	handlingNumber  uint64
}

func NewModuleInternal(mid module.MID, scoreCalculator module.CalculateScore) (ModuleInternal, error) {
	parts, err := module.SplitMID(mid)

}
