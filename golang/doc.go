/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package golang builds Go source from composable tokens. Package-level
// constructors create token sequences, and fluent methods append declarations,
// expressions, and statements to them.
//
// Render applies go/format by default. Tokens.Render writes the readable token
// layout without formatting, and SetRawMode disables formatting for subsequent
// package-level Render calls. Raw tokens are always emitted verbatim.
package golang
