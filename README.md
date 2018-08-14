# xxtea

[![GoDoc](https://godoc.org/github.com/hillu/go-xxtea?status.svg)](https://godoc.org/github.com/hillu/go-xxtea) ![Travis CI](https://api.travis-ci.org/syndeck/xxtea.svg?branch=master)

This is a Go implementation of the "Corrected Block TEA" or "XXTEA"
block cipher algorithm described in [Correction to
xtea](http://www.movable-type.co.uk/scripts/xxtea.pdf).  
  
This implementation differs from the original one implemented by [Hilko Bengen](https://github.com/hillu)
by how the data and the key are handled: in this implementation we use Big Endian Signed Integers instead 
of unsigned int32s.  
  
We've corrected this implementation in order to decrypt data that is 
encrypted [with this method](https://github.com/francesco-ficarola/OpenBeaconLogger/blob/master/src/main/java/it/uniroma1/dis/wsngroup/utils/XXTEA.java)
(aka [Mathieu Carbou](https://github.com/mathieucarbou)'s / [Ovea](https://github.com/Ovea)'s implementation).


## License

BSD 2-clause, see LICENSE file in the source distribution.

## Original Author
- Hilko Bengen <bengen@hilluzination.de>

## Contributors
- Denys Vitali <denys@denv.it>