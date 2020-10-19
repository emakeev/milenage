/*
This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package milenage

import "fmt"

// mockRNG yields a constant byte sequence instead of generating a new random sequence each time
type mockRNG []byte

func (rng mockRNG) Read(b []byte) (int, error) {
	copy(b, rng)
	if len(b) <= len(rng) {
		return len(b), nil
	}
	return len(rng), fmt.Errorf("not enough data to read")
}

// NewMockCipher instantiates the Milenage algo using non-mutable mockRNG for rng
func NewMockCipher(amf []byte, rand []byte) (*Cipher, error) {
	cipher, err := NewCipher(amf)
	if err != nil {
		return nil, err
	}
	cipher.SetRng(mockRNG(rand))
	return cipher, nil
}
