package routes

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters() *chi.Mux

func (r *Router) setConfigRouters()

func RouterHealth()

func RouterProduct()

func EnableTimeout()

func EnableCORS()

func EnableRecovery()

func EnableRequestID()

func EnableRealIP()
