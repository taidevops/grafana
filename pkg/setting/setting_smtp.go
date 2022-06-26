package setting

type SmtpSettings struct {
	Enabled        bool
	Host           string
	User           string
	Password       string
	CertFile       string
	KeyFile        string
	FromAddress    string
	FromName       string
	EhloIdentity   string
	StartTLSPolicy string
	SkipVerify     bool

	SendWelcomeEmailOnSignUp bool
	TemplatesPatterns        []string
	ContentTypes             []string
}

func (cfg *Cfg) readSmtpSettings() {
	sec := cfg.Raw.Section("smtp")
}
