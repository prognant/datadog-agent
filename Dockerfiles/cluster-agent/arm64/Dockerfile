########################################
# Preparation stage: layout and chmods #
########################################

FROM ubuntu:21.04 as builder

WORKDIR /output

COPY datadog-cluster-agent.arm64 opt/datadog-agent/bin/datadog-cluster-agent
COPY ./conf.d etc/datadog-agent/conf.d
COPY ./datadog-cluster.yaml etc/datadog-agent/datadog-cluster.yaml
COPY ./security-agent-policies/compliance/containers/ etc/datadog-agent/compliance.d
COPY entrypoint.sh .
COPY readsecret.sh readsecret_multiple_providers.sh ./

RUN chmod 755 entrypoint.sh \
    && chown root:root readsecret.sh readsecret_multiple_providers.sh \
    && chmod 550 readsecret.sh readsecret_multiple_providers.sh\
    && chmod g+r,g+w,g+X -R etc/datadog-agent/ \
    && chmod +x opt/datadog-agent/bin/datadog-cluster-agent \
    && ln -s /opt/datadog-agent/bin/datadog-cluster-agent opt/datadog-agent/bin/agent

####################################
# Actual docker image construction #
####################################

FROM ubuntu:21.04

LABEL maintainer "Datadog <package@datadoghq.com>"

ENV PATH="/opt/datadog-agent/bin/:$PATH" \
    DOCKER_DD_AGENT="true" \
    # Allow User Group to exec the secret backend script.
    DD_SECRET_BACKEND_COMMAND_ALLOW_GROUP_EXEC_PERM="true"

RUN apt-get update \
    && apt full-upgrade -y \
    && apt-get install --no-install-recommends -y ca-certificates curl \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY --from=builder /output /

# Allow running as an unprivileged user:
# - General case is the dd-agent user
# - OpenShift uses a random UID in the root group
#
# Containerd does not preserve permissions when mounting a volume on top
# of an empty folder. Creating .placeholder files as a workaround.
#
RUN adduser --system --no-create-home --disabled-password --ingroup root dd-agent \
    && mkdir -p /var/log/datadog/ /conf.d \
    && touch /var/log/datadog/.placeholder \
    && touch /tmp/.placeholder \
    && chown -R dd-agent:root /etc/datadog-agent/ /var/log/datadog/ /conf.d /tmp/ \
    && chmod g+r,g+w,g+X -R /etc/datadog-agent/ /var/log/datadog/ /conf.d /tmp/

# Incompatible with the custom metrics API on port 443
# Set DD_EXTERNAL_METRICS_PROVIDER_PORT to a higher value to run as non-root
# USER dd-agent

# Leave following directories RW to allow use of readonly rootfs
VOLUME ["/etc/datadog-agent", "/var/log/datadog", "/tmp"]

ENTRYPOINT ["/entrypoint.sh"]

# No docker healthcheck, use a HTTP check
# on port 5005 and/or 443 on Kubernetes

CMD ["datadog-cluster-agent", "start"]
