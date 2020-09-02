package http

import "gitlab.kroger.com/platform/krogo/pkg/krogo"

type Fav_loc interface {
	krogo.RestCreator
	krogo.RestReader
	krogo.RestDeleter
}
