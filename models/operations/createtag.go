// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
)

type CreateTagGlobals struct {
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	WorkspaceID *string `queryParam:"style=form,explode=true,name=workspaceId"`
}

func (o *CreateTagGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

// Color - The color of the tag. If not provided, a random color will be used from the list: red, yellow, green, blue, purple, pink, brown.
type Color string

const (
	ColorRed    Color = "red"
	ColorYellow Color = "yellow"
	ColorGreen  Color = "green"
	ColorBlue   Color = "blue"
	ColorPurple Color = "purple"
	ColorPink   Color = "pink"
	ColorBrown  Color = "brown"
)

func (e Color) ToPointer() *Color {
	return &e
}
func (e *Color) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "red":
		fallthrough
	case "yellow":
		fallthrough
	case "green":
		fallthrough
	case "blue":
		fallthrough
	case "purple":
		fallthrough
	case "pink":
		fallthrough
	case "brown":
		*e = Color(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Color: %v", v)
	}
}

type CreateTagRequestBody struct {
	// The name of the tag to create.
	Name *string `json:"name,omitempty"`
	// The color of the tag. If not provided, a random color will be used from the list: red, yellow, green, blue, purple, pink, brown.
	Color *Color `json:"color,omitempty"`
	// The name of the tag to create.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Tag *string `json:"tag,omitempty"`
}

func (o *CreateTagRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateTagRequestBody) GetColor() *Color {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *CreateTagRequestBody) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}
