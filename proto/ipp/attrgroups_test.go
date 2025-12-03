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
	_ = attributesGroup(CUPSDeviceAttributesGroup{})
	_ = attributesGroup(CUPSPPDAttributesGroup{})
	_ = attributesGroup(CUPSPrinterClassAttributesGroup{})
	_ = attributesGroup(DocumentDescriptionGroup{})
	_ = attributesGroup(DocumentStatusGroup{})
	_ = attributesGroup(DocumentTemplateGroup{})
	_ = attributesGroup(EventNotificationsGroup{})
	_ = attributesGroup(JobDescriptionGroup{})
	_ = attributesGroup(JobStatusGroup{})
	_ = attributesGroup(JobTemplateGroup{})
	_ = attributesGroup(OperationGroup{})
	_ = attributesGroup(PrinterDescriptionGroup{})
	_ = attributesGroup(PrinterStatusGroup{})
	_ = attributesGroup(ResourceDescriptionGroup{})
	_ = attributesGroup(ResourceStatusGroup{})
	_ = attributesGroup(SubscriptionStatusGroup{})
	_ = attributesGroup(SubscriptionTemplateGroup{})
	_ = attributesGroup(SystemDescriptionGroup{})
	_ = attributesGroup(SystemStatusGroup{})
)
