// Copyright 2022 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

//go:build freebsd || linux || netbsd || openbsd
// +build freebsd linux netbsd openbsd

package ec2detector

import (
	"testing"

	"github.com/aws/amazon-ssm-agent/common/identity/availableidentities/ec2/ec2detector/helper"
	"github.com/aws/amazon-ssm-agent/common/identity/availableidentities/ec2/ec2detector/nitrodetector"
	"github.com/aws/amazon-ssm-agent/common/identity/availableidentities/ec2/ec2detector/xendetector"
	"github.com/stretchr/testify/assert"
)

func TestAssertDetectorSize(t *testing.T) {
	detectors := helper.GetAllDetectors()
	assert.Equal(t, 2, len(detectors))

	// assert priority order
	assert.Equal(t, nitrodetector.Name, detectors[0].GetName())
	assert.Equal(t, xendetector.Name, detectors[1].GetName())
}
