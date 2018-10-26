package pakLThre

import "pakF/pakLTwo"

func World() string {
	return "world"
}
func HelloWold() string {
	return pakLTwo.Hello() + World()
}
