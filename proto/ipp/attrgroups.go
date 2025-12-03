// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Groups of attributes

package ipp

import "github.com/OpenPrinting/go-mfp/proto/ipp/iana"

// AttributesGroup one of the following types.
//
//	[CUPSDeviceAttributesGroup]
//	[CUPSPPDAttributesGroup]
//	[CUPSPrinterClassAttributesGroup]
//	[DocumentDescriptionGroup]
//	[DocumentStatusGroup]
//	[DocumentTemplateGroup]
//	[EventNotificationsGroup]
//	[JobDescriptionGroup]
//	[JobStatusGroup]
//	[JobTemplateGroup]
//	[OperationGroup]
//	[PrinterDescriptionGroup]
//	[PrinterStatusGroup]
//	[ResourceDescriptionGroup]
//	[ResourceStatusGroup]
//	[SubscriptionStatusGroup]
//	[SubscriptionTemplateGroup]
//	[SystemDescriptionGroup]
//	[SystemStatusGroup]
//
// The member of this type needs to be embedded into each structure
// that implements the [Object] interface to specify which IPP attributes
// are registered by IANA for that type of Object.
type AttributesGroup interface {
	// registrations returns attributes that belongs to the group
	registrations() map[string]*iana.DefAttr
}

// CUPSDeviceAttributesGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.CUPSDeviceAttributes] group.
type CUPSDeviceAttributesGroup struct{}

// registrations returns attributes that belongs to the group
func (CUPSDeviceAttributesGroup) registrations() map[string]*iana.DefAttr {
	return iana.CUPSDeviceAttributes
}

// CUPSPPDAttributesGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.CUPSPPDAttributes] group.
type CUPSPPDAttributesGroup struct{}

// registrations returns attributes that belongs to the group
func (CUPSPPDAttributesGroup) registrations() map[string]*iana.DefAttr {
	return iana.CUPSPPDAttributes
}

// CUPSPrinterClassAttributesGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.CUPSPrinterClassAttributes] group.
type CUPSPrinterClassAttributesGroup struct{}

// registrations returns attributes that belongs to the group
func (CUPSPrinterClassAttributesGroup) registrations() map[string]*iana.DefAttr {
	return iana.CUPSPrinterClassAttributes
}

// DocumentDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.DocumentDescription] group.
type DocumentDescriptionGroup struct{}

// registrations returns attributes that belongs to the group
func (DocumentDescriptionGroup) registrations() map[string]*iana.DefAttr {
	return iana.DocumentDescription
}

// DocumentStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.DocumentStatusGroup] group.
type DocumentStatusGroup struct{}

// registrations returns attributes that belongs to the group
func (DocumentStatusGroup) registrations() map[string]*iana.DefAttr {
	return iana.DocumentStatus
}

// DocumentTemplateGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.DocumentTemplate] group.
type DocumentTemplateGroup struct{}

// registrations returns attributes that belongs to the group
func (DocumentTemplateGroup) registrations() map[string]*iana.DefAttr {
	return iana.DocumentTemplate
}

// EventNotificationsGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.EventNotifications] group.
type EventNotificationsGroup struct{}

// registrations returns attributes that belongs to the group
func (EventNotificationsGroup) registrations() map[string]*iana.DefAttr {
	return iana.EventNotifications
}

// JobDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.JobDescriptionGroup] group.
type JobDescriptionGroup struct{}

// registrations returns attributes that belongs to the group
func (JobDescriptionGroup) registrations() map[string]*iana.DefAttr {
	return iana.JobDescription
}

// JobStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.JobStatus] group.
type JobStatusGroup struct{}

// registrations returns attributes that belongs to the group
func (JobStatusGroup) registrations() map[string]*iana.DefAttr {
	return iana.JobStatus
}

// JobTemplateGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.JobTemplate] group.
type JobTemplateGroup struct{}

// registrations returns attributes that belongs to the group
func (JobTemplateGroup) registrations() map[string]*iana.DefAttr {
	return iana.JobTemplate
}

// OperationGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.Operation] group.
type OperationGroup struct{}

// registrations returns attributes that belongs to the group
func (OperationGroup) registrations() map[string]*iana.DefAttr {
	return iana.Operation
}

// PrinterDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.PrinterDescription] group.
type PrinterDescriptionGroup struct{}

// registrations returns attributes that belongs to the group
func (PrinterDescriptionGroup) registrations() map[string]*iana.DefAttr {
	return iana.PrinterDescription
}

// PrinterStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.PrinterStatus] group.
type PrinterStatusGroup struct{}

// registrations returns attributes that belongs to the group
func (PrinterStatusGroup) registrations() map[string]*iana.DefAttr {
	return iana.PrinterStatus
}

// ResourceDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.ResourceDescription] group.
type ResourceDescriptionGroup struct{}

// registrations returns attributes that belongs to the group
func (ResourceDescriptionGroup) registrations() map[string]*iana.DefAttr {
	return iana.ResourceDescription
}

// ResourceStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.ResourceStatus] group.
type ResourceStatusGroup struct{}

// registrations returns attributes that belongs to the group
func (ResourceStatusGroup) registrations() map[string]*iana.DefAttr {
	return iana.ResourceStatus
}

// SubscriptionStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.SubscriptionStatus] group.
type SubscriptionStatusGroup struct{}

// registrations returns attributes that belongs to the group
func (SubscriptionStatusGroup) registrations() map[string]*iana.DefAttr {
	return iana.SubscriptionStatus
}

// SubscriptionTemplateGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.SubscriptionTemplate] group.
type SubscriptionTemplateGroup struct{}

// registrations returns attributes that belongs to the group
func (SubscriptionTemplateGroup) registrations() map[string]*iana.DefAttr {
	return iana.SubscriptionTemplate
}

// SystemDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.SystemDescription] group.
type SystemDescriptionGroup struct{}

// registrations returns attributes that belongs to the group
func (SystemDescriptionGroup) registrations() map[string]*iana.DefAttr {
	return iana.SystemDescription
}

// SystemStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.SystemStatus] group.
type SystemStatusGroup struct{}

// registrations returns attributes that belongs to the group
func (SystemStatusGroup) registrations() map[string]*iana.DefAttr {
	return iana.SystemStatus
}
