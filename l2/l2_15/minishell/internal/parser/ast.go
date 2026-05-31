package parser

// Redirect is an input or output file.
type Redirect struct {
	Path string
}

// Command is a single command with optional redirects.
type Command struct {
	Args      []string
	Stdin     *Redirect
	Stdout    *Redirect
	AppendOut bool
}

// Pipeline is a chain of commands connected by |.
type Pipeline struct {
	Commands []Command
}

// Job is a full input line: pipelines and && / || operators.
type Job struct {
	Pipelines []Pipeline
	Ops       []string
}
