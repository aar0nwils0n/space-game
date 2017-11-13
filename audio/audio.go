package audio

import (
	"honnef.co/go/js/dom"
)

//Store stores audio files
type Store struct {
	Files map[string]*File
}

//CreateStore creates file store
func CreateStore() Store {
	store := Store{}
	store.Files = make(map[string]*File)
	return store
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
	as.Files[key] = &file
}

//StartLoop plays an audio file and loops at the designated positions
func (f *File) StartLoop(loopStart float64, loopEnd float64) {
	f.Playing = true
	f.el.Play()
	f.timeout = dom.GetWindow().SetTimeout(func() {
		f.el.Set("currentTime", loopStart)
		f.loop(loopStart, loopEnd)
	}, int(loopEnd*1000))
}

//Play will play the file once
func (f *File) Play() {
	f.el.Play()
}

func (f *File) loop(start float64, end float64) {
	f.timeout = dom.GetWindow().SetTimeout(func() {
		f.el.Set("currentTime", start)
		f.loop(start, end)
	}, int(end-start)*1000)
}

//StopLoop stops the audo from playing and resets the timeout and currentTime
func (f *File) StopLoop() {
	f.Playing = false
	dom.GetWindow().ClearTimeout(f.timeout)
	f.el.Set("currentTime", 0)
	f.el.Pause()
}
