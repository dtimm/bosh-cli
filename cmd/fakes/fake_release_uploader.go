// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/cmd"
)

type FakeReleaseUploader struct {
	UploadReleasesStub        func([]byte) ([]byte, error)
	uploadReleasesMutex       sync.RWMutex
	uploadReleasesArgsForCall []struct {
		arg1 []byte
	}
	uploadReleasesReturns struct {
		result1 []byte
		result2 error
	}
}

func (fake *FakeReleaseUploader) UploadReleases(arg1 []byte) ([]byte, error) {
	fake.uploadReleasesMutex.Lock()
	fake.uploadReleasesArgsForCall = append(fake.uploadReleasesArgsForCall, struct {
		arg1 []byte
	}{arg1})
	fake.uploadReleasesMutex.Unlock()
	if fake.UploadReleasesStub != nil {
		return fake.UploadReleasesStub(arg1)
	} else {
		return fake.uploadReleasesReturns.result1, fake.uploadReleasesReturns.result2
	}
}

func (fake *FakeReleaseUploader) UploadReleasesCallCount() int {
	fake.uploadReleasesMutex.RLock()
	defer fake.uploadReleasesMutex.RUnlock()
	return len(fake.uploadReleasesArgsForCall)
}

func (fake *FakeReleaseUploader) UploadReleasesArgsForCall(i int) []byte {
	fake.uploadReleasesMutex.RLock()
	defer fake.uploadReleasesMutex.RUnlock()
	return fake.uploadReleasesArgsForCall[i].arg1
}

func (fake *FakeReleaseUploader) UploadReleasesReturns(result1 []byte, result2 error) {
	fake.UploadReleasesStub = nil
	fake.uploadReleasesReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

var _ cmd.ReleaseUploader = new(FakeReleaseUploader)
