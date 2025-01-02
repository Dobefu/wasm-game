package game

import "time"

func Update() {
	FRAME = FRAME + 1

	DELTA_TIME = float64(time.Now().UnixMilli()-_LAST_TIME.UnixMilli()) / 1000
	_LAST_TIME = time.Now()
}
