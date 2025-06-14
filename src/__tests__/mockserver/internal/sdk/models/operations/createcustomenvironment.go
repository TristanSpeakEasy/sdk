// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
)

// CreateCustomEnvironmentTypeRequest - Type of matcher. One of \"equals\", \"startsWith\", or \"endsWith\".
type CreateCustomEnvironmentTypeRequest string

const (
	CreateCustomEnvironmentTypeRequestEquals     CreateCustomEnvironmentTypeRequest = "equals"
	CreateCustomEnvironmentTypeRequestStartsWith CreateCustomEnvironmentTypeRequest = "startsWith"
	CreateCustomEnvironmentTypeRequestEndsWith   CreateCustomEnvironmentTypeRequest = "endsWith"
)

func (e CreateCustomEnvironmentTypeRequest) ToPointer() *CreateCustomEnvironmentTypeRequest {
	return &e
}
func (e *CreateCustomEnvironmentTypeRequest) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "equals":
		fallthrough
	case "startsWith":
		fallthrough
	case "endsWith":
		*e = CreateCustomEnvironmentTypeRequest(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCustomEnvironmentTypeRequest: %v", v)
	}
}

// CreateCustomEnvironmentBranchMatcherRequest - How we want to determine a matching branch. This is optional.
type CreateCustomEnvironmentBranchMatcherRequest struct {
	// Type of matcher. One of \"equals\", \"startsWith\", or \"endsWith\".
	Type CreateCustomEnvironmentTypeRequest `json:"type"`
	// Git branch name or portion thereof.
	Pattern string `json:"pattern"`
}

func (o *CreateCustomEnvironmentBranchMatcherRequest) GetType() CreateCustomEnvironmentTypeRequest {
	if o == nil {
		return CreateCustomEnvironmentTypeRequest("")
	}
	return o.Type
}

func (o *CreateCustomEnvironmentBranchMatcherRequest) GetPattern() string {
	if o == nil {
		return ""
	}
	return o.Pattern
}

type CreateCustomEnvironmentRequestBody struct {
	// The slug of the custom environment to create.
	Slug *string `json:"slug,omitempty"`
	// Description of the custom environment. This is optional.
	Description *string `json:"description,omitempty"`
	// How we want to determine a matching branch. This is optional.
	BranchMatcher *CreateCustomEnvironmentBranchMatcherRequest `json:"branchMatcher,omitempty"`
	// Where to copy environment variables from. This is optional.
	CopyEnvVarsFrom *string `json:"copyEnvVarsFrom,omitempty"`
}

func (o *CreateCustomEnvironmentRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateCustomEnvironmentRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateCustomEnvironmentRequestBody) GetBranchMatcher() *CreateCustomEnvironmentBranchMatcherRequest {
	if o == nil {
		return nil
	}
	return o.BranchMatcher
}

func (o *CreateCustomEnvironmentRequestBody) GetCopyEnvVarsFrom() *string {
	if o == nil {
		return nil
	}
	return o.CopyEnvVarsFrom
}

type CreateCustomEnvironmentRequest struct {
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                             `queryParam:"style=form,explode=true,name=slug"`
	RequestBody *CreateCustomEnvironmentRequestBody `request:"mediaType=application/json"`
}

func (o *CreateCustomEnvironmentRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *CreateCustomEnvironmentRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *CreateCustomEnvironmentRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateCustomEnvironmentRequest) GetRequestBody() *CreateCustomEnvironmentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

// CreateCustomEnvironmentTypeResponseBody - The type of environment (production, preview, or development)
type CreateCustomEnvironmentTypeResponseBody string

const (
	CreateCustomEnvironmentTypeResponseBodyProduction  CreateCustomEnvironmentTypeResponseBody = "production"
	CreateCustomEnvironmentTypeResponseBodyPreview     CreateCustomEnvironmentTypeResponseBody = "preview"
	CreateCustomEnvironmentTypeResponseBodyDevelopment CreateCustomEnvironmentTypeResponseBody = "development"
)

func (e CreateCustomEnvironmentTypeResponseBody) ToPointer() *CreateCustomEnvironmentTypeResponseBody {
	return &e
}
func (e *CreateCustomEnvironmentTypeResponseBody) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "production":
		fallthrough
	case "preview":
		fallthrough
	case "development":
		*e = CreateCustomEnvironmentTypeResponseBody(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCustomEnvironmentTypeResponseBody: %v", v)
	}
}

// CreateCustomEnvironmentBranchMatcherTypeResponseBody - The type of matching to perform
type CreateCustomEnvironmentBranchMatcherTypeResponseBody string

const (
	CreateCustomEnvironmentBranchMatcherTypeResponseBodyEndsWith   CreateCustomEnvironmentBranchMatcherTypeResponseBody = "endsWith"
	CreateCustomEnvironmentBranchMatcherTypeResponseBodyStartsWith CreateCustomEnvironmentBranchMatcherTypeResponseBody = "startsWith"
	CreateCustomEnvironmentBranchMatcherTypeResponseBodyEquals     CreateCustomEnvironmentBranchMatcherTypeResponseBody = "equals"
)

func (e CreateCustomEnvironmentBranchMatcherTypeResponseBody) ToPointer() *CreateCustomEnvironmentBranchMatcherTypeResponseBody {
	return &e
}
func (e *CreateCustomEnvironmentBranchMatcherTypeResponseBody) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "endsWith":
		fallthrough
	case "startsWith":
		fallthrough
	case "equals":
		*e = CreateCustomEnvironmentBranchMatcherTypeResponseBody(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCustomEnvironmentBranchMatcherTypeResponseBody: %v", v)
	}
}

// CreateCustomEnvironmentBranchMatcherResponseBody - Configuration for matching git branches to this environment
type CreateCustomEnvironmentBranchMatcherResponseBody struct {
	// The type of matching to perform
	Type CreateCustomEnvironmentBranchMatcherTypeResponseBody `json:"type"`
	// The pattern to match against branch names
	Pattern string `json:"pattern"`
}

func (o *CreateCustomEnvironmentBranchMatcherResponseBody) GetType() CreateCustomEnvironmentBranchMatcherTypeResponseBody {
	if o == nil {
		return CreateCustomEnvironmentBranchMatcherTypeResponseBody("")
	}
	return o.Type
}

func (o *CreateCustomEnvironmentBranchMatcherResponseBody) GetPattern() string {
	if o == nil {
		return ""
	}
	return o.Pattern
}

// CreateCustomEnvironmentVerification - A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
type CreateCustomEnvironmentVerification struct {
	Type   string `json:"type"`
	Domain string `json:"domain"`
	Value  string `json:"value"`
	Reason string `json:"reason"`
}

func (o *CreateCustomEnvironmentVerification) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *CreateCustomEnvironmentVerification) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *CreateCustomEnvironmentVerification) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *CreateCustomEnvironmentVerification) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

// CreateCustomEnvironmentDomain - List of domains associated with this environment
type CreateCustomEnvironmentDomain struct {
	Name                string   `json:"name"`
	ApexName            string   `json:"apexName"`
	ProjectID           string   `json:"projectId"`
	Redirect            *string  `json:"redirect,omitempty"`
	RedirectStatusCode  *float64 `json:"redirectStatusCode,omitempty"`
	GitBranch           *string  `json:"gitBranch,omitempty"`
	CustomEnvironmentID *string  `json:"customEnvironmentId,omitempty"`
	UpdatedAt           *float64 `json:"updatedAt,omitempty"`
	CreatedAt           *float64 `json:"createdAt,omitempty"`
	// `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
	Verified bool `json:"verified"`
	// A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
	Verification []CreateCustomEnvironmentVerification `json:"verification,omitempty"`
}

func (o *CreateCustomEnvironmentDomain) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateCustomEnvironmentDomain) GetApexName() string {
	if o == nil {
		return ""
	}
	return o.ApexName
}

func (o *CreateCustomEnvironmentDomain) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *CreateCustomEnvironmentDomain) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *CreateCustomEnvironmentDomain) GetRedirectStatusCode() *float64 {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

func (o *CreateCustomEnvironmentDomain) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *CreateCustomEnvironmentDomain) GetCustomEnvironmentID() *string {
	if o == nil {
		return nil
	}
	return o.CustomEnvironmentID
}

func (o *CreateCustomEnvironmentDomain) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CreateCustomEnvironmentDomain) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CreateCustomEnvironmentDomain) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

func (o *CreateCustomEnvironmentDomain) GetVerification() []CreateCustomEnvironmentVerification {
	if o == nil {
		return nil
	}
	return o.Verification
}

// CreateCustomEnvironmentResponseBody - Internal representation of a custom environment with all required properties
type CreateCustomEnvironmentResponseBody struct {
	// Unique identifier for the custom environment (format: env_*)
	ID string `json:"id"`
	// URL-friendly name of the environment
	Slug string `json:"slug"`
	// The type of environment (production, preview, or development)
	Type CreateCustomEnvironmentTypeResponseBody `json:"type"`
	// Optional description of the environment's purpose
	Description *string `json:"description,omitempty"`
	// Configuration for matching git branches to this environment
	BranchMatcher *CreateCustomEnvironmentBranchMatcherResponseBody `json:"branchMatcher,omitempty"`
	// List of domains associated with this environment
	Domains []CreateCustomEnvironmentDomain `json:"domains,omitempty"`
	// List of aliases for the current deployment
	CurrentDeploymentAliases []string `json:"currentDeploymentAliases,omitempty"`
	// Timestamp when the environment was created
	CreatedAt float64 `json:"createdAt"`
	// Timestamp when the environment was last updated
	UpdatedAt float64 `json:"updatedAt"`
}

func (o *CreateCustomEnvironmentResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateCustomEnvironmentResponseBody) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *CreateCustomEnvironmentResponseBody) GetType() CreateCustomEnvironmentTypeResponseBody {
	if o == nil {
		return CreateCustomEnvironmentTypeResponseBody("")
	}
	return o.Type
}

func (o *CreateCustomEnvironmentResponseBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateCustomEnvironmentResponseBody) GetBranchMatcher() *CreateCustomEnvironmentBranchMatcherResponseBody {
	if o == nil {
		return nil
	}
	return o.BranchMatcher
}

func (o *CreateCustomEnvironmentResponseBody) GetDomains() []CreateCustomEnvironmentDomain {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *CreateCustomEnvironmentResponseBody) GetCurrentDeploymentAliases() []string {
	if o == nil {
		return nil
	}
	return o.CurrentDeploymentAliases
}

func (o *CreateCustomEnvironmentResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *CreateCustomEnvironmentResponseBody) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

type CreateCustomEnvironmentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *CreateCustomEnvironmentResponseBody
}

func (o *CreateCustomEnvironmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateCustomEnvironmentResponse) GetObject() *CreateCustomEnvironmentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
