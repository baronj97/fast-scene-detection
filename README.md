# Fast Scene Detection

## Purpose

This command-line utility provides usage to extract i-frames from video files and process them to determine scene changes.

## User Features

1. Extract i-frames from video file
1. Compute the average frame color and generate a spectrum for visualization.
1. Determine time intervals for scenes and produce an output text file with scene times.

## Tool Usage

        NAME:
            fast-scene-detector - CLI tool to extract i-frames and approximate scene changes

        USAGE:
            fast-scene-detector [global options] command [command options] [arguments...]

        COMMANDS:
            extract, e      - Extract generates a directory of i-frames from the input file
            visualize, v    - Visualize creates a spectrum of colors by analyzing the average color of each i-frame
            predict, p      - Predict produces a text file with the number of scenes and their time intervals 
            help, h         - Help shows the usage of the tool and a short description of each command

        GLOBAL OPTIONS:
            --frames-dir value       - The directory to save and load the i-frames (default: "./frames")
            --spectrum-path value    - The path to save the spectrum file (default: "./spectrum.png")
            --scenes-path value      - The path to save the scene listing file (default: "./scenes.txt")

## Command Usage

1. Extract

        NAME:
            fast-scene-detector extract - extract the i-frames from the video file

        USAGE:
            fast-scene-detector extract [command options] /path/to/video/file

        OPTIONS:
            --interval, i  - The number of frames to skip in between i-frames. If i == 10, then every 10th i-frame is used. (default: 10)

1. Visualize

        NAME:
            fast-scene-detector visualize - visualize the average frame color of all selected i-frames

        USAGE:
            fast-scene-detector visualize [command options]

        OPTIONS:
            --format, f  - The file format to save the output image to.

1. Predict

        NAME:
            fast-scene-detector preduct - predict the scene intervals based on the average color

        USAGE:
            fast-scene-detector visualize

## Examples

TODO
