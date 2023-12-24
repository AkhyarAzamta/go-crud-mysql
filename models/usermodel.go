package models

import (
	"database/sql"
	"fmt"
	"time"

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

	rows, err := p.conn.Query("select * from dataku")
	if err != nil {
		return []entities.Dataku{}, err
	}
	defer rows.Close()

	var dataUser []entities.Dataku
	for rows.Next() {
		var dataku entities.Dataku
		rows.Scan(&dataku.Id,
			&dataku.NamaLengkap,
			&dataku.NIK,
			&dataku.JenisKelamin,
			&dataku.TempatLahir,
			&dataku.TanggalLahir,
			&dataku.Alamat,
			&dataku.NoHp)

		if dataku.JenisKelamin == "1" {
			dataku.JenisKelamin = "Laki-laki"
		} else {
			dataku.JenisKelamin = "Perempuan"
		}
		// 2006-01-02 => yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2006-01-02", dataku.TanggalLahir)
		// 02-01-2006 => dd-mm-yyyy
		dataku.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataUser = append(dataUser, dataku)
	}

	return dataUser, nil

}

func (p *UserModel) Create(dataku entities.Dataku) bool {

	result, err := p.conn.Exec("insert into dataku (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values(?,?,?,?,?,?,?)",
		dataku.NamaLengkap, dataku.NIK, dataku.JenisKelamin, dataku.TempatLahir, dataku.TanggalLahir, dataku.Alamat, dataku.NoHp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *UserModel) Find(id int64, dataku *entities.Dataku) error {

	return p.conn.QueryRow("select * from dataku where id = ?", id).Scan(
		&dataku.Id,
		&dataku.NamaLengkap,
		&dataku.NIK,
		&dataku.JenisKelamin,
		&dataku.TempatLahir,
		&dataku.TanggalLahir,
		&dataku.Alamat,
		&dataku.NoHp)
}

func (p *UserModel) Update(dataku entities.Dataku) error {

	_, err := p.conn.Exec(
		"update dataku set nama_lengkap = ?, nik = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_hp = ? where id = ?",
		dataku.NamaLengkap, dataku.NIK, dataku.JenisKelamin, dataku.TempatLahir, dataku.TanggalLahir, dataku.Alamat, dataku.NoHp, dataku.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *UserModel) Delete(id int64) {
	p.conn.Exec("delete from dataku where id = ?", id)
}
