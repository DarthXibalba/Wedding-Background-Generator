package cmd

import (
	"fmt"
	"image"
	"image/draw"
	"path/filepath"
	"strings"

	"github.com/DarthXibalba/Wedding-Background-Generator/internal/imgproc"
	"github.com/spf13/cobra"
)

// splitAppendCmd represents the splitter command
var splitAppendCmd = &cobra.Command{
	Use:   "split",
	Short: "Split an image into parts",
	Long:  `Split an image into multiple parts based on specified criteria.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		imgPath := args[0]
		img, err := imgproc.LoadImage(imgPath)
		if err != nil {
			return err
		}

		// Create a mirror image for the backside
		mirroredImg := imgproc.MirrorImage(img)

		// Calculate the bounds and the middle of the image
		origBounds := img.Bounds()
		origHeight := origBounds.Dy()
		origWidth := origBounds.Dx()
		origMiddle := origWidth / 2

		newHeight := origHeight
		newWidth := origWidth * 2
		if flag_outputHeight > 0 {
			newHeight = flag_outputHeight
		}
		if flag_outputWidth > 0 {
			newWidth = flag_outputWidth
		}

		// Create a new canvas with double the width to accommodate the RLRL pattern
		newRect := image.Rect(0, 0, newWidth, newHeight)
		newImg := image.NewRGBA(newRect)

		// Draw the 2nd image in RM(LR)L pattern
		draw.Draw(newImg, image.Rect(0, 0, origMiddle, newHeight), img, image.Point{X: origMiddle, Y: 0}, draw.Src)
		draw.Draw(newImg, image.Rect(origMiddle, 0, origWidth+origMiddle, newHeight), mirroredImg, image.Point{}, draw.Src)
		draw.Draw(newImg, image.Rect(origWidth+origMiddle, 0, newWidth, newHeight), img, image.Point{}, draw.Src)

		// Save images
		ext := filepath.Ext(imgPath)
		base := strings.TrimSuffix(imgPath, ext)
		newExt := ".png"
		suffix := fmt.Sprintf("_cropped_%dx%d", newWidth, newHeight)

		newImgPath := base + suffix + newExt
		err = imgproc.SaveImage(newImg, newImgPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(splitAppendCmd)
}
