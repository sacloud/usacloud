package command

import (
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"io/ioutil"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ObjectStoragePut(ctx Context, params *PutObjectStorageParam) error {
	if ctx.NArgs() != 2 {
		return fmt.Errorf("ObjectStoragePut is failed: %s", "Only two argument can be specified")
	}

	// validate filepath
	filePath := ""
	filePath = filepath.Clean(ctx.Args()[0])
	info, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("file[%s] is not exists: %s", filePath, err)
	}

	// path
	path := ""
	path = ctx.Args()[1]
	if path != "" && strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 1)
	}
	// if path is dir, set filename from filePath
	if strings.HasSuffix(path, "/") {
		path = fmt.Sprintf("%s%s", path, filepath.Base(filePath))
	}

	// on SakuraCloud, bucket name is same as AccessKey
	if params.Bucket == "" {
		params.Bucket = params.AccessKey
	}

	auth, err := aws.GetAuth(params.AccessKey, params.SecretKey)
	if err != nil {
		return fmt.Errorf("ObjectStoragePut is failed: %s", err)
	}
	client := s3.New(auth, aws.Region{
		Name:       "us-west-2",
		S3Endpoint: "https://b.sakurastorage.jp",
	})

	bucket := client.Bucket(params.Bucket)

	if info.IsDir() {
		params.ContentType = "" // when directory mode, set empty to content-type
		err := objectStoragePutRecursive(path, filePath, filePath, params.Recursive, bucket, params.ContentType)
		if err != nil {
			return fmt.Errorf("ObjectStoragePut is failed: %s", err)
		}

	} else {
		err := objectStoragePut(path, filePath, bucket, params.ContentType)
		if err != nil {
			return fmt.Errorf("ObjectStoragePut is failed: %s", err)
		}

	}

	return nil
}

func objectStoragePutRecursive(remotePath, baseDir, targetDir string, rec bool, bucket *s3.Bucket, contentType string) error {

	// if recursive is false , process only files under targetDir
	if !rec && targetDir != baseDir {
		return nil
	}

	entries, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return err
	}

	for _, fi := range entries {
		src := filepath.Join(targetDir, fi.Name())
		// this is used by object storage , so use path.Join(not filepath.Join)
		dest := path.Join(remotePath, fi.Name())
		if fi.IsDir() {
			err := objectStoragePutRecursive(dest, baseDir, src, rec, bucket, contentType)
			if err != nil {
				return err
			}
		} else {
			err := objectStoragePut(dest, src, bucket, contentType)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func objectStoragePut(destPath, srcPath string, bucket *s3.Bucket, contentType string) error {

	if contentType == "" {
		// set content-type from extension
		ext := filepath.Ext(srcPath)
		contentType = mime.TypeByExtension(ext)
	}

	// target file
	file, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	// put file
	err = bucket.PutReader(destPath, file, fi.Size(), contentType, s3.PublicRead)
	if err != nil {
		return err
	}
	return nil
}
