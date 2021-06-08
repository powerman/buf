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

package organizationcreate

import (
	"context"
	"fmt"

	"github.com/powerman/buf/internal/buf/bufcli"
	"github.com/powerman/buf/internal/buf/bufcore/bufmodule"
	"github.com/powerman/buf/internal/buf/bufprint"
	"github.com/powerman/buf/internal/pkg/app/appcmd"
	"github.com/powerman/buf/internal/pkg/app/appflag"
	"github.com/powerman/buf/internal/pkg/rpc"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const formatFlagName = "format"

// NewCommand returns a new Command
func NewCommand(
	name string,
	builder appflag.Builder,
) *appcmd.Command {
	flags := newFlags()
	return &appcmd.Command{
		Use:   name + " <buf.build/organization>",
		Short: "Create a new organization.",
		Args:  cobra.ExactArgs(1),
		Run: builder.NewRunFunc(
			func(ctx context.Context, container appflag.Container) error {
				return run(ctx, container, flags)
			},
			bufcli.NewErrorInterceptor(),
		),
		BindFlags: flags.Bind,
	}
}

type flags struct {
	Format string
}

func newFlags() *flags {
	return &flags{}
}

func (f *flags) Bind(flagSet *pflag.FlagSet) {
	flagSet.StringVar(
		&f.Format,
		formatFlagName,
		bufprint.FormatText.String(),
		fmt.Sprintf(`The output format to use. Must be one of %s`, bufprint.AllFormatsString),
	)
}

func run(
	ctx context.Context,
	container appflag.Container,
	flags *flags,
) error {
	moduleOwner, err := bufmodule.ModuleOwnerForString(container.Arg(0))
	if err != nil {
		return appcmd.NewInvalidArgumentError(err.Error())
	}
	apiProvider, err := bufcli.NewRegistryProvider(ctx, container)
	if err != nil {
		return err
	}
	service, err := apiProvider.NewOrganizationService(ctx, moduleOwner.Remote())
	if err != nil {
		return err
	}
	organization, err := service.CreateOrganization(
		ctx,
		moduleOwner.Owner(),
	)
	if err != nil {
		if rpc.GetErrorCode(err) == rpc.ErrorCodeAlreadyExists {
			return bufcli.NewOrganizationNameAlreadyExistsError(container.Arg(0))
		}
		return err
	}
	return bufcli.PrintOrganizations(ctx, moduleOwner.Remote(), container.Stdout(), flags.Format, organization)
}
