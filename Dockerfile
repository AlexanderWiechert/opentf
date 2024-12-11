# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

FROM alpine:3.21.0

LABEL maintainer="OpenTF Team <opentf@opentf.org>"

COPY opentf /usr/local/bin/opentf

ENTRYPOINT ["/usr/local/bin/opentf"]
