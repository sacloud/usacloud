package params

import (
	"encoding/json"
	"fmt"
	"io"
)

type skeletonFiller interface {
	fillValueToSkeleton()
}

func writeSkeleton(in interface{}, writer io.Writer) error {
	if fill, ok := in.(skeletonFiller); ok {
		fill.fillValueToSkeleton()
	}

	data, err := json.Marshal(in)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(writer, data)
	return err
}
