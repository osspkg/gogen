/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import "go.osspkg.com/gogen/internal/models"

func (v *Tokens) Package(arg string) *Tokens {
	return v.Join(
		create().Join(
			&models.Keyword{D: "package"},
			&models.Keyword{D: arg},
		),
	)
}

func Package(arg string) *Tokens {
	return create().Package(arg)
}

//------------------------------------------------------

func (v *Tokens) Import(name, module string) *Tokens {
	return v.Join(
		create().Join(
			&models.Keyword{D: "import"},
			&models.Keyword{D: name, Verify: true},
			&models.Text{D: module},
		),
	)
}

func Import(name, module string) *Tokens {
	return create().Import(name, module)
}
