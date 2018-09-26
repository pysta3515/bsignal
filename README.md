# bsignal
use bsignal to control other goroutine do something and do sth in its own routine after other goroutine all been done!

# usage
```
import "github.com/pysta/bsignal"

func main() {
	bSignal := bsignal.NewBSignal()
	go BTask(bSignal)
	go CTask(bSignal)
	bSignal.BroadCast() // notify BTask & CTask to do sth.
	bSignal.AllDone()   //BTask & CTask all done
	// do sth ...
}

func BTask(bSignal *bsignal.BSignal) {
	bSignal.Subscribe()
	for {
		if bSignal.Listened() {
			//do your job
			bSignal.Done()
		} else {
			// to other job
		}
	}
}

func CTask(bSignal *bsignal.BSignal) {
	bSignal.Subscribe()
	for {
		if bSignal.Listened() {
			//do your job
			bSignal.Done()
		} else {
			// to other job
		}
	}
}
```
