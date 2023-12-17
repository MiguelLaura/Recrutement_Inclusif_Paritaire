package serveur

type InitReq struct {
	IdSimu string `json:"id_simulation"`
}

type ActionReq struct {
	IdSimu string `json:"id_simulation"`
	T      string `json:"type"`
	D      string `json:"data"`
}

type Resp struct {
	Msg string `json:"msg"`
}
