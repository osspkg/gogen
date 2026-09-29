/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package types

import "io"

// Token renders source text to a writer.
type Token interface {
	// Render writes the token's source text to w and returns any write or render error.
	Render(w io.Writer) error
}
