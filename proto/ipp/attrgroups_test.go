// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Groups of attributes

package ipp

// Here we verify that every group of attributes implement the
// AttributesGroup interface.
var (
	_ = AttributesGroup(CUPSDeviceAttributesGroup{})
	_ = AttributesGroup(CUPSPPDAttributesGroup{})
	_ = AttributesGroup(CUPSPrinterClassAttributesGroup{})
	_ = AttributesGroup(DocumentDescriptionGroup{})
	_ = AttributesGroup(DocumentStatusGroup{})
	_ = AttributesGroup(DocumentTemplateGroup{})
	_ = AttributesGroup(EventNotificationsGroup{})
	_ = AttributesGroup(JobDescriptionGroup{})
	_ = AttributesGroup(JobStatusGroup{})
	_ = AttributesGroup(JobTemplateGroup{})
	_ = AttributesGroup(OperationGroup{})
	_ = AttributesGroup(PrinterDescriptionGroup{})
	_ = AttributesGroup(PrinterStatusGroup{})
	_ = AttributesGroup(ResourceDescriptionGroup{})
	_ = AttributesGroup(ResourceStatusGroup{})
	_ = AttributesGroup(SubscriptionStatusGroup{})
	_ = AttributesGroup(SubscriptionTemplateGroup{})
	_ = AttributesGroup(SystemDescriptionGroup{})
	_ = AttributesGroup(SystemStatusGroup{})
)
