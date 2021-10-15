package slayer

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const PLAGUE_NAME = ".DS_Store"

// Don't initialize more workers than available threads since its useless here
var MAXWORKERS = runtime.GOMAXPROCS

// Describes an error. We add the path since it is not included by default.
type slayError struct {
	err  error
	path string
}

func (s slayError) Error() string {
	return s.path + ": " + s.err.Error()
}

// The slayer handles concurrent processing of directories
type slayer struct {
	directory string
	jobs      chan string
	errChan   chan error
	wg        sync.WaitGroup
}

func (s *slayer) path(p string) string {
	return filepath.Join(s.directory, p)
}

func (s *slayer) writeError(path string, err error) {
	if err != nil {
		s.errChan <- slayError{
			err:  err,
			path: s.path(path),
		}
	}
}

func (s *slayer) readDir(path string) []os.FileInfo {
	// If ye can't handle the directory, then just quit yer job lol
	if file, err := os.Open(s.path(path)); err != nil {
		s.writeError(path, err)
		return nil
	} else if files, err := file.Readdir(0); err != nil {
		s.writeError(path, err)
		return nil
	} else {
		return files
	}
}

func (s *slayer) handleFile(path string) {
	// Eradicate the abomination
	if filepath.Base(path) == PLAGUE_NAME {
		s.writeError(path, os.Remove(s.path(path)))
	}
}

func (s *slayer) process(path string) {
	// Mark one job as done
	defer s.wg.Done()

	files := s.readDir(path)

	for _, fileinfo := range files {
		fullpath := filepath.Join(path, fileinfo.Name())
		// Regular file encountered
		if fileinfo.Mode()&os.ModeType == 0 {
			s.handleFile(fullpath)
		} else if fileinfo.IsDir() {
			s.newJob(fullpath)
		}
	}
}

func (s *slayer) worker() {
	for file := range s.jobs {
		s.process(file)

	}
}

func (s *slayer) newJob(path string) {
	// So that we wait until the job is processed completely
	s.wg.Add(1)

	select {
	case s.jobs <- path: // Added new job to the queue
	default: // Process ourselves

	}
}

// Concurrently remove all .DS_Store files from directory pointed to by directory
// If errChan is not nil then errors are written to it
func Slay(directory string, errChan chan error) {

}
