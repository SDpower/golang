// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import "math"

/*
	arg1 ^ arg2 (exponentiation)
 */

export func Pow(arg1,arg2 float64) float64 {
	if arg2 < 0 {
		return 1/Pow(arg1, -arg2);
	}
	if arg1 <= 0 {
		if(arg1 == 0) {
			if arg2 <= 0 {
				return sys.NaN();
			}
			return 0;
		}

		temp := Floor(arg2);
		if temp != arg2 {
			panic(sys.NaN());
		}

		l := int32(temp);
		if l&1 != 0 {
			return -Pow(-arg1, arg2);
		}
		return Pow(-arg1, arg2);
	}

	temp := Floor(arg2);
	if temp != arg2 {
		if arg2-temp == .5 {
			if temp == 0 {
				return Sqrt(arg1);
			}
			return Pow(arg1, temp) * Sqrt(arg1);
		}
		return Exp(arg2 * Log(arg1));
	}

	l := int32(temp);
	temp = 1;
	for {
		if l&1 != 0 {
			temp = temp*arg1;
		}
		l >>= 1;
		if l == 0 {
			return temp;
		}
		arg1 *= arg1;
	}
	panic("unreachable")
}
