package base

// CommonParameter 全コマンド共通フィールド
type CommonParameter struct {
	Parameters       string `cli:",category=input,desc=Input parameters in JSON format"`
	ParameterFile    string `cli:",category=input,desc=Input parameters in JSON format(from file)"`
	GenerateSkeleton bool   `cli:",category=input,desc=Output skeleton of parameters with JSON format"`
}
