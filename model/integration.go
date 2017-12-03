// Copyright 2017 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package model

// Integration represents user integration settings.
type Integration struct {
	UserID               int64
	PinboardEnabled      bool
	PinboardToken        string
	PinboardTags         string
	PinboardMarkAsUnread bool
	InstapaperEnabled    bool
	InstapaperUsername   string
	InstapaperPassword   string
}