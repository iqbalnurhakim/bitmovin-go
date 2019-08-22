package models

import (
	"github.com/iqbalnurhakim/bitmovin-go/bitmovintypes"
)

type Filter struct {
	ID          *string                `json:"id"`
	Name        *string                `json:"name"`
	Description *string                `json:"description"`
	CustomData  map[string]interface{} `json:"customData"`
}

type WatermarkFilter struct {
	Filter
	Image  *string `json:"image"`
	Left   *int64  `json:"left"`
	Right  *int64  `json:"right"`
	Top    *int64  `json:"top"`
	Bottom *int64  `json:"bottom"`
}

type CropFilter struct {
	Filter
	Left   *int64 `json:"left"`
	Right  *int64 `json:"right"`
	Top    *int64 `json:"top"`
	Bottom *int64 `json:"bottom"`
}

type RotationFilter struct {
	Filter
	Rotation *int64 `json:"rotation"`
}

type DenoiseFilter struct {
	Filter
	LumaSpatial   *float64 `json:"lumaSpatial,omitempty"`
	ChromaSpatial *float64 `json:"chromaSpatial,omitempty"`
	LumaTmp       *float64 `json:"lumaTmp,omitempty"`
	ChromaTmp     *float64 `json:"chromaTmp,omitempty"`
}

type DenoiseFilterResponse struct {
	RequestID *string                      `json:"requestId,omitempty"`
	Status    bitmovintypes.ResponseStatus `json:"status,omitempty"`
	Data      DenoiseFilterData            `json:"data,omitempty"`
}

type DenoiseFilterData struct {
	Result           DenoiseFilter `json:"result,omitempty"`
	Messages         []Message     `json:"messages,omitempty"`
	Code             *int64        `json:"code,omitempty"`
	Message          *string       `json:"message,omitempty"`
	DeveloperMessage *string       `json:"developerMessage,omitempty"`
	Links            []Link        `json:"links,omitempty"`
	Details          []Detail      `json:"details,omitempty"`
}

type DeinterlacingFilter struct {
	Filter
	Mode   bitmovintypes.DeinterlacingMode  `json:"mode"`
	Parity bitmovintypes.PictureFieldParity `json:"parity"`
}

type DeinterlacingFilterResponse struct {
	RequestID *string                      `json:"requestId,omitempty"`
	Status    bitmovintypes.ResponseStatus `json:"status,omitempty"`
	Data      DeinterlacingFilterData      `json:"data,omitempty"`
}

type DeinterlacingFilterData struct {
	Result           DeinterlacingFilter `json:"result,omitempty"`
	Messages         []Message           `json:"messages,omitempty"`
	Code             *int64              `json:"code,omitempty"`
	Message          *string             `json:"message,omitempty"`
	DeveloperMessage *string             `json:"developerMessage,omitempty"`
	Links            []Link              `json:"links,omitempty"`
	Details          []Detail            `json:"details,omitempty"`
}

type AddFilter struct {
	ID       string `json:"id"`
	Position *int64 `json:"position"`
}
type AddFilterResponseData struct {
	Result           AddFilter `json:"result,omitempty"`
	Messages         []Message `json:"messages,omitempty"`
	Code             *int64    `json:"code,omitempty"`
	Message          *string   `json:"message,omitempty"`
	DeveloperMessage *string   `json:"developerMessage,omitempty"`
	Links            []Link    `json:"links,omitempty"`
	Details          []Detail  `json:"details,omitempty"`
}
type AddFilterResponse struct {
	RequestID *string                      `json:"requestId,omitempty"`
	Status    bitmovintypes.ResponseStatus `json:"status,omitempty"`
	Data      AddFilterResponseData        `json:"data,omitempty"`
}
