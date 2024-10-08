# avoid finding the local fake git repo by looking one level up when looking for base_directory
BASE_DIRECTORY:=$(shell git -C .. rev-parse --show-toplevel)
GIT_TAG=$(shell cat GIT_TAG)
GOLANG_VERSION=$(shell cat GOLANG_VERSION)

REPO=image-builder
REPO_OWNER=aws
EKS_ANYWHERE_CLONE_URL=https://github.com/aws/eks-anywhere.git
EKS_ANYWHERE_REPO=eks-anywhere
EKS_ANYWHERE_REPO_SPARSE_CHECKOUT=release/triggers/bundle-release/development/CLI_MAX_VERSION scripts/eksa_version.sh
EKS_ANYWHERE_RELEASE_MANIFEST_URL=https://dev-release-assets.eks-anywhere.model-rocket.aws.dev/$(if $(filter $(BRANCH_NAME),main),,$(BRANCH_NAME)/)eks-a-release.yaml
EKS_ANYWHERE_EMBED_VERSION=$(shell cd eks-anywhere && source ./scripts/eksa_version.sh && eksa-version::latest_release_version_in_manifest "$(EKS_ANYWHERE_RELEASE_MANIFEST_URL)")

BINARY_TARGET_FILES=image-builder
GO_MOD_PATHS=..
SIMPLE_CREATE_BINARIES=true
HAS_S3_ARTIFACTS=true
IMAGE_NAMES=
EXCLUDE_FROM_UPGRADE_BUILDSPEC=true

REPO_NO_CLONE=true
SKIP_CHECKSUM_VALIDATION=true
GIT_CHECKOUT_TARGET=main.go
FAKE_GIT_REPO_TARGET=.git

EXTRA_GO_LDFLAGS=-X github.com/aws/eks-anywhere-build-tooling/image-builder/cmd.version=$(GIT_TAG) \
	-X github.com/aws/eks-anywhere-build-tooling/image-builder/builder.eksaVersion=$(EKS_ANYWHERE_EMBED_VERSION) \
	-X github.com/aws/eks-anywhere-build-tooling/image-builder/builder.eksaReleaseManifest=$(EKS_ANYWHERE_RELEASE_MANIFEST_URL)

# Go 1.18 automatically embeds vcs information if there is a .git folder in the current path
# this makes builds not reproducible due a sync timestamp 
EXTRA_GOBUILD_FLAGS=-buildvcs=false


include $(BASE_DIRECTORY)/Common.mk


build: unit-test

$(REPO):
	@mkdir $@
	source $(BUILD_LIB)/common.sh && retry git clone --depth 1 --filter=blob:none --sparse -b $(BRANCH_NAME) $(EKS_ANYWHERE_CLONE_URL) $(EKS_ANYWHERE_REPO)
	git -C $(EKS_ANYWHERE_REPO) sparse-checkout set $(EKS_ANYWHERE_REPO_SPARSE_CHECKOUT) --cone --skip-checks

ifeq ($(CODEBUILD_CI),true)
ifneq ($(shell echo $(CODEBUILD_BUILD_ID) | grep "aws-staging-bundle-build"),)
$(BINARY_TARGETS): EKS_ANYWHERE_EMBED_VERSION=$(shell cat $(EKS_ANYWHERE_REPO)/$(EKS_ANYWHERE_REPO_SPARSE_CHECKOUT))
$(BINARY_TARGETS): EKS_ANYWHERE_RELEASE_MANIFEST_URL=https://anywhere-assets.eks.amazonaws.com/releases/eks-a/manifest.yaml
endif
endif

$(GATHER_LICENSES_TARGETS): | $(FAKE_GIT_REPO_TARGET)

$(FAKE_GIT_REPO_TARGET):
	@git init
	@git remote add origin https://github.com/aws/eks-anywhere-build-tooling.git

.PHONY: clean-extra
clean-extra:
	@rm -rf vendor .git $(GO_MOD_DOWNLOAD_TARGETS) $(REPO) $(EKS_ANYWHERE_REPO)

clean: clean-extra

unit-test: | $$(ENABLE_DOCKER)
	@source $(BUILD_LIB)/common.sh && build::common::use_go_version $(GOLANG_VERSION); \
	go test -v ./...


########### DO NOT EDIT #############################
# To update call: make add-generated-help-block
# This is added to help document dynamic targets and support shell autocompletion
# Run make help for a formatted help block with all targets
include Help.mk
########### END GENERATED ###########################
