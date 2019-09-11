/*
 * Copyright (c) 2019. ENNOO - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package file 文件操作工具
package gnomon

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type fileCommon struct{}

// PathExists 判断路径是否存在
func (f *fileCommon) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// ReadFileFirstLine 从文件中逐行读取并返回字符串数组
func (f *fileCommon) ReadFileFirstLine(filePath string) (string, error) {
	fileIn, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer fileIn.Close()
	finReader := bufio.NewReader(fileIn)
	inputString, _ := finReader.ReadString('\n')
	return String().TrimN(inputString), nil
}

// ReadFileByLine 从文件中逐行读取并返回字符串数组
func (f *fileCommon) ReadFileByLine(filePath string) ([]string, error) {
	fileIn, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileIn.Close()
	finReader := bufio.NewReader(fileIn)
	var fileList []string
	for {
		inputString, err := finReader.ReadString('\n')
		//fmt.Println(inputString)
		if err == io.EOF {
			fileList = append(fileList, String().TrimN(inputString))
			break
		}
		fileList = append(fileList, String().TrimN(inputString))
	}
	//fmt.Println("fileList",fileList)
	return fileList, nil
}

// CreateAndWrite 创建并写入内容到文件中
func (f *fileCommon) CreateAndWrite(filePath string, data []byte, force bool) error {
	if exist := f.PathExists(filePath); exist && !force {
		return errors.New("file exist")
	}
	lastIndex := strings.LastIndex(filePath, "/")
	parentPath := filePath[0:lastIndex]
	if err := os.MkdirAll(parentPath, os.ModePerm); nil != err {
		return err
	}
	// 创建文件，如果文件已存在，会将文件清空
	if file, err := os.Create(filePath); err != nil {
		return err
	} else {
		defer file.Close()
		// 将数据写入文件中
		//file.WriteString(string(data)) //写入字符串
		if n, err := file.Write(data); nil != err { // 写入byte的slice数据
			return err
		} else {
			Log().Debug("CreateAndWrite", Log().Field("byte count", n))
			return nil
		}
	}
}

// LoopDirFromDir 遍历文件夹下的所有子文件夹
func (f *fileCommon) LoopDirFromDir(pathname string) ([]string, error) {
	var s []string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		Log().Debug("read dir fail", Log().Err(err))
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

// LoopAllFileFromDir 遍历文件夹及子文件夹下的所有文件
func (f *fileCommon) LoopAllFileFromDir(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		Log().Debug("read dir fail", Log().Err(err))
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = f.LoopAllFileFromDir(fullDir, s)
			if err != nil {
				Log().Debug("read dir fail", Log().Err(err))
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}