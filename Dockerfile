FROM cgr.dev/chainguard/static@sha256:288b818c1b3dd89776d176f07f5f671b118fe836c4d80ec2cc3299b596fe71b7
COPY imgverter \
	/usr/bin/imgverter
ENTRYPOINT ["/usr/bin/automated-word-translator"]