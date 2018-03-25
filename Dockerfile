FROM alpine
ADD aws-madness.linux /aws-madness
ADD scripts/docker-entrypoint.sh /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
