package writer

type writerAdapter struct {
	writerTaskContainer
	writerApiLambda
}

type Writer interface {
	WriterFunc
	WriterTask
	PlugExecutor() Writer
}

func (wr writerAdapter) PlugExecutor() Writer {

}
