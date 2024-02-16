package main

import (
	"GoWarp/warp"
)

func init() {
	//cobra.AddTemplateFunc("getActiveMachine", )
}

func main() {
	warp, err := warp.GetWarpClient()
	if err != nil {
		panic(err)
	}
	//htb.SetNewAPIKey("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI1IiwianRpIjoiNDJkZmQwYjVkNDMwZmJlNWI1MzFiYmFlZDZhZTZiMmY0ZmNlM2ZiZGRjMjUzMWM5ZmZmM2RlMzg1N2I2MDc4NTRlZWUyNDJmMjAwNWVhOTAiLCJpYXQiOjE3MDgwNDE0NTUuMTYzMzEsIm5iZiI6MTcwODA0MTQ1NS4xNjMzMTEsImV4cCI6MTczOTU3NzQ1NS4xNTQ0ODMsInN1YiI6IjEzMTM0MjUiLCJzY29wZXMiOlsiMmZhIl19.o9J5W1hX8amzQase1G2UKnAEZMHNdrfboCGK0n3epaCClVCRNR78BIVCykKNVC1LIhzv7fm9qOJxSnCHlFGwkUiclEXSFASJ1ywiWorcLXWSGIhdqUQ11v3_Jv4XIRoR74Lp2PVVizhyas55-wzmw0j-qLWynsjUOVsPR5TJHo_IOj8bErckXh-VcUEAXZTNNsXZS3vML3QEJyQpWrecyGvWlsAsKRO_PuI7FYygB8_SLjf65sB1QsAlFUxhriB04ogC4grocXf5efOfbF-0QTqRmv26qWwoSxUYsAcVCYvySPIk20HtvLkOF0ik04VEboxSQ0GCZ3q1sWNunwX-PdU8HVKpnGRT3I5hM-oAJBPTB53iq-1zt5tz6yjBWRYH50AGb6lgdtqZ1tHc18R5jyYSE6-3ICZKrT0kysH6eSFzqi04cAoRTXIE4dA739xb9dZZUU4qqVx4JFx3TP74BYAXWcx5ocmvoKYBU8NY2r6_4h1w61Cffpgeo4gmnxuNIfAGnYJmww5NGG_LyXqgYzOKCwEoahhKDYrEGu9ulZlHMyaGXifmyG9mhGpvHzMATDgn-gMvBNhKTV5QHFynloDmCBNlyNQQuHm6KiLF2J-QMJtGdXWwRdDqe4tc88qjL1Mks5BuK8UhQPRH6lVrB8mJf5qJHDVYAwGsMoTfhK4")
	//m, _ := warp.ListMachines()
	//fmt.Sprintf("%v", m) 575
	// Cannot find valid machine? Also what user do we need to use my user has no data?
	// ServerID req for some user id works for none
	//warp.GetValidateMachineOwnage(575)
	//warp.GetListEndgames()
	//warp.GetEndgameProfile(302)
	//warp.ListRetiredMachines()
	warp.Start()
}
