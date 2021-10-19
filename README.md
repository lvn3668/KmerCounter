# KmerCounter
Kmer counter is written in GO Lang v 1.16.5
To install GO on Windows, follow the instructions at https://golang.org/doc/install
4 GO implementation of N-mer counter in DNA sequences which tests for validity of input.
It reads in file name (of fasta file)
It reads the size length  (kmer length) for which counts are desired and writes out to file, counts of all overlapping kmers of size 1 through the specified input.
It checks if fasta file is empty amd whether kmer length is specified. 
