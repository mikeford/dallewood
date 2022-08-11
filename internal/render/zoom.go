package render

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func InfiniteZoom(framesDir string, initialCrop float64, duration uint, easingMode AnimationEasingMode) error {
	frames, err := os.ReadDir(framesDir)
	if err != nil {
		return fmt.Errorf("failed to open frames directory: %v", err)
	}

	segmentsDir := "segments" // temporary artifacts directory
	err = os.RemoveAll(segmentsDir)
	if err != nil {
		return fmt.Errorf("failed to clear out segments directory: %v", err)
	}

	err = os.MkdirAll(segmentsDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create segments directory: %v", err)
	}

	rendersDir := "renders"
	if _, err := os.Stat(rendersDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(rendersDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create renders directory: %v", err)
		}
	}

	// Calculate total number of segments to be produced
	// so we can perform progress-based animation easing.
	numSegments := 0
	for _, file := range frames {
		if filepath.Ext(file.Name()) == ".png" {
			numSegments++
		}
	}

	// Calculate keyframes used to interpolate animation easing durations for each segment
	keyframes := []int{}
	for i := 0; i < numSegments; i++ {
		p := float64(i+1) / float64(numSegments)
		keyframes = append(keyframes, int(float64(duration)*calculateNormalizedTime(p, easingMode)))
	}

	currentSegment := 0 // Progress counter for easing
	segmentFilenames := []string{}
	// Produce each segment in order (TODO: eventually process segments concurrently)
	for _, file := range frames {
		if filepath.Ext(file.Name()) != ".png" {
			continue
		}

		segmentFile := fmt.Sprintf("%s/%03d.mp4", segmentsDir, currentSegment)
		segmentFilenames = append(segmentFilenames, fmt.Sprintf("file %03d.mp4", currentSegment))
		err = ffmpeg.Input(fmt.Sprintf("%s/%s", framesDir, file.Name())).
			Filter("scale", ffmpeg.Args{"-2:10*ih"}).
			ZoomPan(ffmpeg.KwArgs{
				"fps": "25",
				"x":   "iw/2-(iw/zoom/2)",
				"y":   "ih/2-(ih/zoom/2)",
				"z":   calculateZoomExpression(initialCrop),
				"d":   calculateDurationExpression(keyframes, currentSegment, numSegments),
				"s":   "1024x1024"}).
			Output(segmentFile, ffmpeg.KwArgs{
				"pix_fmt": "yuv420p",
				"c:v":     "libx264"}).
			OverWriteOutput().ErrorToStdOut().Run()
		if err != nil {
			return fmt.Errorf("failed to render segment for %s: %v", file.Name(), err)
		}

		currentSegment++
	}

	segmentListFilename := fmt.Sprintf("%s/segments.txt", segmentsDir)
	segmentListFile, err := os.OpenFile(segmentListFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to create segment list file: %s", err)
	}

	writer := bufio.NewWriter(segmentListFile)
	for _, data := range segmentFilenames {
		_, err = writer.WriteString(data + "\n")
		if err != nil {
			return fmt.Errorf("failed to write contents of segment file list: %v", err)
		}
	}
	writer.Flush()
	segmentListFile.Close()

	// We manually invoke ffmpeg because of a bug in which ffmpeg-go's concat lexically sorts stream filenames,
	//    causing the concatenation of > 9 streams to output out of order. (i.e. 1, 10 instead of 1, 2)
	//    See https://github.com/u2takey/ffmpeg-go/issues/58
	cmd := exec.Command("ffmpeg", "-f", "concat", "-i", segmentListFilename, "-c", "copy",
		fmt.Sprintf("renders/render_%d.mp4", time.Now().Unix()))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error concatening segments with ffmpeg: %v", err)
	}

	return nil
}

// Calculates the expression necessary to zoom out to a full view from an initial cropped view
func calculateZoomExpression(initialCrop float64) string {
	initialZoomFactor := 1 / initialCrop                                            // e.g. for a 30% initial crop, 1 / 0.3 = 3.333 zoom factor
	return fmt.Sprintf("%f-on/duration*%f", initialZoomFactor, initialZoomFactor-1) // 3.333-on/duration*(initialZoomFactor-1) where (on/duration) progresses from 0 to 1 for each frame
}

// Calculates the expression representing the total number of frames for a given segment
func calculateDurationExpression(keyframes []int, segment int, numSegments int) string {
	var fps int = 25
	var segmentSeconds int
	if segment == 0 {
		segmentSeconds = keyframes[0]
	} else {
		segmentSeconds = keyframes[segment] - keyframes[segment-1]
	}
	return fmt.Sprintf("%d", uint(fps*segmentSeconds))
}
