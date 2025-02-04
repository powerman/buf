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

package build

import (
	"context"
	"fmt"

	"github.com/powerman/buf/internal/buf/bufanalysis"
	"github.com/powerman/buf/internal/buf/bufcli"
	"github.com/powerman/buf/internal/buf/bufconfig"
	"github.com/powerman/buf/internal/buf/bufcore/bufimage"
	"github.com/powerman/buf/internal/buf/buffetch"
	"github.com/powerman/buf/internal/buf/bufwork"
	"github.com/powerman/buf/internal/pkg/app"
	"github.com/powerman/buf/internal/pkg/app/appcmd"
	"github.com/powerman/buf/internal/pkg/app/appflag"
	"github.com/powerman/buf/internal/pkg/storage/storageos"
	"github.com/powerman/buf/internal/pkg/stringutil"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	asFileDescriptorSetFlagName = "as-file-descriptor-set"
	errorFormatFlagName         = "error-format"
	excludeImportsFlagName      = "exclude-imports"
	excludeSourceInfoFlagName   = "exclude-source-info"
	pathsFlagName               = "path"
	outputFlagName              = "output"
	outputFlagShortName         = "o"
	configFlagName              = "config"

	// deprecated
	sourceFlagName = "source"
	// deprecated
	sourceConfigFlagName = "source-config"
	// deprecated
	filesFlagName = "file"
)

// NewCommand returns a new Command.
func NewCommand(
	name string,
	builder appflag.Builder,
	moduleResolverReaderProvider bufcli.ModuleResolverReaderProvider,
	deprecated string,
	hidden bool,
) *appcmd.Command {
	flags := newFlags()
	return &appcmd.Command{
		Use:        name + " <input>",
		Short:      "Build all files from the input location and output an image.",
		Long:       bufcli.GetInputLong(`the source or module to build, or image to convert`),
		Args:       cobra.MaximumNArgs(1),
		Deprecated: deprecated,
		Hidden:     hidden,
		Run: builder.NewRunFunc(
			func(ctx context.Context, container appflag.Container) error {
				return run(ctx, container, flags, moduleResolverReaderProvider)
			},
			bufcli.NewErrorInterceptor(),
		),
		BindFlags: flags.Bind,
	}
}

type flags struct {
	AsFileDescriptorSet bool
	ErrorFormat         string
	ExcludeImports      bool
	ExcludeSourceInfo   bool
	Paths               []string
	Output              string
	Config              string

	// deprecated
	Source string
	// deprecated
	SourceConfig string
	// deprecated
	Files []string
	// special
	InputHashtag string
}

func newFlags() *flags {
	return &flags{}
}

func (f *flags) Bind(flagSet *pflag.FlagSet) {
	bufcli.BindInputHashtag(flagSet, &f.InputHashtag)
	bufcli.BindAsFileDescriptorSet(flagSet, &f.AsFileDescriptorSet, asFileDescriptorSetFlagName)
	bufcli.BindExcludeImports(flagSet, &f.ExcludeImports, excludeImportsFlagName)
	bufcli.BindExcludeSourceInfo(flagSet, &f.ExcludeSourceInfo, excludeSourceInfoFlagName)
	bufcli.BindPathsAndDeprecatedFiles(flagSet, &f.Paths, pathsFlagName, &f.Files, filesFlagName)
	flagSet.StringVar(
		&f.ErrorFormat,
		errorFormatFlagName,
		"text",
		fmt.Sprintf(
			"The format for build errors, printed to stderr. Must be one of %s.",
			stringutil.SliceToString(bufanalysis.AllFormatStrings),
		),
	)
	flagSet.StringVarP(
		&f.Output,
		outputFlagName,
		outputFlagShortName,
		app.DevNullFilePath,
		fmt.Sprintf(
			`The location to write the image to. Must be one of format %s.`,
			buffetch.ImageFormatsString,
		),
	)
	flagSet.StringVar(
		&f.Config,
		configFlagName,
		"",
		`The config file or data to use.`,
	)

	// deprecated
	flagSet.StringVar(
		&f.Source,
		sourceFlagName,
		"",
		fmt.Sprintf(
			`The source or module to build, or image to convert. Must be one of format %s.`,
			buffetch.AllFormatsString,
		),
	)
	_ = flagSet.MarkDeprecated(
		sourceFlagName,
		`input as the first argument instead.`+bufcli.FlagDeprecationMessageSuffix,
	)
	_ = flagSet.MarkHidden(sourceFlagName)
	// deprecated
	flagSet.StringVar(
		&f.SourceConfig,
		sourceConfigFlagName,
		"",
		`The config file or data to use.`,
	)
	_ = flagSet.MarkDeprecated(
		sourceConfigFlagName,
		fmt.Sprintf("use --%s instead.%s", configFlagName, bufcli.FlagDeprecationMessageSuffix),
	)
	_ = flagSet.MarkHidden(sourceConfigFlagName)
}

func run(
	ctx context.Context,
	container appflag.Container,
	flags *flags,
	moduleResolverReaderProvider bufcli.ModuleResolverReaderProvider,
) error {
	if flags.Output == "" {
		return bufcli.NewFlagIsRequiredError(outputFlagName)
	}
	input, err := bufcli.GetInputValue(container, flags.InputHashtag, flags.Source, sourceFlagName, ".")
	if err != nil {
		return err
	}
	inputConfig, err := bufcli.GetStringFlagOrDeprecatedFlag(
		flags.Config,
		configFlagName,
		flags.SourceConfig,
		sourceConfigFlagName,
	)
	if err != nil {
		return err
	}
	paths, err := bufcli.GetStringSliceFlagOrDeprecatedFlag(
		flags.Paths,
		pathsFlagName,
		flags.Files,
		filesFlagName,
	)
	if err != nil {
		return err
	}
	ref, err := buffetch.NewRefParser(container.Logger()).GetRef(ctx, input)
	if err != nil {
		return err
	}
	configProvider := bufconfig.NewProvider(container.Logger())
	workspaceConfigProvider := bufwork.NewProvider(container.Logger())
	moduleResolver, err := moduleResolverReaderProvider.GetModuleResolver(ctx, container)
	if err != nil {
		return err
	}
	moduleReader, err := moduleResolverReaderProvider.GetModuleReader(ctx, container)
	if err != nil {
		return err
	}
	storageosProvider := storageos.NewProvider(storageos.ProviderWithSymlinks())
	imageConfigs, fileAnnotations, err := bufcli.NewWireImageConfigReader(
		container.Logger(),
		storageosProvider,
		configProvider,
		workspaceConfigProvider,
		moduleResolver,
		moduleReader,
	).GetImageConfigs(
		ctx,
		container,
		ref,
		inputConfig,
		paths,
		false,
		flags.ExcludeSourceInfo,
	)
	if err != nil {
		return err
	}
	if len(fileAnnotations) > 0 {
		// stderr since we do output to stdout potentially
		if err := bufanalysis.PrintFileAnnotations(
			container.Stderr(),
			fileAnnotations,
			flags.ErrorFormat,
		); err != nil {
			return err
		}
		return bufcli.ErrFileAnnotation
	}
	imageRef, err := buffetch.NewImageRefParser(container.Logger()).GetImageRef(ctx, flags.Output)
	if err != nil {
		return fmt.Errorf("--%s: %v", outputFlagName, err)
	}
	images := make([]bufimage.Image, 0, len(imageConfigs))
	for _, imageConfig := range imageConfigs {
		images = append(images, imageConfig.Image())
	}
	image, err := bufimage.MergeImages(images...)
	if err != nil {
		return err
	}
	return bufcli.NewWireImageWriter(
		container.Logger(),
	).PutImage(
		ctx,
		container,
		imageRef,
		image,
		flags.AsFileDescriptorSet,
		flags.ExcludeImports,
	)
}
