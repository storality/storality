package exceptions

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")