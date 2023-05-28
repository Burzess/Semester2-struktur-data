package db

import (
	"progres/entity/antrian"
	"progres/entity/dokter"
	"progres/entity/pasien"
	"progres/entity/suster"
	"progres/entity/users"
)

var DBUsers users.LinkedlistUsers

var DBDokter dokter.LinkedlistDokter
var DBSuster suster.LinkedlistSuster
var DBPasien pasien.LinkedlistPasien

var DBAntrian antrian.LinkedlistAntrian

var NoAntrian int = 0
var JumlahAntrian int = 1
