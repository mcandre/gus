package gus

import (
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"

	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

// Version is semver.
const Version = "0.0.1"

const gitIndexBasename = ".git"

const gitConfigBasename = "config"
const gitConfigMode = 0644

const gitModulesBasename = ".gitmodules"
const gitModulesMode = 0644

const modulesBasename = "modules"

const gitCommand = "git"

const submoduleSubCommand = "submodule"

const addSubCommand = "add"
const branchFlag = "-b"

const removeSubCommand = "rm"
const recursiveFlag = "-r"
const cachedFlag = "--cached"

// Init ensures a local git repository and submodule index are provisioned
// with the given top-level git project path.
func Init(top string) error {
	fsTop := osfs.New(top)
	fsGitIndex := osfs.New(path.Join(top, gitIndexBasename))
	storage := filesystem.NewStorage(fsGitIndex, &cache.ObjectLRU{})

	repo, err := git.Init(storage, fsTop)

	if err == git.ErrRepositoryAlreadyExists {
		rp, err2 := git.PlainOpen(top)

		if err2 != nil {
			return err2
		}

		repo = rp
	} else if err != nil {
		return err
	}

	worktree, err := repo.Worktree()

	if err != nil {
		return err
	}

	submodules, err := worktree.Submodules()

	if err != nil {
		return err
	}

	return submodules.Init()
}

// AddSubmodule registers a new submodule from the given
// top-level git project path,
// submodule URL,
// target path (may be empty),
// and branch (may be empty).
func AddSubmodule(top string, url string, target string, branch string) error {
	//
	// Pending https://github.com/src-d/go-git/issues/597
	//

	args := []string{submoduleSubCommand, addSubCommand}

	if branch != "" {
		args = append(args, branchFlag, branch)
	}

	args = append(args, url)

	if target != "" {
		args = append(args, target)
	}

	cmd := exec.Command(gitCommand, args...)
	cmd.Dir = top
	return cmd.Run()
}

// RemoveSubmodule unregisters a submodule from the given
// top-level git project path,
// by submodule url.
func RemoveSubmodule(top string, url string) error {
	modules := config.NewModules()

	gitModulesPath := path.Join(top, gitModulesBasename)

	gitModulesData, err := ioutil.ReadFile(gitModulesPath)

	if err != nil {
		return err
	}

	if err2 := modules.Unmarshal(gitModulesData); err2 != nil {
		return err2
	}

	var name *string
	var pth *string

	for n, sub := range modules.Submodules {
		if sub.URL == url {
			name = &n
			pth = &sub.Path
		}
	}

	if name == nil {
		return fmt.Errorf("No submodule registered with URL: %v", url)
	}

	delete(modules.Submodules, *name)

	if len(modules.Submodules) == 0 {
		if err2 := os.Remove(gitModulesPath); err2 != nil {
			return err2
		}
	} else {
		gitModulesData, err = modules.Marshal()

		if err != nil {
			return err
		}

		if err2 := ioutil.WriteFile(gitModulesPath, gitModulesData, gitModulesMode); err2 != nil {
			return err2
		}
	}

	//
	// Pending https://github.com/src-d/go-git/issues/1287
	//

	cmd := exec.Command(gitCommand, addSubCommand, gitModulesBasename)
	cmd.Dir = top

	if err2 := cmd.Run(); err2 != nil {
		return nil
	}

	repo, err := git.PlainOpen(top)

	if err != nil {
		return err
	}

	c, err := repo.Config()

	if err != nil {
		return err
	}

	delete(c.Submodules, *name)

	gitConfigData, err := c.Marshal()

	if err != nil {
		return err
	}

	gitConfigPath := path.Join(top, gitIndexBasename, gitConfigBasename)

	if err2 := ioutil.WriteFile(gitConfigPath, gitConfigData, gitConfigMode); err2 != nil {
		return err2
	}

	//
	// Pending https://github.com/src-d/go-git/issues/1288
	//

	cmd = exec.Command(gitCommand, removeSubCommand, cachedFlag, recursiveFlag, *pth)
	cmd.Dir = top

	if err2 := cmd.Run(); err2 != nil {
		return err2
	}

	modulesPath := path.Join(top, gitIndexBasename, modulesBasename, *pth)

	if err2 := os.RemoveAll(modulesPath); err2 != nil {
		return err2
	}

	//
	// TODO:
	//
	// For each empty directory leaf on modulesPath,
	// Remove that directory...
	//

	if err2 := os.RemoveAll(*pth); err2 != nil {
		return err2
	}

	//
	// TODO:
	//
	// For each empty directory leaf on pth,
	// Remove that directory...
	//

	return nil
}

// GetSubmodules enumerates registered git submodules as URL, path pairs,
// based on the given top-level git project path.
func GetSubmodules(top string) (map[string]string, error) {
	submodules := make(map[string]string)

	repo, err := git.PlainOpen(top)

	if err != nil {
		return nil, err
	}

	worktree, err2 := repo.Worktree()

	if err2 != nil {
		return nil, err2
	}

	subs, err2 := worktree.Submodules()

	if err2 != nil {
		return nil, err2
	}

	for _, sub := range subs {
		c := sub.Config()
		submodules[c.URL] = c.Path
	}

	return submodules, nil
}
