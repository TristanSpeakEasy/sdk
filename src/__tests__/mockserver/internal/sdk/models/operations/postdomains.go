// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

// PostDomainsRequestBody3 - transfer-in
type PostDomainsRequestBody3 struct {
	// The domain name you want to add.
	Name string `json:"name"`
	// The domain operation to perform. It can be either `add` or `transfer-in`.
	Method string `json:"method"`
	// The authorization code assigned to the domain.
	AuthCode *string `json:"authCode,omitempty"`
	// The price you expect to be charged for the required 1 year renewal.
	ExpectedPrice *float64 `json:"expectedPrice,omitempty"`
}

func (o *PostDomainsRequestBody3) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostDomainsRequestBody3) GetMethod() string {
	if o == nil {
		return ""
	}
	return o.Method
}

func (o *PostDomainsRequestBody3) GetAuthCode() *string {
	if o == nil {
		return nil
	}
	return o.AuthCode
}

func (o *PostDomainsRequestBody3) GetExpectedPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.ExpectedPrice
}

// PostDomainsRequestBody2 - move-in
type PostDomainsRequestBody2 struct {
	// The domain name you want to add.
	Name string `json:"name"`
	// The domain operation to perform. It can be either `add` or `transfer-in`.
	Method string `json:"method"`
	// The move-in token from Move Requested email.
	Token *string `json:"token,omitempty"`
}

func (o *PostDomainsRequestBody2) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostDomainsRequestBody2) GetMethod() string {
	if o == nil {
		return ""
	}
	return o.Method
}

func (o *PostDomainsRequestBody2) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// PostDomainsRequestBody1 - add
type PostDomainsRequestBody1 struct {
	// The domain name you want to add.
	Name string `json:"name"`
	// Whether the domain has the Vercel Edge Network enabled or not.
	CdnEnabled *bool `json:"cdnEnabled,omitempty"`
	Zone       *bool `json:"zone,omitempty"`
	// The domain operation to perform. It can be either `add` or `transfer-in`.
	Method *string `json:"method,omitempty"`
}

func (o *PostDomainsRequestBody1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostDomainsRequestBody1) GetCdnEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdnEnabled
}

func (o *PostDomainsRequestBody1) GetZone() *bool {
	if o == nil {
		return nil
	}
	return o.Zone
}

func (o *PostDomainsRequestBody1) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

type PostDomainsRequestType string

const (
	PostDomainsRequestTypePostDomainsRequestBody1 PostDomainsRequestType = "post_domains_RequestBody_1"
	PostDomainsRequestTypePostDomainsRequestBody2 PostDomainsRequestType = "post_domains_RequestBody_2"
	PostDomainsRequestTypePostDomainsRequestBody3 PostDomainsRequestType = "post_domains_RequestBody_3"
)

type PostDomainsRequest struct {
	PostDomainsRequestBody1 *PostDomainsRequestBody1 `queryParam:"inline"`
	PostDomainsRequestBody2 *PostDomainsRequestBody2 `queryParam:"inline"`
	PostDomainsRequestBody3 *PostDomainsRequestBody3 `queryParam:"inline"`

	Type PostDomainsRequestType
}

func CreatePostDomainsRequestPostDomainsRequestBody1(postDomainsRequestBody1 PostDomainsRequestBody1) PostDomainsRequest {
	typ := PostDomainsRequestTypePostDomainsRequestBody1

	return PostDomainsRequest{
		PostDomainsRequestBody1: &postDomainsRequestBody1,
		Type:                    typ,
	}
}

func CreatePostDomainsRequestPostDomainsRequestBody2(postDomainsRequestBody2 PostDomainsRequestBody2) PostDomainsRequest {
	typ := PostDomainsRequestTypePostDomainsRequestBody2

	return PostDomainsRequest{
		PostDomainsRequestBody2: &postDomainsRequestBody2,
		Type:                    typ,
	}
}

func CreatePostDomainsRequestPostDomainsRequestBody3(postDomainsRequestBody3 PostDomainsRequestBody3) PostDomainsRequest {
	typ := PostDomainsRequestTypePostDomainsRequestBody3

	return PostDomainsRequest{
		PostDomainsRequestBody3: &postDomainsRequestBody3,
		Type:                    typ,
	}
}

func (u *PostDomainsRequest) UnmarshalJSON(data []byte) error {

	var postDomainsRequestBody2 PostDomainsRequestBody2 = PostDomainsRequestBody2{}
	if err := utils.UnmarshalJSON(data, &postDomainsRequestBody2, "", true, true); err == nil {
		u.PostDomainsRequestBody2 = &postDomainsRequestBody2
		u.Type = PostDomainsRequestTypePostDomainsRequestBody2
		return nil
	}

	var postDomainsRequestBody1 PostDomainsRequestBody1 = PostDomainsRequestBody1{}
	if err := utils.UnmarshalJSON(data, &postDomainsRequestBody1, "", true, true); err == nil {
		u.PostDomainsRequestBody1 = &postDomainsRequestBody1
		u.Type = PostDomainsRequestTypePostDomainsRequestBody1
		return nil
	}

	var postDomainsRequestBody3 PostDomainsRequestBody3 = PostDomainsRequestBody3{}
	if err := utils.UnmarshalJSON(data, &postDomainsRequestBody3, "", true, true); err == nil {
		u.PostDomainsRequestBody3 = &postDomainsRequestBody3
		u.Type = PostDomainsRequestTypePostDomainsRequestBody3
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PostDomainsRequest", string(data))
}

func (u PostDomainsRequest) MarshalJSON() ([]byte, error) {
	if u.PostDomainsRequestBody1 != nil {
		return utils.MarshalJSON(u.PostDomainsRequestBody1, "", true)
	}

	if u.PostDomainsRequestBody2 != nil {
		return utils.MarshalJSON(u.PostDomainsRequestBody2, "", true)
	}

	if u.PostDomainsRequestBody3 != nil {
		return utils.MarshalJSON(u.PostDomainsRequestBody3, "", true)
	}

	return nil, errors.New("could not marshal union type PostDomainsRequest: all fields are null")
}

// PostDomainsCreator - An object containing information of the domain creator, including the user's id, username, and email.
type PostDomainsCreator struct {
	Username         string  `json:"username"`
	Email            string  `json:"email"`
	CustomerID       *string `json:"customerId,omitempty"`
	IsDomainReseller *bool   `json:"isDomainReseller,omitempty"`
	ID               string  `json:"id"`
}

func (o *PostDomainsCreator) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *PostDomainsCreator) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *PostDomainsCreator) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *PostDomainsCreator) GetIsDomainReseller() *bool {
	if o == nil {
		return nil
	}
	return o.IsDomainReseller
}

func (o *PostDomainsCreator) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// PostDomainsServiceType - The type of service the domain is handled by. `external` if the DNS is externally handled, `zeit.world` if handled with Vercel, or `na` if the service is not available.
type PostDomainsServiceType string

const (
	PostDomainsServiceTypeZeitWorld PostDomainsServiceType = "zeit.world"
	PostDomainsServiceTypeExternal  PostDomainsServiceType = "external"
	PostDomainsServiceTypeNa        PostDomainsServiceType = "na"
)

func (e PostDomainsServiceType) ToPointer() *PostDomainsServiceType {
	return &e
}
func (e *PostDomainsServiceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zeit.world":
		fallthrough
	case "external":
		fallthrough
	case "na":
		*e = PostDomainsServiceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostDomainsServiceType: %v", v)
	}
}

type PostDomainsDomain struct {
	// If the domain has the ownership verified.
	Verified bool `json:"verified"`
	// A list of the current nameservers of the domain.
	Nameservers []string `json:"nameservers"`
	// A list of the intended nameservers for the domain to point to Vercel DNS.
	IntendedNameservers []string `json:"intendedNameservers"`
	// A list of custom nameservers for the domain to point to. Only applies to domains purchased with Vercel.
	CustomNameservers []string `json:"customNameservers,omitempty"`
	// An object containing information of the domain creator, including the user's id, username, and email.
	Creator PostDomainsCreator `json:"creator"`
	// The domain name.
	Name string `json:"name"`
	// If it was purchased through Vercel, the timestamp in milliseconds when it was purchased.
	BoughtAt *float64 `json:"boughtAt"`
	// Timestamp in milliseconds when the domain was created in the registry.
	CreatedAt float64 `json:"createdAt"`
	// Timestamp in milliseconds at which the domain is set to expire. `null` if not bought with Vercel.
	ExpiresAt *float64 `json:"expiresAt"`
	// The unique identifier of the domain.
	ID string `json:"id"`
	// Timestamp in milliseconds at which the domain was ordered.
	OrderedAt *float64 `json:"orderedAt,omitempty"`
	// Indicates whether the domain is set to automatically renew.
	Renew *bool `json:"renew,omitempty"`
	// The type of service the domain is handled by. `external` if the DNS is externally handled, `zeit.world` if handled with Vercel, or `na` if the service is not available.
	ServiceType PostDomainsServiceType `json:"serviceType"`
	// Timestamp in milliseconds at which the domain was successfully transferred into Vercel. `null` if the transfer is still processing or was never transferred in.
	TransferredAt *float64 `json:"transferredAt,omitempty"`
	// If transferred into Vercel, timestamp in milliseconds when the domain transfer was initiated.
	TransferStartedAt *float64 `json:"transferStartedAt,omitempty"`
	UserID            string   `json:"userId"`
	TeamID            *string  `json:"teamId"`
}

func (o *PostDomainsDomain) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

func (o *PostDomainsDomain) GetNameservers() []string {
	if o == nil {
		return []string{}
	}
	return o.Nameservers
}

func (o *PostDomainsDomain) GetIntendedNameservers() []string {
	if o == nil {
		return []string{}
	}
	return o.IntendedNameservers
}

func (o *PostDomainsDomain) GetCustomNameservers() []string {
	if o == nil {
		return nil
	}
	return o.CustomNameservers
}

func (o *PostDomainsDomain) GetCreator() PostDomainsCreator {
	if o == nil {
		return PostDomainsCreator{}
	}
	return o.Creator
}

func (o *PostDomainsDomain) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostDomainsDomain) GetBoughtAt() *float64 {
	if o == nil {
		return nil
	}
	return o.BoughtAt
}

func (o *PostDomainsDomain) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *PostDomainsDomain) GetExpiresAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *PostDomainsDomain) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PostDomainsDomain) GetOrderedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.OrderedAt
}

func (o *PostDomainsDomain) GetRenew() *bool {
	if o == nil {
		return nil
	}
	return o.Renew
}

func (o *PostDomainsDomain) GetServiceType() PostDomainsServiceType {
	if o == nil {
		return PostDomainsServiceType("")
	}
	return o.ServiceType
}

func (o *PostDomainsDomain) GetTransferredAt() *float64 {
	if o == nil {
		return nil
	}
	return o.TransferredAt
}

func (o *PostDomainsDomain) GetTransferStartedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.TransferStartedAt
}

func (o *PostDomainsDomain) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *PostDomainsDomain) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type PostDomainsResponseBody struct {
	Domain PostDomainsDomain `json:"domain"`
}

func (o *PostDomainsResponseBody) GetDomain() PostDomainsDomain {
	if o == nil {
		return PostDomainsDomain{}
	}
	return o.Domain
}

type PostDomainsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *PostDomainsResponseBody
}

func (o *PostDomainsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostDomainsResponse) GetObject() *PostDomainsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
