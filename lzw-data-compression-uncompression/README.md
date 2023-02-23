# LZW Data Compression and Uncompression

LZW is a data compression method that takes advantage of this repetition.
This compression method:

1. start with an initial model
2. read data piece by piece
3. update the model and encode the data as you go along

LZW is a "dictionary"-based compression algorithm.
LZW encodes data by referencing a dictionary.
To encode a sub-string, only a single code number, corresponding to that sub-string's index in the dictionary, needs to be written to the output file.
