// Package selfupdate provides function to update binary
package selfupdate

import (
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"golang.org/x/xerrors"
)

const Version = "1.2.3"
const slug = "mpppk/cli-template"

// Do execute updating binary
func Do() (bool, error) {
	v := semver.MustParse(Version)
	latest, err := selfupdate.UpdateSelf(v, slug)
	if err != nil {
		return false, xerrors.Errorf("Binary update failed: %w", err)
	}
	return !latest.Version.Equals(v), nil
}
