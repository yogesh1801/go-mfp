ALL_LOCAL :=	data.go

include ../../../Rules.mak

data.go: ipp-registrations.xml cups.xml errata.xml
	$(IPP_REGISTRATIONS_GENERATE) -i ipp-registrations.xml -i cups.xml -e errata.xml -o data.go
