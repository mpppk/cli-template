// Package util provides some utilities
package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/comail/colog"
)

// ConvertStringSliceToIntSlice converts string slices to int slices.
func ConvertStringSliceToIntSlice(stringSlice []string) (intSlice []int, err error) {
	for _, s := range stringSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string slice to int slice: %w", err)
		}
		intSlice = append(intSlice, num)
	}
	return
}

// InitializeLog initialize log settings
func InitializeLog(verbose bool) {
	colog.Register()
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LInfo)

	if verbose {
		colog.SetMinLevel(colog.LDebug)
	}
}

// PrettyPrintError print error as pretty
func PrettyPrintError(err error) string {
	var errs []error
	for {
		errs = append(errs, err)
		if e := errors.Unwrap(err); e == nil {
			break
		} else {
			err = e
		}
	}
	var eMsgs []string
	beforeErrMsg := ""
	for i := len(errs) - 1; i >= 0; i-- {
		e := errs[i]
		eMsg := ""
		if beforeErrMsg == "" {
			eMsg = e.Error()
		} else {
			eMsgs := strings.Split(e.Error(), beforeErrMsg)
			eMsg = eMsgs[0]
		}
		eMsgs = append(eMsgs, eMsg)
		beforeErrMsg = e.Error()
	}

	retMsg := ""
	for i := len(eMsgs) - 1; i >= 0; i-- {
		prefix := "  "
		if i == len(eMsgs)-1 {
			prefix = "Error: "
		}
		retMsg = retMsg + fmt.Sprintln(prefix+strings.TrimSuffix(eMsgs[i], ": "))
	}
	return retMsg
}
