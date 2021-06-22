BioRad Coding assignment 1

This code is written in GO Language and the logic is as below:

1. It reads in 2 arguments (directory) and kmercounter
2. In the directory specified, it reads in all Fasta files
3. It checks if the kmerr counter less than 0 or is greater than length of the Fasta sequence
4. It builds a hashmap of the kmers of the specified length in each FASTA file
5. It writes out the kmer and the count to inputfile.kmercounter.counts.out

To run the program:
at the command prompt type "go tool compiler kmercounter.go ; go run kmercounter.go <path to directory> <kmer size>"

GO Lang can be installed on Windows using the instructions at https://golang.org/doc/install 
On successful installation, type "go version" or "go" at the command prompt to see the list of commands.
The GoLand IDE Version 2021.1.3 was used. It can be downloaded from jetbrains.com/go/
The installation requirements are at https://www.jetbrains.com/help/go/installation-guide.html
Under Project settings, the command line settings can be set as "path-to-directory kmer-size" and the project can be run by clicking on "Run go build kmercounter" (or whatever name the project is saved as). This automatically compiles the code, and runs it using the settings specified in the config file (using the IDE interface) and hence runs this tool with the stated command-line-arguments , be it path to the directory containing the fasta files or the kmer size for which counts are to be calculated in each fasta file.

