/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package main

import (
	"io/ioutil"
	"os"
	"path"
)

const (
	SSH_PRIVATE_KEY_DIR = ".ssh"
	SSH_PRIVATE_KEY_FILE_NAME = "id_rsa"
)

func CreateSshPrivateKeyOnDisk(fileId int, sshPrivateKeyContent string) error {

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	sshPrivateKeyFilePath := path.Join(userHomeDirectory, SSH_PRIVATE_KEY_DIR, SSH_PRIVATE_KEY_FILE_NAME)

	// if file exists then delete file
	if _, err := os.Stat(sshPrivateKeyFilePath); os.IsExist(err) {
		os.Remove(sshPrivateKeyFilePath)
	}

	// create file with content
	err = ioutil.WriteFile(sshPrivateKeyFilePath, []byte(sshPrivateKeyContent), 0600)
	if err != nil {
		return err
	}

	return nil
}


func CreateGitCredentialFileAndPutData(data string) error {

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fileName := path.Join(userHomeDirectory, ".git-credentials")

	// if file exists then delete file
	if _, err := os.Stat(fileName); os.IsExist(err) {
		os.Remove(fileName)
	}

	// create file with content
	err = ioutil.WriteFile(fileName, []byte(data), 0600)
	if err != nil {
		return  err
	}

	return nil
}


func CleanupAfterFetchingHttpsSubmodules() error {

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// remove ~/.git-credentials
	gitCredentialsFile := path.Join(userHomeDirectory, ".git-credentials")
	if _, err := os.Stat(gitCredentialsFile); os.IsExist(err) {
		os.Remove(gitCredentialsFile)
	}

	return nil
}