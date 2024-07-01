// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

type CreateDomainGlobals struct {
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	WorkspaceID *string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *CreateDomainGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

func (o *CreateDomainGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

type CreateDomainRequestBody struct {
	// Name of the domain.
	Slug string `json:"slug"`
	// Redirect users to a specific URL when any link under this domain has expired.
	ExpiredURL *string `json:"expiredUrl,omitempty"`
	// Whether to archive this domain. `false` will unarchive a previously archived domain.
	Archived *bool `default:"false" json:"archived"`
	// Provide context to your teammates in the link creation modal by showing them an example of a link to be shortened.
	Placeholder *string `default:"https://dub.co/help/article/what-is-dub" json:"placeholder"`
}

func (c CreateDomainRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateDomainRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateDomainRequestBody) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *CreateDomainRequestBody) GetExpiredURL() *string {
	if o == nil {
		return nil
	}
	return o.ExpiredURL
}

func (o *CreateDomainRequestBody) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *CreateDomainRequestBody) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

type CreateDomainResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The domain was created.
	DomainSchema *components.DomainSchema
}

func (o *CreateDomainResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateDomainResponse) GetDomainSchema() *components.DomainSchema {
	if o == nil {
		return nil
	}
	return o.DomainSchema
}
