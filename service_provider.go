package goal_ten_api

import "github.com/goal-web/contracts"

func NewService() contracts.ServiceProvider {
	return Service{}
}

type Service struct {
}

func (s Service) Register(application contracts.Application) {
	application.Singleton("tenAPI", func(config contracts.Config) TenAPI {
		conf, _ := config.Get("ten-api").(Config)
		return &Client{BaseUrl: conf.BaseUrl}
	})
}

func (s Service) Start() error {
	return nil
}

func (s Service) Stop() {
}
