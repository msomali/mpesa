/*
 * MIT License
 *
 * Copyright (c) 2021 TECHCRAFT TECHNOLOGIES CO LTD.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package mpesa

const (
	TANZANIA = market(0)
	GHANA    = market(1)
)

type market int

//Create
//Description	URL Context Value	input_Country Value	input_Currency Value
//Vodafone Ghana	vodafoneGHA	GHA	GHS
//Vodacom Tanzania	vodacomTZN	TZN	TZS
func (m market) Create() Market {
	switch m {
	case 0:
		return Market{
			URLContextValue: "vodacomTZN",
			Country:         "TZN",
			Currency:        "TZS",
			Description:     "Vodacom Tanzania",
		}

	case 1:
		return Market{
			URLContextValue: "vodafoneGHA",
			Country:         "GHA",
			Currency:        "GHS",
			Description:     "Vodafone Ghana",
		}
	}

	return Market{}
}
