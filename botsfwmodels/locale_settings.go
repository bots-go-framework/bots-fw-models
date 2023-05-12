package botsfwmodels

type LocaleSettings interface {
	// SetPreferredLocale sets preferred locale for the chat as code5 (e.g. en-US)
	SetPreferredLocale(code5 string) error

	// GetPreferredLocale returns preferred locale for the chat as code5 (e.g. en-US)
	GetPreferredLocale() string
}
