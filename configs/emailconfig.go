package configs

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

func LoadEmailConfig() EmailConfig {

	return EmailConfig{
		SMTPHost:     "smtp.gmail.com",
		SMTPPort:     587,
		SMTPUsername: "guptapiyush238@gmail.com",
		SMTPPassword: "dwrbnuhqroeqissu",
		// SMTPPassword: "dwrb nuhq roeq issu",
	}
}
