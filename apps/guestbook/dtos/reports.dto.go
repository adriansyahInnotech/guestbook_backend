package dtos

type Reports struct {
}

type DataCard struct {
	Nik                  int64  `json:"nik"`
	Nama                 string `json:"nama"`
	TempatLahir          string `json:"tempat_lahir"`
	TanggalLahir         string `json:"tanggal_lahir"`
	JenisKelamin         string `json:"jenis_kelamin"`
	GolonganDarah        string `json:"golongan_darah"`
	Alamat               string `json:"alamat"`
	Rt                   string `json:"rt"`
	Rw                   string `json:"rw"`
	Kelurahan            string `json:"kelurahan"`
	Kecamatan            string `json:"kecamatan"`
	Kabupaten            string `json:"kabupaten"`
	Provinsi             string `json:"provinsi"`
	Agama                string `json:"agama"`
	StatusPerkawinan     string `json:"status_perkawinan"`
	Pekerjaan            string `json:"pekerjaan"`
	Kewarganegaraan      string `json:"kewarganegaraan"`
	BerlakuHingga        string `json:"berlaku_hingga"`
	FotoBase64           string `json:"foto_base64"`
	TandaTanganBase64    string `json:"tanda_tangan_base64"`
	BiometrikKiriBase64  string `json:"biometrik_kiri_base64"`
	BiometrikKananBase64 string `json:"biometrik_kanan_base64"`
	IsVerifyFingerprint  bool   `json:"is_verify_fingerprint"`
}
