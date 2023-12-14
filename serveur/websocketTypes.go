package serveur

type InitReq struct {
	IdSimu string `json:"id_simu"`
}

type DefaultReq struct {
	T string `json:"type"`
	D string `json:"data"`
}

type Resp struct {
	Msg string `json:"msg"`
}
