FROM scratch
MAINTAINER Kel Cecil <kelcecil@praisechaos.com>

USER nobody

COPY echoexit /bin/echoexit
ENTRYPOINT /bin/echoexit