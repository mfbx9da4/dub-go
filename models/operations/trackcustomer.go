// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

type TrackCustomerRequestBody struct {
	// This is the unique identifier for the customer in the client's app. This is used to track the customer's journey.
	CustomerID string `json:"customerId"`
	// Name of the customer in the client's app.
	CustomerName *string `default:"null" json:"customerName"`
	// Email of the customer in the client's app.
	CustomerEmail *string `default:"null" json:"customerEmail"`
	// Avatar of the customer in the client's app.
	CustomerAvatar *string `default:"null" json:"customerAvatar"`
}

func (t TrackCustomerRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TrackCustomerRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TrackCustomerRequestBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *TrackCustomerRequestBody) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *TrackCustomerRequestBody) GetCustomerEmail() *string {
	if o == nil {
		return nil
	}
	return o.CustomerEmail
}

func (o *TrackCustomerRequestBody) GetCustomerAvatar() *string {
	if o == nil {
		return nil
	}
	return o.CustomerAvatar
}

// TrackCustomerResponseBody - A customer was tracked.
type TrackCustomerResponseBody struct {
	CustomerID     string  `json:"customerId"`
	CustomerName   *string `json:"customerName"`
	CustomerEmail  *string `json:"customerEmail"`
	CustomerAvatar *string `json:"customerAvatar"`
}

func (o *TrackCustomerResponseBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *TrackCustomerResponseBody) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *TrackCustomerResponseBody) GetCustomerEmail() *string {
	if o == nil {
		return nil
	}
	return o.CustomerEmail
}

func (o *TrackCustomerResponseBody) GetCustomerAvatar() *string {
	if o == nil {
		return nil
	}
	return o.CustomerAvatar
}

type TrackCustomerResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A customer was tracked.
	Object *TrackCustomerResponseBody
}

func (o *TrackCustomerResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TrackCustomerResponse) GetObject() *TrackCustomerResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}