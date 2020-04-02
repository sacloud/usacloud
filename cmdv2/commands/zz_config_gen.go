// Copyright 2017-2020 The Usacloud Authors
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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"errors"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// configCmd represents the command to manage SAKURA Cloud Config
func configCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "A manage command of APIKey settings",
		Long:  `A manage command of APIKey settings`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO not implements: call edit func as default
		},
	}
}

func configCurrentCmd() *cobra.Command {
	configCurrentParam := params.NewCurrentConfigParam()
	cmd := &cobra.Command{
		Use: "current",

		Short: "Current Config",
		Long:  `Current Config`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return configCurrentParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, configCurrentParam)
			if err != nil {
				return err
			}

			if configCurrentParam.GenerateSkeleton {
				return generateSkeleton(ctx, configCurrentParam)
			}

			return funcs.ConfigCurrent(ctx, configCurrentParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&configCurrentParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&configCurrentParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&configCurrentParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&configCurrentParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&configCurrentParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	return cmd
}

func configDeleteCmd() *cobra.Command {
	configDeleteParam := params.NewDeleteConfigParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete Config",
		Long:    `Delete Config`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return configDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, configDeleteParam)
			if err != nil {
				return err
			}

			if configDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, configDeleteParam)
			}

			// confirm
			if !configDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.ConfigDelete(ctx, configDeleteParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&configDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&configDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&configDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&configDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&configDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&configDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	return cmd
}

func configEditCmd() *cobra.Command {
	configEditParam := params.NewEditConfigParam()
	cmd := &cobra.Command{
		Use: "edit",

		Short: "Edit Config (default)",
		Long:  `Edit Config (default)`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return configEditParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, configEditParam)
			if err != nil {
				return err
			}

			if configEditParam.GenerateSkeleton {
				return generateSkeleton(ctx, configEditParam)
			}

			return funcs.ConfigEdit(ctx, configEditParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&configEditParam.Token, "token", "", "", "API Token of SakuraCloud")
	fs.StringVarP(&configEditParam.Secret, "secret", "", "", "API Secret of SakuraCloud")
	fs.StringVarP(&configEditParam.Zone, "zone", "", "", "Target zone of SakuraCloud")
	fs.StringVarP(&configEditParam.DefaultOutputType, "default-output-type", "", "", "Default output format type")
	fs.StringVarP(&configEditParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&configEditParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&configEditParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&configEditParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&configEditParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	return cmd
}

func configListCmd() *cobra.Command {
	configListParam := params.NewListConfigParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List Config",
		Long:    `List Config`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return configListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, configListParam)
			if err != nil {
				return err
			}

			if configListParam.GenerateSkeleton {
				return generateSkeleton(ctx, configListParam)
			}

			return funcs.ConfigList(ctx, configListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&configListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&configListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&configListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&configListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&configListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	return cmd
}

func configMigrateCmd() *cobra.Command {
	configMigrateParam := params.NewMigrateConfigParam()
	cmd := &cobra.Command{
		Use: "migrate",

		Short: "Migrate Config",
		Long:  `Migrate Config`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return configMigrateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, configMigrateParam)
			if err != nil {
				return err
			}

			if configMigrateParam.GenerateSkeleton {
				return generateSkeleton(ctx, configMigrateParam)
			}

			return funcs.ConfigMigrate(ctx, configMigrateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&configMigrateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&configMigrateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&configMigrateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&configMigrateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&configMigrateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	return cmd
}

func configShowCmd() *cobra.Command {
	configShowParam := params.NewShowConfigParam()
	cmd := &cobra.Command{
		Use: "show",

		Short: "Show Config",
		Long:  `Show Config`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return configShowParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, configShowParam)
			if err != nil {
				return err
			}

			if configShowParam.GenerateSkeleton {
				return generateSkeleton(ctx, configShowParam)
			}

			return funcs.ConfigShow(ctx, configShowParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&configShowParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&configShowParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&configShowParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&configShowParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&configShowParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	return cmd
}

func configUseCmd() *cobra.Command {
	configUseParam := params.NewUseConfigParam()
	cmd := &cobra.Command{
		Use: "use",

		Short: "Use Config",
		Long:  `Use Config`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return configUseParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, configUseParam)
			if err != nil {
				return err
			}

			if configUseParam.GenerateSkeleton {
				return generateSkeleton(ctx, configUseParam)
			}

			return funcs.ConfigUse(ctx, configUseParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&configUseParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&configUseParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&configUseParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&configUseParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&configUseParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	return cmd
}

func init() {
	parent := configCmd()
	parent.AddCommand(configCurrentCmd())
	parent.AddCommand(configDeleteCmd())
	parent.AddCommand(configEditCmd())
	parent.AddCommand(configListCmd())
	parent.AddCommand(configMigrateCmd())
	parent.AddCommand(configShowCmd())
	parent.AddCommand(configUseCmd())
	rootCmd.AddCommand(parent)
}
