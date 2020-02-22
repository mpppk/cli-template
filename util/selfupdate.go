// Package selfupdate provides function to update binary
package util

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const Version = "1.2.3"
const slug = "mpppk/cli-template"

// Do execute updating binary
func Do() (bool, error) {
	v := semver.MustParse(Version)
	latest, err := selfupdate.UpdateSelf(v, slug)
	if err != nil {
		return false, fmt.Errorf("Binary update failed: %w", err)
	}
	return !latest.Version.Equals(v), nil
}
