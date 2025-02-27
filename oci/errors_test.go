// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// issue-routing-tag: terraform/default
func TestUnitGetVersionAndDateError(t *testing.T) {
	versionError := getVersionAndDateError()
	assert.Contains(t, versionError, "Provider version: ")
	assert.Contains(t, versionError, Version)
	assert.Contains(t, versionError, ReleaseDate)
	assert.NotContains(t, versionError, "Update(s) behind to current")
}
