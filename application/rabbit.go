package application

import "publisher/domain"

type GetData struct {
	repo domain.Idata
}

func NewGetData(repo domain.Idata)*GetData{
	return &GetData{repo: repo}
}


func (r *GetData)Execute(data domain.DatosSensor)error{
	return r.repo.DataMqtt(data)
}