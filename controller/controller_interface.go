package controller

type IController interface {
	//Serve method need to be implemented to derived controller for serving all routing
	Serve()
}
