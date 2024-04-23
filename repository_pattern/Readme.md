# Repository Pattern

### Intuition:
Repository Pattern is to decouple the concerns of storing and accessing data.


## Resources:
### Repository Pattern:
https://youtu.be/ivJ2s0e7vi0?si=3TcossfW2ekJOOul
### Dependency Injection:
https://youtu.be/UX4XjxWcDB4?si=-W7CUdIfJrhF85eM   

Here is a detailed yet simple implementation of a dependency injection from the video tutorial:
```
package main

import "fmt"

// Different types of rocks: Ice, Sandy, Concrete
/* Therefore the placeSafeties will be heavily dependent on the type of rock being climbed
to decouple this dependency we make an interface of the placingSafties so that the climber
depends upon behaviour of the SafetyPlacer and not on its implementation*/

type SafetyPlacer interface {
	placeSafeties()
}

type RockClimber struct {
	rocksClimbed int
	sp           SafetyPlacer
}

func NewRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{sp: sp}
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

type IceSafetyPlacer struct {
}

type RockSafetyPlacer struct {
}

func (sp IceSafetyPlacer) placeSafeties() {
	fmt.Printf("Placing my Ice safties!\n")
}

func (sp RockSafetyPlacer) placeSafeties() {
	fmt.Printf("Placing my Rock safties!\n")
}

func main() {
	rc := NewRockClimber(IceSafetyPlacer{})
	for i := 0; i < 15; i++ {
		rc.climbRock()
	}
}
```
