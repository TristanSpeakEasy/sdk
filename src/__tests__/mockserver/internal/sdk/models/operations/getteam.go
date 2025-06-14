// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetTeamRequest struct {
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The Team identifier to perform the request on behalf of.
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
}

func (o *GetTeamRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *GetTeamRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

type GetTeamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The requested team
	Team *components.Team
}

func (o *GetTeamResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTeamResponse) GetTeam() *components.Team {
	if o == nil {
		return nil
	}
	return o.Team
}
