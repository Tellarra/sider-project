package model

import "time"

type Orga struct {
	ID                      string        `json:"_id"`
	Name                    string        `json:"name"`
	Tweet                   string        `json:"tweet"`
	CompanySize             string        `json:"companySize"`
	SignupOrigin            string        `json:"signupOrigin"`
	Website                 string        `json:"website"`
	LogoUrl                 string        `json:"logoUrl"`
	VideoUrl                string        `json:"videoUrl"`
	CoverUrl                string        `json:"coverUrl"`
	Users                   []OrgaUser    `json:"users"`
	PhoneNumber             string        `json:"phoneNumber"`
	InvoiceInfos            InvoiceInfo   `json:"invoiceInfos"`
	PaymentMethod           PaymentMethod `json:"paymentMethod"`
	NotPaying               bool          `json:"notPaying"`
	StripeCustomerID        string        `json:"stripeCustomerID"`
	PricingPlan             PricingPlan   `json:"pricingPlan"`
	SCustomerNumber         string        `json:"sCustomerNumber"`
	SCustomerType           string        `json:"sCustomerType"`
	Services                OrgaServices  `json:"services"`
	WorkLegalStatus         string        `json:"workLegalStatus"`
	Siret                   string        `json:"siret"`
	APECode                 string        `json:"APECode"`
	PricingMethod           string        `json:"pricingMethod"`
	Industry                string        `json:"industry"`
	PricingId               string        `json:"pricingId"`
	Blocked                 *bool         `json:"blocked"`
	ParentOrganisationIds   *[]string     `json:"parentOrganisationIds"`
	ChildOrganisationIds    *[]string     `json:"childOrganisationIds"`
	OrganisationType        string        `json:"organisationType"`
	ShopCode                string        `json:"shopCode"`
	Address                 string        `json:"address"`
	City                    string        `json:"city"`
	Zip                     string        `json:"zip"`
	Features                *[]string     `json:"features"`
	IndemnitiesConfirmation Confirmation  `json:"indemnitiesConfirmation"`
	HoursTypesConfirmation  Confirmation  `json:"hoursTypesConfirmation"`
}

type OrgaToFeed struct {
	ID                      string        `json:"id"`
	Name                    string        `json:"name"`
	Tweet                   string        `json:"tweet"`
	CompanySize             string        `json:"companySize"`
	SignupOrigin            string        `json:"signupOrigin"`
	Website                 string        `json:"website"`
	LogoUrl                 string        `json:"logoUrl"`
	VideoUrl                string        `json:"videoUrl"`
	CoverUrl                string        `json:"coverUrl"`
	Users                   []OrgaUser    `json:"users"`
	PhoneNumber             string        `json:"phoneNumber"`
	InvoiceInfos            InvoiceInfo   `json:"invoiceInfos"`
	PaymentMethod           PaymentMethod `json:"paymentMethod"`
	NotPaying               bool          `json:"notPaying"`
	StripeCustomerID        string        `json:"stripeCustomerID"`
	PricingPlan             PricingPlan   `json:"pricingPlan"`
	SCustomerNumber         string        `json:"sCustomerNumber"`
	SCustomerType           string        `json:"sCustomerType"`
	Services                OrgaServices  `json:"services"`
	WorkLegalStatus         string        `json:"workLegalStatus"`
	Siret                   string        `json:"siret"`
	APECode                 string        `json:"APECode"`
	PricingMethod           string        `json:"pricingMethod"`
	Industry                string        `json:"industry"`
	PricingId               string        `json:"pricingId"`
	Blocked                 *bool         `json:"blocked"`
	ParentOrganisationIds   *[]string     `json:"parentOrganisationIds"`
	ChildOrganisationIds    *[]string     `json:"childOrganisationIds"`
	OrganisationType        string        `json:"organisationType"`
	ShopCode                string        `json:"shopCode"`
	Address                 string        `json:"address"`
	City                    string        `json:"city"`
	Zip                     string        `json:"zip"`
	Features                *[]string     `json:"features"`
	IndemnitiesConfirmation Confirmation  `json:"indemnitiesConfirmation"`
	HoursTypesConfirmation  Confirmation  `json:"hoursTypesConfirmation"`
}

type OrgaUser struct {
	ID     string   `json:"id"`
	Status string   `json:"status"`
	Roles  []string `json:"roles"`
}

type InvoiceInfo struct {
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Address InvoiceAddress `json:"address"`
}

type InvoiceAddress struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
	Country    string `json:"country"`
}

type PaymentMethod struct {
	Type  string `json:"type"`
	Last4 string `json:"last4"`
}

type PricingPlan struct {
	StockManager StockManager `json:"stockManager"`
}

type StockManager struct {
	Sider int `json:"sider"`
	Side  int `json:"side"`
}

type OrgaServices struct {
	Google GoogleService `json:"google"`
}

type GoogleService struct {
	SpreadsheetId string `json:"spreadsheetId"`
}

type Confirmation struct {
	Date   time.Time `json:"date"`
	UserID string    `json:"userId"`
}

func (o Orga) ToOrgaToFeed() OrgaToFeed {
	return OrgaToFeed(o)
}
