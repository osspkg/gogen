/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package python builds Python 3.10+ source from composable tokens. It includes
// builders for declarations, control flow, expressions, values, and built-in or
// generic type annotations. Suites use four-space indentation; callers insert
// explicit Line tokens between top-level statements. Rendering writes readable
// source directly and does not format, parse, or type-check the generated code.
package python
