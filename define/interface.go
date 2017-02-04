package define

type CommonListAPI interface {
	SetEmpty()
	SetLimit(int)
	SetOffset(int)
	SetInclude(string)
	SetExclude(string)
	SetSortBy(string, bool)
	SetFilterBy(string, interface{})
	SetFilterMultiBy(string, interface{})
}
