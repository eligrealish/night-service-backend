package policy

type Message struct {
	PrivacyNotice struct {
		Introduction         string `json:"Introduction" bson:"Introduction"`
		InformationWeCollect struct {
			PersonalInformation    []string `json:"PersonalInformation" bson:"PersonalInformation"`
			NonPersonalInformation []string `json:"NonPersonalInformation" bson:"NonPersonalInformation"`
			Purpose                string   `json:"Purpose" bson:"Purpose"`
		} `json:"InformationWeCollect" bson:"InformationWeCollect"`
		HowWeUseYourInformation struct {
			Purpose   string `json:"Purpose" bson:"Purpose"`
			Marketing string `json:"Marketing" bson:"Marketing"`
		} `json:"HowWeUseYourInformation" bson:"HowWeUseYourInformation"`
		InformationSharing struct {
			Policy              string `json:"Policy" bson:"Policy"`
			TrustedThirdParties string `json:"TrustedThirdParties" bson:"TrustedThirdParties"`
		} `json:"InformationSharing" bson:"InformationSharing"`
		DataSecurity struct {
			Measures    string `json:"Measures" bson:"Measures"`
			Limitations string `json:"Limitations" bson:"Limitations"`
		} `json:"DataSecurity" bson:"DataSecurity"`
		YourChoices struct {
			ManageYourInformation string `json:"ManageYourInformation" bson:"ManageYourInformation"`
			OptOut                string `json:"OptOut" bson:"OptOut"`
		} `json:"YourChoices" bson:"YourChoices"`
		LegalCompliance struct {
			Regulations string `json:"Regulations" bson:"Regulations"`
		} `json:"LegalCompliance" bson:"LegalCompliance"`
		Consent            string `json:"Consent" bson:"Consent"`
		ContactInformation string `json:"ContactInformation" bson:"ContactInformation"`
		LastUpdated        string `json:"LastUpdated" bson:"LastUpdated"`
	} `json:"PrivacyNotice" bson:"PrivacyNotice"`
}
