package audio

import (
	"honnef.co/go/js/dom"
)

//Store stores audio files
type Store struct {
	Files map[string]IFile
}

//CreateStore creates file store
func CreateStore() Store {
	store := Store{}
	store.Files = make(map[string]IFile)
	return store
}

//IFile is an interface for interacting with an audiofile
type IFile interface {
	StartLoop(float64, float64)
	Pause()
	Play()
	Loop(float64, float64)
	LoopFull()
	StopLoop()
}

//File which can be played
type File struct {
	el      *dom.HTMLAudioElement
	timeout int
	Playing bool
}

//Add and preload audio file
func (as *Store) Add(key string, url string) {
	audio := dom.GetWindow().Document().CreateElement("audio").(*dom.HTMLAudioElement)
	audio.SetAttribute("src", url)
	file := File{el: audio}
	var iFile IFile
	iFile = &file
	as.Files[key] = iFile
}

//StartLoop plays an audio file and loops at the designated positions
func (f *File) StartLoop(loopStart float64, loopEnd float64) {
	if f.Playing {
		return
	}

	f.Playing = true
	f.el.Play()
	f.timeout = dom.GetWindow().SetTimeout(func() {
		f.el.Set("currentTime", loopStart)
		f.Loop(loopStart, loopEnd)
	}, int(loopEnd*1000))
}

//Pause the file
func (f *File) Pause() {
	f.el.Pause()
}

//Play will play the file once
func (f *File) Play() {
	f.el.Play()
}

//Loop starts a loop which recursively repeats itself
func (f *File) Loop(start float64, end float64) {
	f.timeout = dom.GetWindow().SetTimeout(func() {
		f.el.Set("currentTime", start)
		f.Loop(start, end)
	}, int(end-start)*1000)
}

//LoopFull loops the full length of the audio
func (f *File) LoopFull() {
	f.el.SetAttribute("loop", "true")
	f.el.Play()
}

//StopLoop stops the audo from playing and resets the timeout and currentTime
func (f *File) StopLoop() {
	if !f.Playing {
		return
	}

	f.Playing = false
	dom.GetWindow().ClearTimeout(f.timeout)
	f.el.Set("currentTime", 0)
	f.el.Pause()
}
