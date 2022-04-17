package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type nasabah struct {
	cabang       string `json:"cabang"`
	noID         string `json:"noid"`
	nama         string `json:"nama"`
	alias        string `json:"alias"`
	mName        string `json:"mname"`
	tipeID       string `json:"tipeid"`
	isLT         string `json:"islt"`
	noNPWP       string `json:"nonpwp"`
	tmptLahir    string `json:"tmptlahir"`
	tglLahir     string `json:"tgllahir"`
	almtID       string `json:"almtid"`
	rtID         string `json:"rtid"`
	rwID         string `json:"rwid"`
	provID       string `json:"provid"`
	kokabID      string `json:"kokabid"`
	kecID        string `json:"kecid"`
	kelID        string `json:"kelid"`
	kdposID      string `json:"kdposid"`
	apiProvID    string `json:"apiprovid"`
	apiKokabID   string `json:"apikokabid"`
	apiKecID     string `json:"apikecid"`
	almtNew      string `json:"almtnew"`
	rtNew        string `json:"rtnew"`
	rwNew        string `json:"rwnew"`
	provNew      string `json:"provnew"`
	kokabNew     string `json:"kokabnew"`
	kecNew       string `json:"kecnew"`
	kelNew       string `json:"kelnew"`
	kdPosNew     string `json:"kdposnew"`
	noHP         string `json:"nohp"`
	email        string `json:"email"`
	jk           string `json:"jk"`
	religion     string `json:"religion"`
	stKawin      string `json:"stkawin"`
	kwrgngraan   string `json:"kwrgngraan"`
	pekerjaan    string `json:"pekerjaan"`
	nmPerusahaan string `json:"nmperusahaan"`
	bdUsaha      string `json:"bdusaha"`
	jabatan      string `json:"jabatan"`
	income       string `json:"income"`
	lastPend     string `json:"lastpend"`
	hobby        string `json:"hobby"`
	sumberDana   string `json:"sumberdana"`
	tujPembukaan string `json:"tujpembukaan"`
	idPath       string `json:"idpath"`
	swaPath      string `json:"swapath"`
}

var nasabahs = []nasabah{
	{cabang: "1", noID: "3216061206970012", nama: "DIMAS RAKHA WISNU WARDHANA", alias: "DIMAS", mName: "NURSITI", tipeID: "ktp", isLT: "1", noNPWP: "94.177.430.9-435.000", tmptLahir: "BEKASI", tglLahir: "1997-06-12", almtID: "PAPAN MAS BLOK G52 NO 14", rtID: "003", rwID: "015", provID: "JAWA BARAT", kokabID: "BEKASI", kecID: "TAMBUN SELATAN", kelID: "MANGUNJAYA", kdposID: "17510", apiProvID: "32", apiKokabID: "16", apiKecID: "6", almtNew: "PAPAN MAS BLOK G52 NO 14", rtNew: "003", rwNew: "015", provNew: "JAWA BARAT", kokabNew: "BEKASI", kecNew: "TAMBUN SELATAN", kelNew: "MANGUNJAYA", kdPosNew: "17510", noHP: "08973204649", email: "DRAWISNU@GMAIL.COM", jk: "L", religion: "ISLAM", stKawin: "LAJANG", kwrgngraan: "ID", pekerjaan: "PEGAWAI SWASTA", nmPerusahaan: "UNIVERSAL BPR", bdUsaha: "PERBANKAN", jabatan: "STAFF", income: "Rp 5 Jt s/d Rp 25 Jt", lastPend: "STRATA 1", hobby: "KESENIAN", sumberDana: "GAJI", tujPembukaan: "SIMPANAN", idPath: "null", swaPath: "null"},
}

func getNasabahs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nasabahs)
}


func main() {
	router := gin.Default()
	router.GET("/nasabahs", getNasabahs)
	router.Run("localhost:8080")
}