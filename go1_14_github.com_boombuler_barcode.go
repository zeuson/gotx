// Code generated by 'goexports github.com/boombuler/barcode'. DO NOT EDIT.

// +build go1.14,!go1.15

package main

import (
	"github.com/boombuler/barcode"
	"image"
	"image/color"
	"reflect"
)

func init() {
	GotxSymbols["github.com/boombuler/barcode"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Scale":               reflect.ValueOf(barcode.Scale),
		"Type2of5":            reflect.ValueOf(barcode.Type2of5),
		"Type2of5Interleaved": reflect.ValueOf(barcode.Type2of5Interleaved),
		"TypeAztec":           reflect.ValueOf(barcode.TypeAztec),
		"TypeCodabar":         reflect.ValueOf(barcode.TypeCodabar),
		"TypeCode128":         reflect.ValueOf(barcode.TypeCode128),
		"TypeCode39":          reflect.ValueOf(barcode.TypeCode39),
		"TypeCode93":          reflect.ValueOf(barcode.TypeCode93),
		"TypeDataMatrix":      reflect.ValueOf(barcode.TypeDataMatrix),
		"TypeEAN13":           reflect.ValueOf(barcode.TypeEAN13),
		"TypeEAN8":            reflect.ValueOf(barcode.TypeEAN8),
		"TypePDF":             reflect.ValueOf(barcode.TypePDF),
		"TypeQR":              reflect.ValueOf(barcode.TypeQR),

		// type definitions
		"Barcode":      reflect.ValueOf((*barcode.Barcode)(nil)),
		"BarcodeIntCS": reflect.ValueOf((*barcode.BarcodeIntCS)(nil)),
		"Metadata":     reflect.ValueOf((*barcode.Metadata)(nil)),

		// interface wrapper definitions
		"_Barcode":      reflect.ValueOf((*_github_com_boombuler_barcode_Barcode)(nil)),
		"_BarcodeIntCS": reflect.ValueOf((*_github_com_boombuler_barcode_BarcodeIntCS)(nil)),
	}
}

// _github_com_boombuler_barcode_Barcode is an interface wrapper for Barcode type
type _github_com_boombuler_barcode_Barcode struct {
	WAt         func(x int, y int) color.Color
	WBounds     func() image.Rectangle
	WColorModel func() color.Model
	WContent    func() string
	WMetadata   func() barcode.Metadata
}

func (W _github_com_boombuler_barcode_Barcode) At(x int, y int) color.Color { return W.WAt(x, y) }
func (W _github_com_boombuler_barcode_Barcode) Bounds() image.Rectangle     { return W.WBounds() }
func (W _github_com_boombuler_barcode_Barcode) ColorModel() color.Model     { return W.WColorModel() }
func (W _github_com_boombuler_barcode_Barcode) Content() string             { return W.WContent() }
func (W _github_com_boombuler_barcode_Barcode) Metadata() barcode.Metadata  { return W.WMetadata() }

// _github_com_boombuler_barcode_BarcodeIntCS is an interface wrapper for BarcodeIntCS type
type _github_com_boombuler_barcode_BarcodeIntCS struct {
	WAt         func(x int, y int) color.Color
	WBounds     func() image.Rectangle
	WCheckSum   func() int
	WColorModel func() color.Model
	WContent    func() string
	WMetadata   func() barcode.Metadata
}

func (W _github_com_boombuler_barcode_BarcodeIntCS) At(x int, y int) color.Color { return W.WAt(x, y) }
func (W _github_com_boombuler_barcode_BarcodeIntCS) Bounds() image.Rectangle     { return W.WBounds() }
func (W _github_com_boombuler_barcode_BarcodeIntCS) CheckSum() int               { return W.WCheckSum() }
func (W _github_com_boombuler_barcode_BarcodeIntCS) ColorModel() color.Model     { return W.WColorModel() }
func (W _github_com_boombuler_barcode_BarcodeIntCS) Content() string             { return W.WContent() }
func (W _github_com_boombuler_barcode_BarcodeIntCS) Metadata() barcode.Metadata  { return W.WMetadata() }