ALL_LOCAL :=	data.go

include ../../../Rules.mak

data.go:
	$(IPP_REGISTRATIONS_GENERATE) -i ipp-registrations.xml -e errata.xml -o data.go -w
