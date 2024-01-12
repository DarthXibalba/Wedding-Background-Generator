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
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a wedding background based on the image specified",
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
		doubleWidth := origWidth * 2
		newWidth := doubleWidth
		if flag_outputHeight > 0 {
			newHeight = flag_outputHeight
		}
		if flag_outputWidth > 0 {
			newWidth = flag_outputWidth
		}

		// Calculate Offset
		xOffset := equalizedOffset(doubleWidth, newWidth)
		//yOffset := equalizedOffset(origHeight, newHeight)

		// Create a new canvas with double the width to accommodate the RLRL pattern
		newRect := image.Rect(0, 0, newWidth, origHeight)
		newImg := image.NewRGBA(newRect)

		// Draw the 2nd image in [R][M(LR)][L] pattern
		x1 := 0
		x2 := origMiddle + xOffset
		y1 := 0
		y2 := origHeight
		draw.Draw(newImg, image.Rect(x1, y1, x2, y2), img, image.Point{X: origMiddle - xOffset, Y: 0}, draw.Src)

		x1 = origMiddle + xOffset
		x2 = 0
		y1 = origWidth + origMiddle + xOffset
		y2 = origHeight
		draw.Draw(newImg, image.Rect(x1, x2, y1, y2), mirroredImg, image.Point{}, draw.Src)

		x1 = origWidth + origMiddle + xOffset
		x2 = 0
		y1 = newWidth
		y2 = origHeight
		draw.Draw(newImg, image.Rect(x1, x2, y1, y2), img, image.Point{}, draw.Src)

		// Scale the image to the specified newHeight

		// Save images
		ext := filepath.Ext(imgPath)
		base := strings.TrimSuffix(imgPath, ext)
		newExt := ".png"
		suffix := fmt.Sprintf("_generated_%dx%d", newWidth, origHeight)

		newImgPath := base + suffix + newExt
		err = imgproc.SaveImage(newImg, newImgPath)
		if err != nil {
			return err
		}

		return nil
	},
}

//func max(a int, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func equalizedOffset(old int, new int) int {
	return (new - old) / 2
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
