package main

import (
	"fmt"
)

type mobile struct {
	general          General
	display_features Display_features
	os_Pro           Os_Pro
	mem_Feat         Mem_Feat
	cam_feat         Cam_feat
	conn_Feat        Conn_Feat
}

type General struct {
	Mod_num   string `json:"Id"`
	Mod_name  string
	Color     string
	Sim_type  string
	Touch_scr bool
}

type Display_features struct {
	Disp_size    float32
	Resolution   int
	Display_type string
}

type Os_Pro struct {
	Os       string
	Pro_type string
	Pro_core string
}

type Mem_Feat struct {
	Int_stor    int
	Ram         int
	Expand_stor int32
}

type Cam_feat struct {
	Prim_cam_avail   bool
	Prim_cam         string
	Second_cam_avail bool
	Second_cam       int
}

type Conn_Feat struct {
	Net_type  string
	_3G       bool
	GPRS      bool
	Bluetooth bool
	Blu_ver   float32
	wifi      bool
	GPS       bool
}

func main() {

	fmt.Println("Server is starting ................")
	var xGeneral General

	xGeneral.Mod_num = "1"
	xGeneral.Mod_name = "Xiaomi"
	xGeneral.Color = "white"
	xGeneral.Sim_type = "Dual"
	xGeneral.Touch_scr = true

	var xDispFeat Display_features

	xDispFeat.Disp_size = 6.5
	xDispFeat.Display_type = "AMOLED"
	xDispFeat.Resolution = 4500

	var xOs_pro Os_Pro

	xOs_pro.Os = "Android"
	xOs_pro.Pro_core = "octa core"
	xOs_pro.Pro_type = "MediaTek"

	var xMemFeat Mem_Feat

	xMemFeat.Int_stor = 64
	xMemFeat.Expand_stor = 512
	xMemFeat.Ram = 8

	var xCamFeat Cam_feat

	xCamFeat.Prim_cam_avail = true
	xCamFeat.Prim_cam = "48 x 2 x 5"
	xCamFeat.Second_cam_avail = true
	xCamFeat.Second_cam = 48

	var xConn_Feat Conn_Feat

	xConn_Feat._3G = true
	xConn_Feat.wifi = true
	xConn_Feat.Bluetooth = true
	xConn_Feat.Blu_ver = 5.1
	xConn_Feat.wifi = true
	xConn_Feat.GPRS = true
	xConn_Feat.GPS = true
	xConn_Feat.Net_type = "4G"

	var yGeneral General

	yGeneral.Mod_num = "2"
	yGeneral.Mod_name = "Xiaomi"
	yGeneral.Color = "white"
	yGeneral.Sim_type = "Dual"
	yGeneral.Touch_scr = true

	var yDispFeat Display_features

	yDispFeat.Disp_size = 6.5
	yDispFeat.Display_type = "AMOLED"
	yDispFeat.Resolution = 4500

	var yOs_pro Os_Pro

	yOs_pro.Os = "Android"
	yOs_pro.Pro_core = "octa core"
	yOs_pro.Pro_type = "MediaTek"

	var yMemFeat Mem_Feat

	yMemFeat.Int_stor = 64
	yMemFeat.Expand_stor = 512
	yMemFeat.Ram = 8

	var yCamFeat Cam_feat

	yCamFeat.Prim_cam_avail = true
	yCamFeat.Prim_cam = "48 x 2 x 5"
	yCamFeat.Second_cam_avail = true
	yCamFeat.Second_cam = 48

	var yConn_Feat Conn_Feat

	yConn_Feat._3G = true
	yConn_Feat.wifi = true
	yConn_Feat.Bluetooth = true
	yConn_Feat.Blu_ver = 5.1
	yConn_Feat.wifi = true
	yConn_Feat.GPRS = true
	yConn_Feat.GPS = true
	yConn_Feat.Net_type = "4G"

	mob :=[]mobile {
		{
			general: xGeneral,
			cam_feat: xCamFeat,
			conn_Feat: ,
		}
	

	}

	x[0].general = xGeneral
	x[0].cam_feat = xCamFeat
	x[0].conn_Feat = xConn_Feat
	x[0].os_Pro = xOs_pro
	x[0].display_features = xDispFeat
	x[0].mem_Feat = xMemFeat

	x[1].general = yGeneral
	x[1].cam_feat = yCamFeat
	x[1].conn_Feat = yConn_Feat
	x[1].os_Pro = yOs_pro
	x[1].display_features = yDispFeat
	x[1].mem_Feat = yMemFeat

	/*
		var x mobile

		x.general = xGeneral
		x.cam_feat = xCamFeat
		x.conn_Feat = xConn_Feat
		x.os_Pro = xOs_pro
		x.display_features = xDispFeat
		x.mem_Feat = xMemFeat
	*/

	fmt.Println(x[0])
	fmt.Println(x[1])
}
