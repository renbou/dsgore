package slayer

import (
	"os"
	"path/filepath"
	"sync"
)

const PLAGUE_NAME = ".DS_Store"
const MAXWORKERS = 100

// The slayer handles concurrent processing of directories
type slayer struct {
	jobs    chan string
	errChan chan error
	wg      sync.WaitGroup
}

func (s *slayer) writeError(path string, err error) {
	if err != nil && s.errChan != nil {
		s.errChan <- err
	}
}

func (s *slayer) readDir(path string) ([]os.FileInfo, error) {
	if file, err := os.Open(path); err != nil {
		return nil, err
	} else if files, err := file.Readdir(0); err != nil {
		_ = file.Close()
		return nil, err
	} else {
		_ = file.Close()
		return files, nil
	}
}

func (s *slayer) handleFile(path string) {
	// Eradicate the abomination
	if filepath.Base(path) == PLAGUE_NAME {
		s.writeError(path, os.Remove(path))
	}
}

func (s *slayer) process(path string) {
	// Mark one job as done
	defer s.wg.Done()

	files, err := s.readDir(path)
	// If ye can't handle the directory, then just quit yer job lol
	if err != nil {
		s.writeError(path, err)
		return
	}

	// Time to bring out da real gunz
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

func (s *slayer) newJob(path string) {
	// So that we wait until the job is processed completely
	s.wg.Add(1)

	select {
	case s.jobs <- path: // Added new job to the queue
	default: // Process ourselves
		s.process(path)
	}
}

func (s *slayer) worker() {
	for file := range s.jobs {
		s.process(file)
	}
}

// Concurrently remove all .DS_Store files from directory pointed to by directory
// If errChan is not nil then errors are written to it
// Thus if you want to call this synchronously, then make errChan nil
// Otherwise call this as a goroutine and read all errors from errChan to wait for completion
func Slay(directory string, errChan chan error) {
	slayer := &slayer{
		jobs:    make(chan string, MAXWORKERS),
		errChan: errChan,
	}

	// Release the hunters
	for i := 0; i < MAXWORKERS; i++ {
		go slayer.worker()
	}

	// Begin the first job, which will quit once it processes the directory and leave only the workers
	slayer.newJob(directory)
	// Wait for the cleansing to be completed and close the channels,
	// notifying the caller that we're done and closing the workers
	slayer.wg.Wait()
	close(slayer.jobs)
	close(slayer.errChan)
}
