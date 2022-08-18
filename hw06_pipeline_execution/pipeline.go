package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

// ExecutePipeline builds pipeline from stages.
// If stage is nil it will be skipped.
func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := wrapWithDone(in, done)
	for _, s := range stages {
		out = s(wrapWithDone(out, done))
	}
	return out
}

func wrapWithDone(in In, done In) Out {
	out := make(Bi)

	go func() {
		defer func() {
			close(out)
			for range in {
			}
		}()

		for {
			select {
			case <-done:
				return
			default:
			}

			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()

	return out
}
