package ui

import (
	"sort"
)

const (
	runningString = "Running"
	skippedString = "Skipping"

	environmentMessage = "Set environment variables"
)

// PrintCommand prints the command to be executed.
func PrintCommand(command string) {
	if Verbosity <= VerbosityLevelQuiet {
		return
	}

	printf(
		LoggerStderr,
		"[%s] %s\n",
		blue(runningString),
		bold(command),
	)
}

// PrintEnvironment prints when environment variables are set.
func PrintEnvironment(variables map[string]*string) {
	if Verbosity <= VerbosityLevelQuiet {
		return
	}

	if len(variables) == 0 {
		return
	}

	printf(
		LoggerStderr,
		"[%s] %s\n",
		blue(runningString),
		environmentMessage,
	)

	// Print in deterministic order
	var keys []string
	for key := range variables {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := variables[key]
		if value == nil {
			continue
		}

		printf(
			LoggerStderr,
			"%sset %s=%s",
			cyan(outputPrefix),
			bold(key),
			*value,
		)
	}

	for _, key := range keys {
		value := variables[key]
		if value != nil {
			continue
		}

		printf(
			LoggerStderr,
			"%sunset %s",
			cyan(outputPrefix),
			bold(key),
		)
	}
}

// PrintSkipped prints the command skipped and the reason.
func PrintSkipped(command string, reason string) {
	if Verbosity < VerbosityLevelVerbose {
		return
	}

	printf(
		LoggerStderr,
		"[%s] %s\n%s%s\n",
		yellow(skippedString),
		bold(command),
		cyan(outputPrefix),
		reason,
	)
}

// PrintCommandError prints an error from a running command.
func PrintCommandError(err error) {
	if Verbosity <= VerbosityLevelQuiet {
		return
	}

	printf(
		LoggerStderr,
		"%s%s\n",
		red(outputPrefix),
		err.Error(),
	)
}
