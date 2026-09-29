/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package golang builds Go source from composable tokens. Render formats the
// generated source with go/format by default; Tokens.Render exposes the readable
// token layout before formatting.
package golang
