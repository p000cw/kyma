APP_NAME = common
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)

.PHONY: build
build:
	./before-commit.sh

.PHONY: ci-pr
ci-pr: build

.PHONY: ci-master
ci-master: build

.PHONY: ci-release
ci-release: build
