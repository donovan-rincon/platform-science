package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/platform-science/internal"
)

func main() {

	// Define flags
	inputFilePath := flag.String("inputFilePath", "", "File path for the file that contains the drivers and shipments in json format ")
	driversFilePath := flag.String("driversFilePath", "", "File path for the file that contains only the drivers in json format (requires to also include --shipmentsFilePath)")
	shipmentsFilePath := flag.String("shipmentsFilePath", "", "File path for the file that contains only the shipments in json format (requires to also include --driversFilePath)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s \n", os.Args[0])

		fmt.Fprintln(os.Stderr, "\nThis program takes an input of drivers and shipments and assign the best pair based in the algorithm:")
		fmt.Fprintln(os.Stderr, "- If the length of the shipment's destination street name is even, the base suitability score (SS) is the number of vowels in the driver’s name multiplied by 1.5.")
		fmt.Fprintln(os.Stderr, "- If the length of the shipment's destination street name is odd, the base SS is the number of consonants in the driver’s name multiplied by 1.")
		fmt.Fprintln(os.Stderr, "- If the length of the shipment's destination street name shares any common factors (besides 1) with the length of the driver’s name, the SS is increased by 50% above the base SS`)")

		fmt.Fprintln(os.Stderr, "\nParameters:")
		flag.PrintDefaults()
	}

	flag.Parse()
	// If the user specified the "--help" option, display the usage message and exit
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if *inputFilePath != "" {
		internal.ReadInputFile(*inputFilePath)
	} else if *driversFilePath != "" && *shipmentsFilePath != "" {
		internal.ReadInputFiles(*shipmentsFilePath, *driversFilePath)
	} else {
		panic("Incorrect flags provided. Please provide either inputFilePath (drivers and shipments) or driversFilePath and shipmentsFilePath")
	}

	assignments := internal.ProcessAssignments()

	totalSS := 0.0
	for _, a := range assignments {
		totalSS += a.SS
		fmt.Printf("Shipment: %s -> Driver: %s (SS: %.2f)\n", a.Shipment, a.Driver, a.SS)
	}
	fmt.Printf("Total SS: %f\n", totalSS)
}
