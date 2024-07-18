package tango_helpers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// CSVReader es una "clase" que proporciona un método para leer datos desde un archivo CSV
type CSVReader struct {
	CSVRows   int // Número de columnas del cvs
	Separator rune
	BatchSize uint
	Reader    *csv.Reader
	file      *os.File
}

func NewCSVReader() *CSVReader {
	return &CSVReader{
		CSVRows:   0,
		Separator: ';',
		BatchSize: 1000,
	}
}

// LeerCSVNativos lee un archivo CSV y retorna un slice de instancias de Persona
// ejecucion: r.Parser("./personas.csv",10).All2Func()
func (r *CSVReader) Parse(filepath string) {

	// Abrir el archivo CSV
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Cannot open file: %v\n", err)
	}
	r.file = file
	//defer file.Close()

	// Crear un lector CSV
	r.Reader = csv.NewReader(file)
	r.Reader.Comma = r.Separator

	r.AutoSetRows()

}

func (r *CSVReader) AutoSetRows() {
	rows, err := r.ReadOneLine()
	if err != nil {
		log.Printf("Error reading first line: %v \n", err)
	}
	r.CSVRows = len(rows)
}

func (r *CSVReader) ReadOneLine() ([]string, error) {

	line, err := r.Reader.Read()
	if err != nil {
		if r.IsEOFError(err) {
			log.Printf("CSV > Error: %v\n", err)
			return nil, err
		}
	}
	return line, nil
}

func (r *CSVReader) CloseFile() {
	r.file.Close()
}

func (r *CSVReader) All(f func([]string) error) error {
	defer r.CloseFile()
	// Leer todas las filas del archivo CSV
	rows, err := r.Reader.ReadAll()
	if err != nil {
		return err
	}

	// Iterar sobre las filas y almacenar en la estructura Persona
	for _, row := range rows {
		if len(row) != r.CSVRows {
			return fmt.Errorf("Incorrect line file format: %v", row)
		}

		// aqui se ejecuta la funcion
		err := f(row)
		if err != nil {
			return err
		}
	}

	return nil

}

// TODO: Revisar esta sección de código.
func (r *CSVReader) Batch(f func(*[][]string) error) error {
	defer r.CloseFile()
	var counter uint = 0
	var batch [][]string
	// Itera sobre las líneas del archivo
	for {
		// Lee un lote de registros
		row, err := r.Reader.Read()
		if err != nil {
			if r.IsEOFError(err) {
				log.Printf("CSV > Error: %v\n", err)
				break // Fin del archivo
			}
		}

		batch = append(batch, row)
		counter = +1
		// Si el tamaño del lote alcanza el límite, procesa el lote
		if counter == r.BatchSize {
			err := f(&batch)
			if err != nil {
				return err
			}
			batch = nil
			counter = 0
		}
	}
	return nil

}

func (r *CSVReader) IsEOFError(err error) bool {
	err_str := fmt.Sprintf("%s", err)
	if err_str != "EOF" {
		return true
	}
	return false
}

func (r *CSVReader) StrToInt(s string) int {
	if s != "" {

		num, err := strconv.Atoi(s)
		if err != nil {
			msg := fmt.Errorf("Error trying to convert string (\"%s\") to int: %v", s, err)
			log.Println(msg)
		}
		return num
	}
	return 0
}
