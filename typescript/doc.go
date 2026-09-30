/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package typescript builds TypeScript and TSX source from composable tokens.
// Package-level builders create token sequences and fluent methods append to
// them. Render writes readable token layout without formatting, compiling, or
// type-checking the output.
package typescript
