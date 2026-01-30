package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// definisikan driver dan datasourcenya, datasource dikirim ketika memakai fungsi ini melalui argument connectionString
	db, err := sql.Open("pgx", connectionString)

	if err != nil {
		// return dari error, juga harus mengikuti kontrak awalnya, yaitu (*sql.DB, error)
		// karena error dan hanya ingin mengatakan bahwa error, maka nilai pertama diisi nil dan nilai kedua diisi error
		return nil, err
	}

	// tes koneksi database.
	err = db.Ping()
	if err != nil {
		// ini juga, jangan lupa kontrak kembaliannya. 2 ya
		return nil, err
	}

	// set database connection pool, rekomendasi di golang
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)

	// cetak log
	log.Println("Database connected successfully.")

	// terakhir, karena tidak ada yang error, maka kembalikan object sql.DB nya
	// ini adalah object yang dipakai untuk melakukan query ke database.
	// ingat kontraknya, karena tidak ada error, maka errornya diganti dengan nil, terbalik dari sebelumnya.
	return db, nil

}
