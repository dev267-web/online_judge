package executor

import (
	"context"
	"os/exec"
	"time"
)

type Submission struct {
	ID              int64
	Language        string
	SourceCode      string
	InputData       string
	ExpectedOutput  string
	TimeLimitMillis int
}

type Result struct {
	Status string
	TimeMs int
}

type Executor struct {
	WorkDir string
}

func NewExecutor(workDir string) *Executor {
	return &Executor{WorkDir: workDir}
}

func (e *Executor) RunSubmission(ctx context.Context, sub Submission) (Result, error) {
	start := time.Now()

	// TEMP dummy execution (real docker logic later)
	cmd := exec.CommandContext(ctx, "echo", "hello")
	_ = cmd.Run()

	return Result{
		Status: "accepted",
		TimeMs: int(time.Since(start).Milliseconds()),
	}, nil
}
