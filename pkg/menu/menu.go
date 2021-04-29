package menu

import pearls "github.com/minibase-app/pearls/types"

func New(items ...string) pearls.Menu {
	return pearls.Menu{Items: items}
}
