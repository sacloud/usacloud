package base

type ConfirmParameter struct {
	AssumeYes bool `cli:"assumeyes,short=y,category=input,desc=Assume that the answer to any question which would be asked is yes"`
}
