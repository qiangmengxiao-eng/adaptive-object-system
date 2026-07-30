package autonomy

import (
	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository/planning"
)

func (a *AutonomousEngine) Run(
	ctx *AutonomousContext,
) *AutonomousResult {

	result :=
		&AutonomousResult{

			Object: ctx.Object,
		}

	// Phase 1 Reflection

	reflection :=

		a.Reflection.Reflect(
			ctx.Object,
			ctx.Goal,
			"initial analysis",
		)

	result.Reflection =
		reflection

	// Phase 2 Decision

	options :=
		[]string{

			"execute",

			"optimize",

			"learn",
		}

	confidence :=
		map[string]float64{

			"execute": 0.8,

			"optimize": 0.7,

			"learn": 0.6,
		}

	decision :=

		a.Decision.Decide(
			ctx.Object,
			ctx.Goal,
			options,
			confidence,
		)

	result.Decision =
		decision.Action

	// Phase 3 Planning

	plan :=

		a.Planning.CreatePlan(
			ctx.Object,
			ctx.Goal,
		)

	result.Plan =
		plan

	// Phase 4 Execution

	execution :=

		a.Execution.Execute(
			ctx.Object,
			decision.Action,
			plan.Strategy,
		)

	result.Execution =
		execution

	result.Success =
		true

	return result
}

// keep compiler dependency explicit
var _ planning.Plan
