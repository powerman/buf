// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufimagemodify

import (
	"context"
	"testing"

	"github.com/powerman/buf/internal/buf/bufcore/bufimage"
	"github.com/powerman/buf/internal/buf/bufcore/bufimage/bufimagebuild"
	"github.com/powerman/buf/internal/buf/bufcore/bufmodule"
	"github.com/powerman/buf/internal/buf/bufcore/bufmodule/bufmodulebuild"
	"github.com/powerman/buf/internal/buf/bufcore/bufmodule/bufmoduletesting"
	"github.com/powerman/buf/internal/pkg/storage/storageos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const (
	testImportPathPrefix = "github.com/foo/bar/internal/gen/proto/go"
	testRepositoryOwner  = "testowner"
	testRepositoryName   = "testrepository"
)

func assertFileOptionSourceCodeInfoEmpty(t *testing.T, image bufimage.Image, fileOptionPath []int32, includeSourceInfo bool) {
	t.Helper()
	for _, imageFile := range image.Files() {
		descriptor := imageFile.Proto()

		if !includeSourceInfo {
			assert.Empty(t, descriptor.SourceCodeInfo)
			continue
		}

		var hasFileOption bool
		for _, location := range descriptor.SourceCodeInfo.Location {
			if len(location.Path) > 0 && int32SliceIsEqual(location.Path, fileOptionPath) {
				hasFileOption = true
				break
			}
		}
		assert.False(t, hasFileOption)
	}
}

func assertFileOptionSourceCodeInfoNotEmpty(t *testing.T, image bufimage.Image, fileOptionPath []int32) {
	t.Helper()
	for _, imageFile := range image.Files() {
		descriptor := imageFile.Proto()

		var hasFileOption bool
		for _, location := range descriptor.SourceCodeInfo.Location {
			if len(location.Path) > 0 && int32SliceIsEqual(location.Path, fileOptionPath) {
				hasFileOption = true
				break
			}
		}
		assert.True(t, hasFileOption)
	}
}

func testGetImage(t *testing.T, dirPath string, includeSourceInfo bool) bufimage.Image {
	t.Helper()
	moduleFileSet := testGetModuleFileSet(t, dirPath)
	var options []bufimagebuild.BuildOption
	if !includeSourceInfo {
		options = append(options, bufimagebuild.WithExcludeSourceCodeInfo())
	}
	image, _, err := bufimagebuild.NewBuilder(zap.NewNop()).Build(
		context.Background(),
		moduleFileSet,
		options...,
	)
	require.NoError(t, err)
	return image
}

func testGetModuleFileSet(t *testing.T, dirPath string) bufmodule.ModuleFileSet {
	t.Helper()
	storageosProvider := storageos.NewProvider()
	readWriteBucket, err := storageosProvider.NewReadWriteBucket(
		dirPath,
	)
	require.NoError(t, err)
	moduleCommit, err := bufmodule.NewModuleCommit(
		"modulerepo.internal",
		testRepositoryOwner,
		testRepositoryName,
		bufmoduletesting.TestCommit,
	)
	require.NoError(t, err)
	module, err := bufmodule.NewModuleForBucket(
		context.Background(),
		readWriteBucket,
		bufmodule.ModuleWithModuleCommit(moduleCommit),
	)
	require.NoError(t, err)
	moduleFileSet, err := bufmodulebuild.NewModuleFileSetBuilder(
		zap.NewNop(),
		bufmodule.NewNopModuleReader(),
	).Build(
		context.Background(),
		module,
	)
	require.NoError(t, err)
	return moduleFileSet
}
