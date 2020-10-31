package base

type FindParameter struct {
	Count int `cli:",aliases=max limit,category=filter"`
	From  int `cli:",aliases=offset,category=filter"`
}
