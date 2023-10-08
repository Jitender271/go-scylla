package cmd

func Execute() error{

	container:= NewContainer()
	container.server.Setup()
	container.routes.Register(container.server)

	if err := container.server.Run(); err!=nil{
		return err

	}

	defer container.dbSession.Close()
	return nil



}