package twitch

const clientId = "tkro9r2rqee1s95hhyecyfq979lky7"

type Integration struct {
	token string
	user  userData
	chat *Chat
}
