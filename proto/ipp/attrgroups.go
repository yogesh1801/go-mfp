// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Groups of attributes

package ipp

import "github.com/OpenPrinting/go-mfp/proto/ipp/iana"

// CUPSDeviceAttributesGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.CUPSDeviceAttributes] group.
type CUPSDeviceAttributesGroup struct{}

// Registrations returns attributes that belongs to the group
func (CUPSDeviceAttributesGroup) Registrations() map[string]*iana.DefAttr {
	return iana.CUPSDeviceAttributes
}

// CUPSPPDAttributesGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.CUPSPPDAttributes] group.
type CUPSPPDAttributesGroup struct{}

// Registrations returns attributes that belongs to the group
func (CUPSPPDAttributesGroup) Registrations() map[string]*iana.DefAttr {
	return iana.CUPSPPDAttributes
}

// CUPSPrinterClassAttributesGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.CUPSPrinterClassAttributes] group.
type CUPSPrinterClassAttributesGroup struct{}

// Registrations returns attributes that belongs to the group
func (CUPSPrinterClassAttributesGroup) Registrations() map[string]*iana.DefAttr {
	return iana.CUPSPrinterClassAttributes
}

// DocumentDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.DocumentDescription] group.
type DocumentDescriptionGroup struct{}

// Registrations returns attributes that belongs to the group
func (DocumentDescriptionGroup) Registrations() map[string]*iana.DefAttr {
	return iana.DocumentDescription
}

// DocumentStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.DocumentStatusGroup] group.
type DocumentStatusGroup struct{}

// Registrations returns attributes that belongs to the group
func (DocumentStatusGroup) Registrations() map[string]*iana.DefAttr {
	return iana.DocumentStatus
}

// DocumentTemplateGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.DocumentTemplate] group.
type DocumentTemplateGroup struct{}

// Registrations returns attributes that belongs to the group
func (DocumentTemplateGroup) Registrations() map[string]*iana.DefAttr {
	return iana.DocumentTemplate
}

// EventNotificationsGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.EventNotifications] group.
type EventNotificationsGroup struct{}

// Registrations returns attributes that belongs to the group
func (EventNotificationsGroup) Registrations() map[string]*iana.DefAttr {
	return iana.EventNotifications
}

// JobDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.JobDescriptionGroup] group.
type JobDescriptionGroup struct{}

// Registrations returns attributes that belongs to the group
func (JobDescriptionGroup) Registrations() map[string]*iana.DefAttr {
	return iana.JobDescription
}

// JobStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.JobStatus] group.
type JobStatusGroup struct{}

// Registrations returns attributes that belongs to the group
func (JobStatusGroup) Registrations() map[string]*iana.DefAttr {
	return iana.JobStatus
}

// JobTemplateGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.JobTemplate] group.
type JobTemplateGroup struct{}

// Registrations returns attributes that belongs to the group
func (JobTemplateGroup) Registrations() map[string]*iana.DefAttr {
	return iana.JobTemplate
}

// OperationGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.Operation] group.
type OperationGroup struct{}

// Registrations returns attributes that belongs to the group
func (OperationGroup) Registrations() map[string]*iana.DefAttr {
	return iana.Operation
}

// PrinterDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.PrinterDescription] group.
type PrinterDescriptionGroup struct{}

// Registrations returns attributes that belongs to the group
func (PrinterDescriptionGroup) Registrations() map[string]*iana.DefAttr {
	return iana.PrinterDescription
}

// PrinterStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.PrinterStatus] group.
type PrinterStatusGroup struct{}

// Registrations returns attributes that belongs to the group
func (PrinterStatusGroup) Registrations() map[string]*iana.DefAttr {
	return iana.PrinterStatus
}

// ResourceDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.ResourceDescription] group.
type ResourceDescriptionGroup struct{}

// Registrations returns attributes that belongs to the group
func (ResourceDescriptionGroup) Registrations() map[string]*iana.DefAttr {
	return iana.ResourceDescription
}

// ResourceStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.ResourceStatus] group.
type ResourceStatusGroup struct{}

// Registrations returns attributes that belongs to the group
func (ResourceStatusGroup) Registrations() map[string]*iana.DefAttr {
	return iana.ResourceStatus
}

// SubscriptionStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.SubscriptionStatus] group.
type SubscriptionStatusGroup struct{}

// Registrations returns attributes that belongs to the group
func (SubscriptionStatusGroup) Registrations() map[string]*iana.DefAttr {
	return iana.SubscriptionStatus
}

// SubscriptionTemplateGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.SubscriptionTemplate] group.
type SubscriptionTemplateGroup struct{}

// Registrations returns attributes that belongs to the group
func (SubscriptionTemplateGroup) Registrations() map[string]*iana.DefAttr {
	return iana.SubscriptionTemplate
}

// SystemDescriptionGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.SystemDescription] group.
type SystemDescriptionGroup struct{}

// Registrations returns attributes that belongs to the group
func (SystemDescriptionGroup) Registrations() map[string]*iana.DefAttr {
	return iana.SystemDescription
}

// SystemStatusGroup should be embedded into the IPP
// structure to indicate that it contains attributes, defined
// in the [iana.SystemStatus] group.
type SystemStatusGroup struct{}

// Registrations returns attributes that belongs to the group
func (SystemStatusGroup) Registrations() map[string]*iana.DefAttr {
	return iana.SystemStatus
}
