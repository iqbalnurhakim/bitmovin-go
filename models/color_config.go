package models

import "github.com/iqbalnurhakim/bitmovin-go/bitmovintypes"

type ColorConfig struct {
	CopyChromaLocationFlag *bool                        `json:"copyChromaLocationFlag,omitempty"`
	CopyColorSpaceFlag     *bool                        `json:"copyColorSpaceFlag,omitempty"`
	CopyColorPrimariesFlag *bool                        `json:"copyColorPrimariesFlag,omitempty"`
	CopyColorRangeFlag     *bool                        `json:"copyColorRangeFlag,omitempty"`
	CopyColorTransferFlag  *bool                        `json:"copyColorTransferFlag,omitempty"`
	ChromaLocation         bitmovintypes.ChromaLocation `json:"chromaLocation,omitempty"`
	ColorTransfer          bitmovintypes.ColorTransfer  `json:"colorTransfer,omitempty"`
	ColorPrimaries         bitmovintypes.ColorPrimaries `json:"colorPrimaries,omitempty"`
	ColorSpace             bitmovintypes.ColorSpace     `json:"colorSpace,omitempty"`
	ColorRange             bitmovintypes.ColorRange     `json:"colorRange,omitempty"`
	InputColorSpace        bitmovintypes.ColorSpace     `json:"inputColorSpace,omitempty"`
	InputColorRange        bitmovintypes.ColorRange     `json:"inputColorRange,omitempty"`
}
