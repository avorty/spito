package vrctFs

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

const VirtualFsPathPrefix = "/tmp/spito-vrct/fs"
const VirtualFilePostfix = ".prototype.bson"

type VRCTFs struct {
	virtualFSPath string
	revertSteps   RevertSteps
}

func MoveFile(source string, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	err = sourceFile.Close()
	if err != nil {
		return err
	}

	if err = os.Remove(source); err != nil {
		return err
	}

	err = destinationFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewFsVRCT() (VRCTFs, error) {
	err := os.MkdirAll(VirtualFsPathPrefix, os.ModePerm)
	revertSteps, err := NewRevertSteps()
	if err != nil {
		return VRCTFs{}, nil
	}

	err = os.MkdirAll(VirtualFsPathPrefix, os.ModePerm)
	if err != nil {
		return VRCTFs{}, err
	}

	dir, err := os.MkdirTemp(VirtualFsPathPrefix, "")
	if err != nil {
		return VRCTFs{}, err
	}

	return VRCTFs{
		virtualFSPath: dir,
		revertSteps:   revertSteps,
	}, nil
}

func (v *VRCTFs) DeleteRuntimeTemp() error {
	if err := v.revertSteps.DeleteRuntimeTemp(); err != nil {
		return err
	}
	return os.RemoveAll(v.virtualFSPath)
}

// Apply
// first parameter takes array of identifiers
// returns revertNumber if serializeRevertSteps is true
func (v *VRCTFs) Apply(rulesHistory []Rule, serializeRevertSteps bool) (int, error) {
	mergeDir, err := os.MkdirTemp("/tmp", "spito-fs-vrct-merge")
	if err != nil {
		return 0, err
	}

	if err := mergePrototypes(v.virtualFSPath, mergeDir); err != nil {
		return 0, err
	}

	if err := v.mergeToRealFs(mergeDir); err != nil {
		return 0, err
	}

	var revertNum int
	if serializeRevertSteps {
		revertNum, err = v.revertSteps.Serialize(rulesHistory)
		if err != nil {
			return 0, err
		}
	}

	return revertNum, os.RemoveAll(mergeDir)
}

func (v *VRCTFs) Revert(fn func(rule Rule) error) error {
	return v.revertSteps.Apply(fn)
}

func (v *VRCTFs) mergeToRealFs(mergeDirPath string) error {
	splitMergePath := strings.Split(mergeDirPath, "/")[3:]
	destPath := strings.Join(splitMergePath, "/")
	if len(destPath) != 0 {
		destPath = "/" + destPath
	}

	entries, err := os.ReadDir(mergeDirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if destPath == "" {
			destPath = "/"
		}
		realFsEntryPath := filepath.Join(destPath, entry.Name())
		mergeDirEntryPath := filepath.Join(mergeDirPath, entry.Name())

		doesRealFsEntryExists, err := pathExists(realFsEntryPath)
		if err != nil {
			return err
		}

		if entry.IsDir() {
			// If originally dir does not exist, then revert should delete it
			if !doesRealFsEntryExists {
				v.revertSteps.RemoveDirAll(realFsEntryPath)
			}
			if err := os.MkdirAll(realFsEntryPath, os.ModePerm); err != nil {
				return err
			}
			if err := v.mergeToRealFs(mergeDirEntryPath); err != nil {
				return err
			}
			continue
		}

		filePrototype := FilePrototype{}
		err = filePrototype.Read(v.virtualFSPath, realFsEntryPath)
		if err != nil {
			return err
		}

		if doesRealFsEntryExists {
			if err := v.revertSteps.BackupOldContent(realFsEntryPath); err != nil {
				return err
			}
		} else {
			v.revertSteps.RemoveFile(realFsEntryPath)
		}

		err = os.Remove(realFsEntryPath)
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		err = MoveFile(mergeDirEntryPath, realFsEntryPath)

		if err != nil {
			return err
		}

		if err := filePrototype.Save(); err != nil {
			return err
		}
	}

	return nil
}

func mergePrototypes(prototypesDirPath, destPath string) error {
	dirs, err := os.ReadDir(prototypesDirPath)
	if err != nil {
		return err
	}

	for _, dirEntry := range dirs {
		if dirEntry.IsDir() {
			dirName := dirEntry.Name()
			if err := os.MkdirAll(filepath.Join(destPath, dirName), os.ModePerm); err != nil {
				return err
			}
			if err := mergePrototypes(filepath.Join(prototypesDirPath, dirName), filepath.Join(destPath, dirName)); err != nil {
				return err
			}
			continue
		}
		prototypeName := dirEntry.Name()
		if strings.Contains(prototypeName, ".prototype.bson") {
			fileName := strings.ReplaceAll(prototypeName, ".prototype.bson", "")

			prototype := FilePrototype{}
			if err := prototype.Read(prototypesDirPath, fileName); err != nil {
				return err
			}
			file, err := prototype.SimulateFile()
			if err != nil {
				return err
			}

			if err := os.WriteFile(filepath.Join(destPath, fileName), file, os.ModePerm); err != nil {
				return err
			}
			continue
		}
	}

	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return false, err
	}

	return !os.IsNotExist(err), nil
}
