package entities

type Dataku struct {
	Id             int64
	NamaOrangTua   string `validate:"required" label:"1. Nama Orang Tua"`
	NoNIK          string `validate:"required" label:"2. No. NIK"`
	Alamat         string `validate:"required" label:"3. Alamat (RT/RW/Desa-Kelurahan/Kecamatan/Kabupaten)"`
	Kota         string `validate:"required" label:"3. Kota (RT/RW/Desa-Kelurahan/Kecamatan/Kabupaten)"`
	TitikKoordinat string `label:"4. Titik Koordinat"`
	KodePos        string `label:"5. No. Kode Pos"`
	NoHP           string `validate:"required" label:"6. No. HP"`
	NamaAnak       string `label:"7. Nama Anak"`
	JenisSekolah   string `label:"8. Jenis Sekolah"`
	Kelas          string `label:"9. Kelas/Setara Kelas"`
	JurusanPilihan string `label:"10. Jurusan yang diminati - Pilihan 1"`
	PemegangKPM string `label:"11. Pemegang KPM BLT/PKH/lainnya" validate:"oneof=ya tidak"`
}
