package model

type Message struct {
	PrivacyNotice struct {
		Introduction         string `json:"Introduction"`
		InformationWeCollect struct {
			PersonalInformation    []string `json:"PersonalInformation"`
			NonPersonalInformation []string `json:"NonPersonalInformation"`
			Purpose                string   `json:"Purpose"`
		} `json:"InformationWeCollect"`
		HowWeUseYourInformation struct {
			Purpose   string `json:"Purpose"`
			Marketing string `json:"Marketing"`
		} `json:"HowWeUseYourInformation"`
		InformationSharing struct {
			Policy              string `json:"Policy"`
			TrustedThirdParties string `json:"TrustedThirdParties"`
		} `json:"InformationSharing"`
		DataSecurity struct {
			Measures    string `json:"Measures"`
			Limitations string `json:"Limitations"`
		} `json:"DataSecurity"`
		YourChoices struct {
			ManageYourInformation string `json:"ManageYourInformation"`
			OptOut                string `json:"OptOut"`
		} `json:"YourChoices"`
		LegalCompliance struct {
			Regulations string `json:"Regulations"`
		} `json:"LegalCompliance"`
		Consent            string `json:"Consent"`
		ContactInformation string `json:"ContactInformation"`
		LastUpdated        string `json:"LastUpdated"`
	} `json:"PrivacyNotice"`
}
