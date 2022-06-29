/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package disk

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"hash/fnv"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gregjones/httpcache"
	"github.com/peterbourgon/diskv"
	"k8s.io/klog/v2"
)

type cacheRoundTripper struct {
	rt *httpcache.Transport
}

// newCacheRoundTripper creates a roundtripper that reads the ETag on
// response headers and send the If-None-Match header on subsequent
// corresponding requests.
func newCacheRoundTripper(cacheDir string, rt http.RoundTripper) http.RoundTripper {
	d := diskv.New(diskv.Options{
		PathPerm: os.FileMode(0750),
		FilePerm: os.FileMode(0660),
		BasePath: cacheDir,
		TempDir:  filepath.Join(cacheDir, ".diskv-temp"),
	})
	t := httpcache.NewTransport(&crcDiskCache{disk: d})
	t.Transport = rt

	return &cacheRoundTripper{rt: t}
}

func (rt *cacheRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt.rt.RoundTrip(req)
}

func (rt *cacheRoundTripper) CancelRequest(req *http.Request) {
	type canceler interface {
		CancelRequest(*http.Request)
	}
	if cr, ok := rt.rt.Transport.(canceler); ok {
		cr.CancelRequest(req)
	} else {
		klog.Errorf("CancelRequest not implemented by %T", rt.rt.Transport)
	}
}

func (rt *cacheRoundTripper) WrappedRoundTripper() http.RoundTripper { return rt.rt.Transport }

// A crcDiskCache is a cache backend for github.com/gregjones/httpcache. It is
// similar to httpcache's diskcache package, but uses checksums to ensure cache
// integrity rather than fsyncing each cache entry in order to avoid performance
// degradation on MacOS.
//
// See https://github.com/kubernetes/kubernetes/issues/110753 for more.
type crcDiskCache struct {
	disk *diskv.Diskv
}

// Get the requested key from the cache on disk. If Get encounters an error, or
// the returned value is not a CRC-32 checksum followed by bytes with a matching
// checksum it will return false to indicate a cache miss.
func (c *crcDiskCache) Get(key string) ([]byte, bool) {
	b, err := c.disk.Read(sanitize(key))
	if err != nil || len(b) < binary.MaxVarintLen32 {
		return []byte{}, false
	}

	response := b[binary.MaxVarintLen32:]
	sum, _ := binary.Uvarint(b[:binary.MaxVarintLen32])
	if crc32.ChecksumIEEE(response) != uint32(sum) {
		return []byte{}, false
	}

	return response, true
}

// Set writes the response to a file on disk. The filename will be the FNV-32a
// hash of the key. The file will contain the CRC-32 checksum of the response
// bytes, followed by said response bytes.
func (c *crcDiskCache) Set(key string, response []byte) {
	sum := make([]byte, binary.MaxVarintLen32)
	_ = binary.PutUvarint(sum, uint64(crc32.ChecksumIEEE(response)))
	_ = c.disk.Write(sanitize(key), append(sum, response...)) // Nothing we can do with this error.
}

func (c *crcDiskCache) Delete(key string) {
	_ = c.disk.Erase(sanitize(key)) // Nothing we can do with this error.
}

// Sanitize an httpcache key such that it can be used as a diskv key, which must
// be a valid filename. The httpcache key will either be the requested URL (if
// the request method was GET) or "<method> <url>" for other methods, per the
// httpcache.cacheKey function.
func sanitize(key string) string {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key)) // Writing to a hash never returns an error.
	return fmt.Sprintf("%X", h.Sum32())
}
