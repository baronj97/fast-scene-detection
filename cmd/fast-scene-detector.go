// fast-scene-detector is a CLI tool to extract i-frames and approximate scene changes
package main

import (
	"fmt"
	"flag"
	"os"
	
	"github.com/baronj97/fast-scene-detection/extraction"
)

func main() {
	// Check for atleast one command before continuing.
	if len(os.Args) == 1 {
		fmt.Println("No commands passed. Pass in \"help\" for usage.")
		os.Exit(1)
	} 

	// Main and Sub Command Flag Setup
	extractCommand := flag.NewFlagSet("extract", flag.ExitOnError)
	interval := extractCommand.Int("i", 10, "Interval: The number of frames to skip in between i-frames. If i == 10, then every 10th i-frame is used. (Optional)")
	path := extractCommand.String("p", "", "Path: Full file path to a video file. (Required)")

	visualizeCommand := flag.NewFlagSet("visualize", flag.ExitOnError)
	format := visualizeCommand.String("f", "", "Format: The file format to save the output image to. (Required)")
	
	predictCommand := flag.NewFlagSet("predict", flag.ExitOnError)
	file := predictCommand.String("f", "", "File: The file path to the output from running the visualize command. (Required)")

	helpCommand := flag.NewFlagSet("help", flag.ExitOnError)
	
	// This section parses the user input based on the main command. 
	// os.Arg[0] is the main command.
    // os.Arg[2:] is the sub command.
	switch os.Args[1] {
	case "extract":
		if len(os.Args) > 2 {
			extractCommand.Parse(os.Args[2:])
		} else {
			extractCommand.PrintDefaults()
			os.Exit(1)
		}
	case "visualize":
		visualizeCommand.Parse(os.Args[2:])
	case "predict":
		visualizeCommand.Parse(os.Args[2:])
	case "help":
		helpCommand.Parse(os.Args[2:])
	default:
		fmt.Println("No commands passed. Pass in help for usage.")
		os.Exit(1)
	}

	// This section runs the appropriate functions based on main command input. Ensures inputs are valid.
	switch os.Args[1] {
	case "extract":
		// Required Flag
		if *path == "" {
			extractCommand.PrintDefaults()
			os.Exit(1)
		}

		// Utilize the extraction package to extract the i-frames from a video file.
		e := extraction.FFMPEGExtractor{Path: "ffmpeg"}
		es := extraction.NewExtractionService(e)
		result := es.IFrames(*interval, *path, "../data/iframes/")

		if result == nil {
			fmt.Println("Extraction Complete")
		} else { 
			fmt.Println(result) // Error from Extraction Package.
		}
	case "visualize":
		fmt.Println("Work in Progress!")

		// Required Flag
		if *format == "" {
			visualizeCommand.PrintDefaults()
			os.Exit(1)
		}

		visualizeCommand.PrintDefaults()
		os.Exit(1)
	case "predict":
		fmt.Println("Work in Progress!")

		// Required Flag
		if *file == "" {
			predictCommand.PrintDefaults()
			os.Exit(1)
		}

		predictCommand.PrintDefaults()
		os.Exit(1)
	case "help":
		fmt.Println("Extract:\n" +
					"NAME:\n" +
					"    fast-scene-detector extract - extract the i-frames from the video file\n\n" +  
					"USAGE:\n" +
					"    fast-scene-detector extract [command options] /path/to/video/file\n\n" +
					"OPTIONS:\n" +
					"    --interval, i  - The number of frames to skip in between i-frames. If i == 10, then every 10th i-frame is used. (default: 10)")

		fmt.Println("\n\nVisualize:\n" +
					"NAME:\n" +
					"    fast-scene-detector visualize - visualize the average frame color of all selected i-frames\n\n" +  
					"USAGE:\n" +
					"    fast-scene-detector visualize [command options]\n\n" +
					"OPTIONS:\n" +
					"    --format, f  - The file format to save the output image to.")

		fmt.Println("\n\nPredict:\n" +
					"NAME:\n" +
					"    fast-scene-detector preduct - predict the scene intervals based on the average color\n\n" +  
					"USAGE:\n" +
					"    fast-scene-detector predict /path/to/image/file\n\n")
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}