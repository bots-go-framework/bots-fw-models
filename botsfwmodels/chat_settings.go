package botsfwmodels

type chatSettings struct {
	PreferredLanguage string   `dalgo:"preferredLanguage,noindex,omitempty" firestore:"preferredLanguage,omitempty"`
	LanguageCodes     []string `dalgo:"languageCodes,noindex,omitempty" firestore:"languageCodes,omitempty"` // UI languages
}

// GetPreferredLanguage returns preferred language
func (e *chatSettings) GetPreferredLanguage() string {
	return e.PreferredLanguage
}

// SetPreferredLanguage sets preferred language
func (e *chatSettings) SetPreferredLanguage(value string) {
	e.PreferredLanguage = value
}

// AddClientLanguage adds client UI language
func (e *chatSettings) AddClientLanguage(languageCode string) (changed bool) {
	if languageCode == "" || languageCode == "root" {
		return false
	}
	for _, lc := range e.LanguageCodes {
		if lc == languageCode {
			return false
		}
	}
	e.LanguageCodes = append(e.LanguageCodes, languageCode)
	if e.PreferredLanguage == "" {
		e.PreferredLanguage = languageCode
	}
	return false
}
