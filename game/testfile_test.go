package game

import "github.com/haronius/space-ship/audio"

type testFile struct {
}

func (t *testFile) StartLoop(float64, float64) {}
func (t *testFile) Pause()                     {}
func (t *testFile) Play()                      {}
func (t *testFile) Loop(float64, float64)      {}
func (t *testFile) LoopFull()                  {}
func (t *testFile) StopLoop()                  {}

func createTestAudioStore() audio.Store {
	store := audio.CreateStore()
	file := testFile{}
	var iFile audio.IFile
	iFile = &file
	store.Files["thruster"] = iFile
	store.Files["explosion"] = iFile
	return store
}
