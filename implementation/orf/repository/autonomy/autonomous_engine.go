package autonomy

import (
	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository/decision"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository/planning"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository/execution"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository/reflection"
)

type AutonomousEngine struct {
	Decision *decision.DecisionEngine

	Planning *planning.PlanningEngine

	Execution *execution.ExecutionEngine

	Reflection *reflection.ReflectionEngine
}

func NewAutonomousEngine(

	d *decision.DecisionEngine,

	p *planning.PlanningEngine,

	e *execution.ExecutionEngine,

	r *reflection.ReflectionEngine,

) *AutonomousEngine {

	return &AutonomousEngine{

		Decision: d,

		Planning: p,

		Execution: e,

		Reflection: r,
	}
}
