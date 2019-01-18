package lib

import (
	"github.com/blang/semver"
	"github.com/pkg/errors"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const Version = "1.2.3"
const slug = "mpppk/cli-template"

func DoSelfUpdate() (bool, error) {
	v := semver.MustParse(Version)
	latest, err := selfupdate.UpdateSelf(v, slug)
	return !latest.Version.Equals(v), errors.Wrap(err, "Binary update failed")
}