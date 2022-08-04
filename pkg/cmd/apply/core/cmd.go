// Copyright Contributors to the Open Cluster Management project
package core

import (
	"fmt"

	"github.com/stolostron/applier/pkg/cmd/apply/common"
	genericclioptionsapplier "github.com/stolostron/applier/pkg/genericclioptions"
	"github.com/stolostron/applier/pkg/helpers"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var example = `
# Apply core-resources templates
%[1]s apply core-resources --values values.yaml --path template_path1 tempalte_path2...
`

// NewCmd ...
func NewCmd(applierFlags *genericclioptionsapplier.ApplierFlags, streams genericclioptions.IOStreams) *cobra.Command {
	o := common.NewOptions(applierFlags, streams)

	cmd := &cobra.Command{
		Use:          "core-resources",
		Short:        "apply core-resources templates located in paths",
		Long:         "apply core-resources templates located in paths with a values.yaml, the list of path can be a path to a file or a directory",
		Example:      fmt.Sprintf(example, helpers.GetExampleHeader()),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			o.ResourcesType = common.CoreResources
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&o.ValuesPath, "values", "", "The files containing the values")
	cmd.Flags().StringArrayVar(&o.Paths, "paths", []string{}, "The list of template paths")
	cmd.Flags().StringVar(&o.OutputFile, "output-file", "", "The generated resources will be copied in the specified file")
	cmd.Flags().BoolVar(&o.SortOnKind, "sort-on-kind", false, "If set the files will be sorted by their kind")
	return cmd
}
