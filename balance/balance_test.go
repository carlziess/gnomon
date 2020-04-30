/*
 *  Copyright (c) 2020. aberic - All Rights Reserved.
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
 */

package balance

import "testing"

func TestNewBalanceRound(t *testing.T) {
	b := NewBalance(Round)
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	for i := 0; i < 100; i++ {
		t.Log(b.Run())
	}
}

func TestNewBalanceRandom(t *testing.T) {
	b := NewBalance(Random)
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	for i := 0; i < 100; i++ {
		t.Log(b.Run())
	}
}

func TestNewBalanceHash(t *testing.T) {
	b := NewBalance(Hash)
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	for i := 0; i < 100; i++ {
		t.Log(b.Run())
	}
}
