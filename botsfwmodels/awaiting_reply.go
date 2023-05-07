package botsfwmodels

import (
	"strings"
)

const (
	// AwaitingReplyToPathSeparator separates parts of the command state
	AwaitingReplyToPathSeparator = "/"

	// AwaitingReplyToPath2QuerySeparator separates path and query parts of state
	AwaitingReplyToPath2QuerySeparator = "?"

	// AwaitingReplyToParamsSeparator separates params of command state
	AwaitingReplyToParamsSeparator = "&"
)

// AwaitingReplyToPath returns just path part of command state
func AwaitingReplyToPath(awaitingReplyTo string) string {
	if i := strings.Index(awaitingReplyTo, AwaitingReplyToPath2QuerySeparator); i >= 0 {
		return awaitingReplyTo[:i]
	}
	return awaitingReplyTo
}

// AwaitingReplyToQuery returns just query part of command state
func AwaitingReplyToQuery(awaitingReplyTo string) string {
	if i := strings.Index(awaitingReplyTo, AwaitingReplyToPath2QuerySeparator); i >= 0 {
		return awaitingReplyTo[i+1:]
	}
	return ""
}
