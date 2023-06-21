package repository

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	  fx.Provide(NewCategoryRepository),
	fx.Provide(NewPostRepository),
	fx.Provide(NewUserRepository),
)
