FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/sample-controller-operatorsdk-take2 \
    USER_UID=1001 \
    USER_NAME=sample-controller-operatorsdk-take2

# install operator binary
COPY build/_output/bin/sample-controller-operatorsdk-take2 ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
