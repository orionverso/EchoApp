package writer

type WriterTaskAdapter struct {
	Writertask writerTask
}

func (wr WriterTaskAdapter) PlugWriter() any {
	return wr.PlugWriter()
}

type WriterApiLambdaAdapter struct {
	Writerapilambda writerApiLambda
}

func (wr WriterApiLambdaAdapter) PlugWriter() any {
	return wr.PlugWriter()
}

type PluginWriter interface {
	PlugWriter() any
}

func MakePlug(pl PluginWriter) any {
	return pl.PlugWriter()
}
