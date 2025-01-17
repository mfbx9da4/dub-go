// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

type TrackCustomerGlobals struct {
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	WorkspaceID *string `queryParam:"style=form,explode=true,name=workspaceId"`
}

func (o *TrackCustomerGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type TrackCustomerRequestBody struct {
	// This is the unique identifier for the customer in the client's app. This is used to track the customer's journey.
	CustomerID string `json:"customerId"`
	// Name of the customer in the client's app.
	CustomerName *string `json:"customerName,omitempty"`
	// Email of the customer in the client's app.
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// Avatar of the customer in the client's app.
	CustomerAvatar *string `json:"customerAvatar,omitempty"`
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
