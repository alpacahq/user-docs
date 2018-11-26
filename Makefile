.PHONY: build publish

build:
	hugo

publish:
	aws s3 sync --acl public-read public $(DEPLOYMENT_TARGET)