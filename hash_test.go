/*
 *
 *  * Copyright (c) 2019. aberic - All Rights Reserved.
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  * http://www.apache.org/licenses/LICENSE-2.0
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 *
 */

package gnomon

import "testing"

func TestHashCommon_MD5(t *testing.T) {
	t.Log(HashMD5("haha"))
	t.Log(HashMD5("haha"))
}

func TestHashCommon_MD516(t *testing.T) {
	t.Log(HashMD516("haha"))
	t.Log(HashMD516("haha"))
}

func TestHashCommon_Sha1(t *testing.T) {
	t.Log(HashSha1("haha"))
	t.Log(HashSha1("haha"))
}

func TestHashCommon_Sha224(t *testing.T) {
	t.Log(HashSha224("haha"))
	t.Log(HashSha224("haha"))
}

func TestHashCommon_Sha256(t *testing.T) {
	t.Log(HashSha256("haha"))
	t.Log(HashSha256("haha"))
}

func TestHashCommon_Sha384(t *testing.T) {
	t.Log(HashSha384("haha"))
	t.Log(HashSha384("haha"))
}

func TestHashCommon_Sha512(t *testing.T) {
	t.Log(HashSha512("haha"))
	t.Log(HashSha512("haha"))
}

func TestHashAll(t *testing.T) {
	t.Log("------------- mad5 -------------")
	t.Log(HashMD5("haha"))
	t.Log(HashMD5("haha"))
	t.Log()
	t.Log("------------- mad516 -------------")
	t.Log(HashMD516("haha"))
	t.Log(HashMD516("haha"))
	t.Log()
	t.Log("------------- sha1 -------------")
	t.Log(HashSha1("haha"))
	t.Log(HashSha1("haha"))
	t.Log()
	t.Log("------------- sha224 -------------")
	t.Log(HashSha224("haha"))
	t.Log(HashSha224("haha"))
	t.Log()
	t.Log("------------- sha256 -------------")
	t.Log(HashSha256("haha"))
	t.Log(HashSha256("haha"))
	t.Log()
	t.Log("------------- sha384 -------------")
	t.Log(HashSha384("haha"))
	t.Log(HashSha384("haha"))
	t.Log()
	t.Log("------------- sha512 -------------")
	t.Log(HashSha512("haha"))
	t.Log(HashSha512("haha"))
	t.Log()
}
