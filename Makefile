GOTARGETS=
GOLIBRARYTARGET=./...
GOTESTFLAGS=-v

default: test

include go-common.mk

post-testcover::
	sed -i.orig -e 's!/v3/!/!' $(GOTESTCOVERRAW) $(GOTESTCOVERHTML)
