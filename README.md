# dallewood
A CLI to convert uncropped DALL·E 2 images into infinite zoom videos



https://user-images.githubusercontent.com/1922944/184796256-95d55338-d19b-4a97-b904-8f47af416d54.mp4



## Getting Started
### Dependencies
You must have the following installed:
- **Go** (such as via Go's official website or Homebrew)
- **ffmpeg** (must be available on your machine's PATH)

### Installing
Simply run `go install github.com/mikeford/dallewood` and verify you have the `dallewood` command available in your terminal.

## Usage
### Overview
`dallewood` v0.1 supports two main operations:

- `dallewood zoom [X]` creates an "infinite zoom" video from PNG images in the specified directory
- `dallewood zoom downsize [X]` downsizes a .png, .jpg, .jpeg, or .webp image and outputs the result to `X_downsized.png`

The `downsize` operation is used to prepare each frame for the next round of inpainting as the underlying images for an infinite zoom video are developed. Its operation is inspired by [this](https://editor.p5js.org/Pi_p5/full/qAqoieAhx) useful p5.js downsizing app. The idea was to create a CLI tool that can combine both the image preparation + video creation stages.

### Command Line Options
Both `zoom` and `zoom downsize` come with help menus accessible by running `zoom --help` and `zoom downsize --help` respectively.

In these, you'll notice each command comes with several options. For example, when calling `zoom downsize`, you can disable the stippled edge transitions around downsized image edges using the `--transition=false` flag. In addition, you might want to change their size using `--transition-scale`, a floating point value representing the percentage of the image the transition should occupy.

When it comes to `zoom`, the most important options are `--duration`, `crop`, and `easing`.

- `--duration` (`-d`) specifies the length, in seconds, of the video to be output
- `--crop` specifies the floating point percentage frame N occupies within frame N+1 for the entire sequence
- `--easing` specifies which animation timing function, otherwise known as *easing*, the video should follow. The choice of this setting dramatically affects the overall feel of your video.

For example, if you're working on a particular video and for each frame you generate, you have DALL·E 2 inpaint the outer 30% to achieve your outcrop, then `--crop` should be set to `0.7`. In other words, the central 70% of the frame is used by DALL·E 2 as the basis for each inpainting.

In fact, 70% cropping is the default used by both `zoom` and `zoom downsize`. If you decide to change the cropping from the default, make sure you also use `dallewood zoom downsize --scale=X` to specify the same decimal value you use with `zoom`.

### Directory Structure / File Naming Convention
As mentioned, `dallewood zoom downsize` is pretty flexible in that it will accept any `.png`, `.jpg`, `.jpeg`, or `.webp` file.

`dallewood zoom` on the other hand is a bit more picky: to ensure proper ordering of your images, keep two requirements in mind:

1. the directory you specify should only contain images you want included in the video (recommended to make a separate folder)
2. the PNG files in the directory should be ordered using *left padded* names. For example, if you have more than ten frames, files should follow the pattern `frame_001`, `frame_002`. Furthermore, if you have more than 100 frames, files should follow the format `frame_0979`, `frame_0980`, and so on.
