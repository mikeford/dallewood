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

	"github.com/mikeford/dallewood/internal/render"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"
)

var easingMode render.AnimationEasingMode

// zoomCmd represents the zoom command
var zoomCmd = &cobra.Command{
	Use:   "zoom",
	Short: "Create an infinite zoom video using progressively outcropped images",
	Long:  `Create an infinite zoom video using progressively outcropped images`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("Incorrect invocation: Must specify directory containing input frames")
		}

		initialCrop, err := cmd.Flags().GetFloat64("crop")
		if err != nil {
			panic(fmt.Errorf("invalid crop argument provided: %v", err))
		}

		duration, err := cmd.Flags().GetUint("duration")
		if err != nil {
			panic(fmt.Errorf("invalid duration argument provided: %v", err))
		}

		if err := render.InfiniteZoom(args[0], initialCrop, uint(duration), easingMode); err != nil {
			panic(fmt.Sprintf("Unable to render infinite zoom: %v", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(zoomCmd)
	zoomCmd.Flags().Float64P("crop", "i", 0.7, "Percentage of each frame to initially crop to (recommended to match downsize scale)")
	zoomCmd.Flags().UintP("duration", "d", 10, "Duration, in seconds, of the video to be rendered")
	zoomCmd.Flags().VarP(
		enumflag.New(&easingMode, "easing", render.AnimationEasingModeIds, enumflag.EnumCaseInsensitive),
		"easing", "e",
		"Specifies the animation easing mode (a.k.a timing function)")
}
