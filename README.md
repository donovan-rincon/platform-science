# Shipments Assignment

This program takes an input json that has shipments and drivers process the assignment using the challeng algorithm


# Challenge Algorithm

- If the length of the shipment's destination street name is even, the base suitability score (SS) is the number of vowels in the driver’s name multiplied by 1.5.
- If the length of the shipment's destination street name is odd, the base SS is the number of consonants in the driver’s name multiplied by 1.
- If the length of the shipment's destination street name shares any common factors (besides 1) with the length of the driver’s name, the SS is increased by 50% above the base SS.

# Assumptions:
- This implementation doesn't handle the case where there are more drivers than shipments or vice versa, and assumes that 
- There's always a one-to-one correspondence.
- It doesn't handle ties between drivers with the same SS, and simply picks the first one that appears in the drivers slice. 
- Inputs are valid and in valid format

# Pre-requisites

If want to build:
- Download the zip of the code or clone this repository
- Have go installed. Instructions in this link https://go.dev/doc/install
- Recommended have Visual Studio Code with go extension installed (https://marketplace.visualstudio.com/items?itemName=golang.Go)
- The program is a module that is already initialized to allow build
- To build go to cmd folder and run 
    - For Macos/Linux 
    ```
    go build -o ShipmentsAssignment
    ```
    - For Windows
    ```
    go build -o ShipmentsAssignment.exe
    ```

# How to run

Included is the binary file of the program, the .exe file for windows and the binary file for unix.
The following command will show the parameters that supports and what to use
- For Macos/Linux
```
./ShipmentsAssigemnt --help
```
- For Windows
```
ShipmentsAssigemnt.exe --help
```


The following command for example takes the input file that contains shipments and drivers in it
- For Macos/Linux
```
./ShipmentsAssigemnt --inputFilePath=[absolute_file_path]
```
- For Windows
```
ShipmentsAssigemnt.exe --inputFilePath=[absolute_file_path]
```
