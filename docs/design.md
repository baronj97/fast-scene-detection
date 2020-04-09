# Design

This page will be used to document how the architectural aspects of the project 
will come together. This includes the use of existing APIs and packages as well 
as any new functionality created as needed.


The objective of this tool is two-fold. First, it is to provide a visual representation of the average color within i-frames. Second, it should analyze
the data and determine where scene splits occur.

The tool will consist of three commands

1. Extract

1. Visualize

1. Predict

## Inputs and Outputs

### Extract

#### Inputs

1. Original Video File

1. The ith number of i-frames to extract, default = 10

1. The output directory to store these files.

#### Outputs

1. A directory of i-frames saved as jpegs or pngs.

### Visualize

#### Inputs

1. A directory of i-frames represented as jpegs or pngs.

#### Outputs

1. Some representation of the frames stored into data with their average frame color.

1. An image that visualizes the average frame color.

### Predict

#### Inputs

1. Some representation of the frames stored into data with their average frame color.

#### Outputs

1. A textfile with approximate time break downs based on the average frame colors. Use contrast to determine scene changes.

## Packages

To support the inputs and outputs stated above, let's implement packages with the following ideas.

### extraction

This package should deal with accepting a video file, extracting the i-frames, and saving them as output.

It might make sense to declare an interface which details the contract methods necessary for this and then provide an exported
struct type which implements the interface. This would allow for hiding the implementation details within the package.

To extract i-frames, I think the best option is to simply fork a process running ffmpeg. Consider the following link.
https://www.bogotobogo.com/FFMpeg/ffmpeg_thumbnails_select_scene_iframe.php?fbclid=IwAR1NicxcCAj0z5vAsun3kSLmzciDv-D3RHZusfKe76ew4Rww8K5j942Yae8

### frames

This package should deal with processing the frames and representing them into some data structure. This will probably be the most complicated package.

The first thing this package needs to do is to read the frames and compute the average color of the frame. This can be done using the GoCV package with
the imread function. This represents the image as a 3D "matrix", where each it is row x col x color channel. We can compute the average value of each color channel independently and save this into some struct, along with other metadata about the frame. Note, this assumes standard range pixel values; that is, 0-255, which is the canonical 8 bit representation of colors, 1 for each of the 3 RGB color channels. This provides a wide range of standard colors. This does not support HDR color mapping and probably isn't worth it anyway.

### scenes

Name to be confirmed. This package should handle the business of determining scene changes. It might honestly make sense to have this as part of the frames package -- not sure yet. What's difficult is that we need to do our best to limit the overlap between packages. My concern is that this package would rely too heavily on the "frames" data structure. In any case, we need something that will take a look at the frames and determine when scene changes occur, then try to approximate time values. If we have an MPEG4 file, we know that the time in between i-frames is 2.0002 seconds. Similarily, it is 0.5000 seconds for MPEG2. If we know the file format, then we can determine how long scenes occur based on the number of i-frames and how many were sampled. We should brainstorm more ideas for this once we get here. For scene changes, I was thinking some sort of mean moving average and when the average changes significantly in either direction, we assume a new scene.   