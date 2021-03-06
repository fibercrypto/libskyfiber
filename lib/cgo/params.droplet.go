package main

import (
	params "github.com/skycoin/skycoin/src/params"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "skytypes.h"
*/
import "C"

//export SKY_params_DropletPrecisionToDivisor
func SKY_params_DropletPrecisionToDivisor(precision uint8) uint64 {
	return params.DropletPrecisionToDivisor(precision)
}

//export SKY_params_DropletPrecisionCheck
func SKY_params_DropletPrecisionCheck(precision uint8, amount uint64) uint32 {
	return libErrorCode(params.DropletPrecisionCheck(precision, amount))
}
