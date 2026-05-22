package backpressure

type Controller struct {
	limit int
}

func NewController(
	limit int,
) *Controller {

	return &Controller{
		limit: limit,
	}
}

func (c *Controller) Allow(
	inflight int,
) bool {

	return inflight < c.limit
}
