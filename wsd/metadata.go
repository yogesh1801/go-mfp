// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Metadata exchange definitions
//
// Specification: Devices Profile for Web Services
// https://specs.xmlsoap.org/ws/2006/02/devprof/devicesprofile.pdf

package wsd

import "github.com/alexpevzner/mfp/xmldoc"

// Dialect attribute values for ThisDevice, ThisModel and Relationship
// sections.
const (
	ThisDeviceDialect   = "http://schemas.xmlsoap.org/ws/2006/02/devprof/ThisDevice"
	ThisModelDialect    = "http://schemas.xmlsoap.org/ws/2006/02/devprof/ThisModel"
	RelationshipDialect = "http://schemas.xmlsoap.org/ws/2006/02/devprof/host"
)

// Relationship types for the needs of Metadata exchange, implemented here.
const (
	RelationshipHost = "http://schemas.xmlsoap.org/ws/2006/02/devprof/host"
)

// Metadata is the device description, returned as response to
// the [Get] request.
type Metadata struct {
	ThisDevice   ThisDeviceMetadata // Device description
	ThisModel    ThisModelMetadata  // Model description
	Relationship Relationship       // Host and hosted services
}

// ThisDeviceMetadata contains information about the particular device.
type ThisDeviceMetadata struct {
	FriendlyName    LocalizedStringList // Device user-friendly name
	FirmwareVersion string              // Firmware version
	SerialNumber    string              // Serial number
}

// ThisModelMetadata contains information about the model.
type ThisModelMetadata struct {
	Manufacturer    LocalizedStringList // Manufacturer name
	ManufacturerURL string              // Manufacturer URL
	ModelName       LocalizedStringList // Model name
	ModelNumber     string              // Model number
	ModelURL        string              // Model URL
	PresentationURL string              // HTML page for this model
}

// Relationship defines relationship between host (i.e., the device)
// and hosted services (i.e., print/scan serviced, implemented by
// the device).
type Relationship struct {
	Host   *ServiceMetadata  // The host inself (very optional)
	Hosted []ServiceMetadata // Hosted services
}

// ServiceMetadata contains information about the host or hosted service.
type ServiceMetadata struct {
	EndpointReference []EndpointReference // Service endpoints
	Types             Types               // Service types
	ServiceID         AnyURI              // Service identifier
}

// ToXML generates XML tree for Metadata.
func (md Metadata) ToXML() xmldoc.Element {
	// Generate sections
	thisDevice := md.ThisDevice.ToXML()
	thisModel := md.ThisModel.ToXML()
	relationship := md.Relationship.ToXML()

	// Build metadata
	metadata := xmldoc.Element{
		Name: NsMex + ":Metadata",
		Children: []xmldoc.Element{
			thisDevice,
			thisModel,
			relationship,
		},
	}

	return metadata
}

// ToXML generates XML tree for ThisDeviceMetadata
func (thisdev ThisDeviceMetadata) ToXML() xmldoc.Element {
	data := xmldoc.Element{
		Name: NsDevprof + ":ThisDevice",
	}

	for _, fn := range thisdev.FriendlyName {
		data.Children = append(data.Children,
			fn.ToXML(NsDevprof+":FriendlyName"))
	}

	data.Children = append(data.Children,
		xmldoc.Element{
			Name: NsDevprof + ":FirmwareVersion",
			Text: thisdev.FirmwareVersion,
		})

	data.Children = append(data.Children,
		xmldoc.Element{
			Name: NsDevprof + ":SerialNumber",
			Text: thisdev.SerialNumber,
		})

	thisDevice := xmldoc.Element{
		Name: NsMex + ":MetadataSection",
		Attrs: []xmldoc.Attr{{
			Name:  "Dialect",
			Value: ThisDeviceDialect,
		}},
		Children: []xmldoc.Element{data},
	}

	return thisDevice
}

// ToXML generates XML tree for ThisModelMetadata
func (thismdl ThisModelMetadata) ToXML() xmldoc.Element {
	data := xmldoc.Element{
		Name: NsDevprof + ":ThisModel",
	}

	for _, mfg := range thismdl.Manufacturer {
		data.Children = append(data.Children,
			mfg.ToXML(NsDevprof+":Manufacturer"))
	}

	if thismdl.ManufacturerURL != "" {
		data.Children = append(data.Children,
			xmldoc.Element{
				Name: NsDevprof + ":ManufacturerUrl",
				Text: thismdl.ManufacturerURL,
			})
	}

	for _, mdl := range thismdl.ModelName {
		data.Children = append(data.Children,
			mdl.ToXML(NsDevprof+":ModelName"))
	}

	data.Children = append(data.Children,
		xmldoc.Element{
			Name: NsDevprof + ":ModelNumber",
			Text: thismdl.ModelNumber,
		})

	if thismdl.ModelURL != "" {
		data.Children = append(data.Children,
			xmldoc.Element{
				Name: NsDevprof + ":ModelUrl",
				Text: thismdl.ModelURL,
			})
	}

	if thismdl.PresentationURL != "" {
		data.Children = append(data.Children,
			xmldoc.Element{
				Name: NsDevprof + ":PresentationUrl",
				Text: thismdl.PresentationURL,
			})
	}

	thisModel := xmldoc.Element{
		Name: NsMex + ":MetadataSection",
		Attrs: []xmldoc.Attr{{
			Name:  "Dialect",
			Value: ThisModelDialect,
		}},
		Children: []xmldoc.Element{data},
	}

	return thisModel
}

// ToXML generates XML tree for Relationship
func (rel Relationship) ToXML() xmldoc.Element {
	root := xmldoc.Element{
		Name: NsDevprof + ":Relationship",
		Attrs: []xmldoc.Attr{{
			Name:  "Type",
			Value: RelationshipHost,
		}},
	}

	if rel.Host != nil {
		root.Children = append(root.Children,
			rel.Host.ToXML(NsDevprof+"Host"))
	}

	for _, hosted := range rel.Hosted {
		root.Children = append(root.Children,
			hosted.ToXML(NsDevprof+"Hosted"))
	}

	return root
}

// ToXML generates XML tree for the ServiceMetadata
func (svcmeta ServiceMetadata) ToXML(name string) xmldoc.Element {
	elm := xmldoc.Element{
		Name: name,
	}

	for _, ep := range svcmeta.EndpointReference {
		elm.Children = append(elm.Children,
			ep.ToXML("EndpointReference"))
	}

	elm.Children = append(elm.Children, svcmeta.Types.ToXML())

	if svcmeta.ServiceID != "" {
		elm.Children = append(elm.Children,
			xmldoc.Element{
				Name: NsDevprof + ":ServiceId",
				Text: string(svcmeta.ServiceID),
			})
	}

	return elm
}
