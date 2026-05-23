package executor

import (
	"fmt"

	ctx "github.com/KandakatlaChandramouli/ElasticKV/internal/runtime/context"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/pipeline/stages"
)

func Execute(
	context ctx.QueryContext,
) {

	fmt.Println(
		"\n=== ElasticKV Runtime ===\n",
	)

	stages.ParseQuery(
		context.Query,
	)

	stages.PlanQuery()

	stages.ExecuteVectorSearch()

	if context.EnableGPU {

		stages.ExecuteGPU()
	}

	stages.AggregateResults()

	fmt.Println(
		"\n=== query completed ===\n",
	)
}
