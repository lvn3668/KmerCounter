BioRad Coding assignment 1

This code is written in GO Language and the logic is as below:

1. It reads in 2 arguments (directory) and kmercounter
2. In the directory specified, it reads in all Fasta files
3. It checks if the kmerr counter  < 0 or is > length of the Fasta sequence
4. It builds a hashmap of the kmers of the specified length in each FASTA file
5. It writes out the kmer and the count to inputfile.kmercounter.counts.out

To run the program:
at the command prompt type "go run kmercounter.go <path to directory> <kmer size>"