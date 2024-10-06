package cmdrunner

type CmdOutput struct {
	Output string
	Err    string
}

func NewCmdOutput(output, err string) CmdOutput {
	return CmdOutput{
		Output: output,
		Err:    err,
	}
}
