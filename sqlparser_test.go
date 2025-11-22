package sqlparser

import "testing"

func TestXxx(t *testing.T) {
	Parse("CREATE TABLE `company` (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text)")
}
