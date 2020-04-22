package position
type Position struct{
	nopol string
	imei string
	tdate string
	sdate string
	speed uint8
	acc uint8
	charge uint8
	fcut uint8
	lat int32
	lng int32
	poi string
	gf string
	gf_id uint16
	poi_id uint16
	addr string
	park_info string
}

func NewPosition()  *Position  {
	return &Position{}
}
func (p Position) toString()string{
	var info string=""
	info="Nopol "+ p.nopol+"\r\n"
	info +="Imei "+ p.imei+"\r\n"
	info +="GDate "+ p.tdate+"\r\n"
	info +="SDate "+ p.sdate+"\r\n"
	return info
}
