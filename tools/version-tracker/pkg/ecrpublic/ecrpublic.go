package ecrpublic

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/aws/eks-anywhere-build-tooling/tools/version-tracker/pkg/constants"
	"github.com/aws/eks-anywhere-build-tooling/tools/version-tracker/pkg/util/command"
	"github.com/aws/eks-anywhere-build-tooling/tools/version-tracker/pkg/util/semver"
)

func GetLatestRevision(imageRepository, currentRevision, branchName string) (string, bool, error) {
	var latestRevision string
	var needsUpgrade bool
	currentRevisionSemver, err := semver.New(currentRevision)
	if err != nil {
		return "", false, fmt.Errorf("getting semver for current version: %v", err)
	}

	skopeoListTagsCmd := exec.Command("skopeo", "list-tags", fmt.Sprintf("docker://%s", imageRepository))
	listTagsOutput, err := command.ExecCommand(skopeoListTagsCmd)
	if err != nil {
		return "", false, fmt.Errorf("running Skopeo list-tags command: %v", err)
	}

	var tagsList interface{}
	err = json.Unmarshal([]byte(listTagsOutput), &tagsList)
	if err != nil {
		return "", false, fmt.Errorf("unmarshalling output of Skopeo list-tags command: %v", err)
	}

	imageTags := tagsList.(map[string]interface{})["Tags"].([]interface{})

	latestRevisionSemver := currentRevisionSemver
	latestRevision = currentRevision
	for _, tag := range imageTags {
		tag := tag.(string)
		if !strings.HasPrefix(tag, "v") {
			continue
		}

		tagSemver, err := semver.New(tag)
		if err != nil {
			return "", false, fmt.Errorf("getting semver for %s tag [%s]: %v", imageRepository, tag, err)
		}

		if branchName != constants.MainBranchName {
			tagRegex := regexp.MustCompile(fmt.Sprintf(`^v?%d(\.|_)%d(\.|_)\d+`, currentRevisionSemver.Major, currentRevisionSemver.Minor))
			if !tagRegex.MatchString(tag) {
				continue
			}
		}

		if tagSemver.GreaterThan(latestRevisionSemver) {
			latestRevisionSemver = tagSemver
			latestRevision = tag
			needsUpgrade = true
		}
	}

	return latestRevision, needsUpgrade, nil
}
