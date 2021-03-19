# Copyright (c) 2020 Gitpod GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License-AGPL.txt in the project root for license information.

FROM alpine:3.13

# Ensure latest packages are present, like security updates.
RUN  apk upgrade --no-cache \
  && apk add --no-cache git bash ca-certificates

COPY components-image-builder--app/image-builder /app/
RUN chmod +x /app/image-builder

COPY components-image-builder-workspace-image-layer--pack/pack.tar /app/workspace-image-layer.tar.gz

RUN touch /app/entrypoint.sh
RUN echo -e "#!/bin/bash\n set -x\n sleep 15 \n /app/image-builder \$@" > /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh
ENTRYPOINT ["/app/entrypoint.sh"]
CMD [ "-v", "help" ]