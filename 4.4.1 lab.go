//
//
//

package main

import "fmt"

type EggCarton struct {
  -brownEggs: int
  -whiteEggs:int
}

+EggCarton()
+EggCarton(brownEggs, whiteEggs)
+setBrownEggs(brownEggs)
+setWhiteEggs(whiteEggs)
+addEggs(eggs, type)
+getBrownEggs():int
+getWhiteEggs(): int
+getTotalEggs():int