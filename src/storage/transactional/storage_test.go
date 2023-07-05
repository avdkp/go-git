package transactional

import (
	"testing"

	"github.com/avdkp/go-git/src/plumbing"
	"github.com/avdkp/go-git/src/plumbing/cache"
	"github.com/avdkp/go-git/src/plumbing/storer"
	"github.com/avdkp/go-git/src/storage"
	"github.com/avdkp/go-git/src/storage/filesystem"
	"github.com/avdkp/go-git/src/storage/memory"
	"github.com/avdkp/go-git/src/storage/test"
	"github.com/go-git/go-billy/v5/memfs"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type StorageSuite struct {
	test.BaseStorageSuite
	temporal func() storage.Storer
}

var _ = Suite(&StorageSuite{
	temporal: func() storage.Storer {
		return memory.NewStorage()
	},
})

var _ = Suite(&StorageSuite{
	temporal: func() storage.Storer {
		fs := memfs.New()
		return filesystem.NewStorage(fs, cache.NewObjectLRUDefault())
	},
})

func (s *StorageSuite) SetUpTest(c *C) {
	base := memory.NewStorage()
	temporal := s.temporal()

	s.BaseStorageSuite = test.NewBaseStorageSuite(NewStorage(base, temporal))
}

func (s *StorageSuite) TestCommit(c *C) {
	base := memory.NewStorage()
	temporal := s.temporal()
	st := NewStorage(base, temporal)

	commit := base.NewEncodedObject()
	commit.SetType(plumbing.CommitObject)

	_, err := st.SetEncodedObject(commit)
	c.Assert(err, IsNil)

	ref := plumbing.NewHashReference("refs/a", commit.Hash())
	c.Assert(st.SetReference(ref), IsNil)

	err = st.Commit()
	c.Assert(err, IsNil)

	ref, err = base.Reference(ref.Name())
	c.Assert(err, IsNil)
	c.Assert(ref.Hash(), Equals, commit.Hash())

	obj, err := base.EncodedObject(plumbing.AnyObject, commit.Hash())
	c.Assert(err, IsNil)
	c.Assert(obj.Hash(), Equals, commit.Hash())
}

func (s *StorageSuite) TestTransactionalPackfileWriter(c *C) {
	base := memory.NewStorage()
	temporal := s.temporal()
	st := NewStorage(base, temporal)

	_, tmpOK := temporal.(storer.PackfileWriter)
	_, ok := st.(storer.PackfileWriter)
	c.Assert(ok, Equals, tmpOK)
}
