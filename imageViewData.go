// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a ImageViewData Record

// ImageViewData Record
type ImageViewData struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// EceInstitutionRoutingNumber contains the routing and transit number of the institution that that creates the
	// bundle header.  Format: TTTTAAAAC, where:
	// TTTT Federal Reserve Prefix
	// AAAA ABA Institution Identifier
	// C Check Digit
	// For a number that identifies a non-financial institution: NNNNNNNNN
	EceInstitutionRoutingNumber string `json:"eceInstitutionRoutingNumber"`
	// BundleBusinessDate is the business date of the bundle.
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31

	//CycleNumber

	//EceInstitutionItemSequenceNumber

	//SecurityOriginatorName

	//SecurityAuthenticatorName

	//SecurityKeyName

	//ClippingOrigin

	//ClippingCoordinateH1

	//ClippingCoordinateH2

	//ClippingCoordinateV1

	//ClippingCoordinateV2

	//LengthImageReferenceKey

	//ImageReferenceKey

	//LengthDigitalSignature

	//DigitalSignature

	//LengthImageData

	//ImageData

}

// NewImageViewData returns a new ImageViewData with default values for non exported fields
func NewImageViewData() *ImageViewData {
	imageData := &ImageViewData{
		recordType: "52",
	}
	return imageData
}

// Parse takes the input record string and parses the ImageViewData values

// String writes the ImageViewData struct to a variable length string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
