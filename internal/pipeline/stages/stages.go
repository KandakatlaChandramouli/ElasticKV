package stages

import "fmt"

func ParseQuery(
	query string,
) {

	fmt.Println(
		"[stage] parsing query",
	)
}

func PlanQuery() {

	fmt.Println(
		"[stage] planning query",
	)
}

func ExecuteVectorSearch() {

	fmt.Println(
		"[stage] vector search",
	)
}

func ExecuteGPU() {

	fmt.Println(
		"[stage] gpu execution",
	)
}

func AggregateResults() {

	fmt.Println(
		"[stage] aggregating results",
	)
}
