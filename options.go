// Copyright 2024 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

type ListOptions struct {
	Page    int `url:"page"`
	PerPage int `url:"per_page"`
}
