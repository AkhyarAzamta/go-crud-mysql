package models

import (
	"database/sql"
	"fmt"
	// "time"

	"go-crud/config"
	"go-crud/entities"
)

type UserModel struct {
	conn *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &UserModel{
		conn: conn,
	}
}

func (p *UserModel) FindAll() ([]entities.Dataku, error) {

	rows, err := p.conn.Query("SELECT * FROM dataku")
	if err != nil {
		return []entities.Dataku{}, err
	}
	defer rows.Close()

	var dataUser []entities.Dataku
	for rows.Next() {
		var dataku entities.Dataku
		rows.Scan(&dataku.Id,
			&dataku.NamaOrangTua,
			&dataku.NoNIK,
			&dataku.Alamat,
			&dataku.Kota,
			&dataku.TitikKoordinat,
			&dataku.KodePos,
			&dataku.NoHP,
			&dataku.NamaAnak,
			&dataku.JenisSekolah,
			&dataku.Kelas,
			&dataku.JurusanPilihan,
			&dataku.PemegangKPM)

		// Menyesuaikan format tanggal
		// tgl_lahir, _ := time.Parse("2006-01-02", dataku.TanggalLahir)
		// dataku.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataUser = append(dataUser, dataku)
	}

	return dataUser, nil
}

func (p *UserModel) Create(dataku entities.Dataku) bool {

	result, err := p.conn.Exec("INSERT INTO dataku (nama_orang_tua, no_nik, alamat, kota, titik_koordinat, kode_pos, no_hp, nama_anak, jenis_sekolah, kelas, jurusan_pilihan, pemegang_kpm) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		dataku.NamaOrangTua, dataku.NoNIK, dataku.Alamat, dataku.Kota, dataku.TitikKoordinat, dataku.KodePos, dataku.NoHP, dataku.NamaAnak, dataku.JenisSekolah, dataku.Kelas, dataku.JurusanPilihan, dataku.PemegangKPM)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *UserModel) Find(id int64, dataku *entities.Dataku) error {

	return p.conn.QueryRow("SELECT * FROM dataku WHERE id = ?", id).Scan(
		&dataku.Id,
		&dataku.NamaOrangTua,
		&dataku.NoNIK,
		&dataku.Alamat,
		&dataku.Kota,
		&dataku.TitikKoordinat,
		&dataku.KodePos,
		&dataku.NoHP,
		&dataku.NamaAnak,
		&dataku.JenisSekolah,
		&dataku.Kelas,
		&dataku.JurusanPilihan,
		&dataku.PemegangKPM)
}

func (p *UserModel) Update(dataku entities.Dataku) error {

	_, err := p.conn.Exec(
		"UPDATE dataku SET nama_orang_tua = ?, no_nik = ?, alamat = ?, kota = ?, titik_koordinat = ?, kode_pos = ?, no_hp = ?, nama_anak = ?, jenis_sekolah = ?, kelas = ?, jurusan_pilihan = ?, pemegang_kpm = ? WHERE id = ?",
		dataku.NamaOrangTua, dataku.NoNIK, dataku.Alamat, dataku.Kota, dataku.TitikKoordinat, dataku.KodePos, dataku.NoHP, dataku.NamaAnak, dataku.JenisSekolah, dataku.Kelas, dataku.JurusanPilihan, dataku.PemegangKPM, dataku.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *UserModel) Delete(id int64) error {
	result, err := p.conn.Exec("DELETE FROM dataku WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Data tidak ditemukan")
	}

	return nil
}
