package botsfwmodels

import (
	"fmt"
	"net/url"
	"strings"
)

type chatState struct {
	changed         bool
	AwaitingReplyTo string `dalgo:",noindex,omitempty" datastore:",noindex,omitempty" firestore:",omitempty"`
}

func (e *chatState) IsChanged() bool {
	return e.changed
}

// GetAwaitingReplyTo returns current state
func (e *chatState) GetAwaitingReplyTo() string {
	return e.AwaitingReplyTo
}

// SetAwaitingReplyTo sets current state
func (e *chatState) SetAwaitingReplyTo(value string) {
	e.AwaitingReplyTo = strings.TrimLeft(value, "/")
	e.changed = true
}

// IsAwaitingReplyTo returns true if bot us awaiting reply to a specific command
func (e *chatState) IsAwaitingReplyTo(code string) bool {
	awaitingReplyToPath := e.getAwaitingReplyToPath()
	return awaitingReplyToPath == code || strings.HasSuffix(awaitingReplyToPath, AwaitingReplyToPathSeparator+code)
}

func (e *chatState) getAwaitingReplyToPath() string {
	pathAndQuery := strings.SplitN(e.AwaitingReplyTo, AwaitingReplyToPath2QuerySeparator, 2)
	if len(pathAndQuery) > 1 {
		return pathAndQuery[0]
	}
	return e.AwaitingReplyTo
}

// PopStepsFromAwaitingReplyUpToSpecificParent go back in state
func (e *chatState) PopStepsFromAwaitingReplyUpToSpecificParent(step string) {
	awaitingReplyTo := e.AwaitingReplyTo
	pathAndQuery := strings.SplitN(awaitingReplyTo, AwaitingReplyToPath2QuerySeparator, 2)
	path := pathAndQuery[0]
	steps := strings.Split(path, AwaitingReplyToPathSeparator)
	for i := len(steps) - 1; i >= 0; i-- {
		if steps[i] == step {
			if i < len(steps)-1 {
				path = strings.Join(steps[:i+1], AwaitingReplyToPathSeparator)
				if len(pathAndQuery) > 1 {
					query := pathAndQuery[1]
					e.SetAwaitingReplyTo(path + AwaitingReplyToPath2QuerySeparator + query)
				} else {
					e.SetAwaitingReplyTo(path)
				}
			}
			//steps = steps[:i]
			break
			// } else {
			// logMessage.Infof(c, "steps[%v]: %v != %v:", i, steps[i], step)
		}
	}
}

// PushStepToAwaitingReplyTo go down in state
func (e *chatState) PushStepToAwaitingReplyTo(step string) {
	awaitingReplyTo := e.AwaitingReplyTo
	pathAndQuery := strings.SplitN(awaitingReplyTo, AwaitingReplyToPath2QuerySeparator, 2)
	if len(pathAndQuery) > 1 { // Has query part - something after "?" character
		if !e.IsAwaitingReplyTo(step) {
			path := pathAndQuery[0]
			query := pathAndQuery[1]
			awaitingReplyTo = strings.Join([]string{path, AwaitingReplyToPathSeparator, step, AwaitingReplyToPath2QuerySeparator, query}, "")
			e.SetAwaitingReplyTo(awaitingReplyTo)
		}
	} else { // Has no query - no "?" character
		if !e.IsAwaitingReplyTo(step) {
			awaitingReplyTo = awaitingReplyTo + AwaitingReplyToPathSeparator + step
			e.SetAwaitingReplyTo(awaitingReplyTo)
		}
	}
}

// AddWizardParam adds context param to state
func (e *chatState) AddWizardParam(key, value string) {
	awaitingReplyTo := e.GetAwaitingReplyTo()
	awaitingURL, err := url.Parse(awaitingReplyTo)
	if err != nil {
		panic(fmt.Sprintf("Failed to call url.Parse(awaitingReplyTo=%v)", awaitingReplyTo))
	}
	query := awaitingURL.Query()
	query.Set(key, value)
	awaitingURL.RawQuery = query.Encode()
	e.SetAwaitingReplyTo(awaitingURL.String())
}

// GetWizardParam returns state param value
func (e *chatState) GetWizardParam(key string) string {
	u, err := url.Parse(e.GetAwaitingReplyTo())
	if err != nil {
		return ""
	}
	return u.Query().Get(key)
}
