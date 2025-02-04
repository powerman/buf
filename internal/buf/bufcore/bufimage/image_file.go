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

package bufimage

import (
	"github.com/powerman/buf/internal/buf/bufcore"
	"github.com/powerman/buf/internal/buf/bufcore/bufmodule"
	"github.com/powerman/buf/internal/pkg/protodescriptor"
	"google.golang.org/protobuf/types/descriptorpb"
)

var _ ImageFile = &imageFile{}

type imageFile struct {
	bufmodule.FileInfo

	fileDescriptorProto *descriptorpb.FileDescriptorProto
}

func newImageFile(
	fileDescriptorProto *descriptorpb.FileDescriptorProto,
	moduleCommit bufmodule.ModuleCommit,
	externalPath string,
	isImport bool,
) (*imageFile, error) {
	if err := protodescriptor.ValidateFileDescriptorProto(fileDescriptorProto); err != nil {
		return nil, err
	}
	coreFileInfo, err := bufcore.NewFileInfo(
		fileDescriptorProto.GetName(),
		externalPath,
		isImport,
	)
	if err != nil {
		return nil, err
	}
	return &imageFile{
		FileInfo:            bufmodule.NewFileInfo(coreFileInfo, moduleCommit),
		fileDescriptorProto: fileDescriptorProto,
	}, nil
}

func (f *imageFile) Proto() *descriptorpb.FileDescriptorProto {
	return f.fileDescriptorProto
}

func (f *imageFile) ImportPaths() []string {
	return f.fileDescriptorProto.GetDependency()
}

func (f *imageFile) withIsImport(isImport bool) ImageFile {
	return &imageFile{
		FileInfo:            bufmodule.NewFileInfo(f.FileInfo.WithIsImport(isImport), f.FileInfo.ModuleCommit()),
		fileDescriptorProto: f.fileDescriptorProto,
	}
}

func (*imageFile) isImageFile() {}
