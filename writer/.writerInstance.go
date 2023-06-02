package writer

import "github.com/aws/constructs-go/constructs/v10"

// Plug-in with other constructs
type WriterInstanceProps struct {
	//insert props from other constructs
}

type writerInstance struct {
	constructs.Construct
}

/*
type Writer interface {
	constructs.Construct
	//insert useful method to Do construct
}
*/

func NewWriterInstance(scope constructs.Construct, id *string, props *WriterInstanceProps) Writer {
	//implement construct
	this := constructs.NewConstruct(scope, id)
	return writerInstance{this}
}
