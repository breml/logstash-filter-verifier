# Copyright (c) 2015-2016 Magnus Bäck <magnus@noun.se>

# Installation root directory. Should be left alone except for
# e.g. package installations. If you want to control the installation
# directory for normal use you should modify PREFIX instead.
DEST_DIR :=

# Installation prefix directory. Could be changed to e.g. /usr or
# /opt/logstash-filter-verifier.
PREFIX := /usr/local

.PHONY: all
all: logstash-filter-verifier

# Depend on this target to force a rebuild every time.
.FORCE:

# The Go compiler is fast and pretty good about figuring out what to
# build so we don't try to to outsmart it.
logstash-filter-verifier: .FORCE
	go get
	go build

.PHONY: clean
clean:
	rm -f logstash-filter-verifier

# To be able to build a Debian package from any commit and get a
# meaningful result, use "git describe" to find the current version
# number and compare it to the most recent entry in debian/changelog
# (which is what Debian build system uses when creating a package).
# If those versions are different, write a new entry with the current
# version.
#
# Replace hyphens with plus signs to comply with Debian policy and
# transform '-rcX' to '~rcX' so that a release candidate is considered
# older than the final release.
.PHONY: deb
deb:
	CURRENT_VERSION=$$(sed -n '1s/^[^ ]* (\([^)]*\)).*/\1/p' \
	        < debian/changelog) && \
	    ACTUAL_VERSION=$$(git describe --tags --always | \
	            sed 's/-rc/~rc/; s/-/+/g') && \
	    if [ "$$CURRENT_VERSION" != "$$ACTUAL_VERSION" ] ; then \
	        dch --force-bad-version --newversion $$ACTUAL_VERSION \
	                "Autogenerated changelog entry" ; \
	    fi
	debuild --preserve-envvar GOPATH -uc -us

.PHONY: install
install: logstash-filter-verifier
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -m 0755 --strip logstash-filter-verifier $(DESTDIR)$(PREFIX)/bin
