/*
Copyright © 2022 Mike Ford mikeford@users.noreply.github.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	"github.com/mikeford/dallewood/internal/downsize"
	"github.com/spf13/cobra"
)

// downsizeCmd represents the downsize command
var downsizeCmd = &cobra.Command{
	Use:   "downsize",
	Short: "Downsize image for infinite zoom outcropping",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("Incorrect invocation: Must specify image file to downsize")
		}
		scale, err := cmd.Flags().GetFloat64("scale")
		if err != nil {
			panic(fmt.Errorf("invalid scale argument provided: %v", err))
		}
		transition, err := cmd.Flags().GetBool("transition")
		if err != nil {
			panic(fmt.Errorf("invalid transition argument provided: %v", err))
		}
		transitionScale, err := cmd.Flags().GetFloat64("transition-scale")
		if err != nil {
			panic(fmt.Errorf("invalid transition scale argument provided: %v", err))
		}
		convertSource, err := cmd.Flags().GetBool("convert-source")
		if err != nil {
			panic(fmt.Errorf("invalid convert source argument provided: %v", err))
		}
		if err := downsize.ImageForInfiniteZoom(args[0], scale, transition, transitionScale, convertSource); err != nil {
			panic(fmt.Errorf("failed to downsize image: %v", err))
		}
	},
}

func init() {
	zoomCmd.AddCommand(downsizeCmd)
	downsizeCmd.Flags().Float64P("scale", "s", 0.7, "Percentage of the new image the original image should occupy")
	downsizeCmd.Flags().Bool("transition", true, "Whether to include edge transitions")
	downsizeCmd.Flags().Float64("transition-scale", 0.03, "Percentage of the image width to use as transition size")
	downsizeCmd.Flags().Bool("convert-source", true, "Whether the original image should be converted in place to PNG (Warning: this operation will delete the original non-PNG input)")
}
